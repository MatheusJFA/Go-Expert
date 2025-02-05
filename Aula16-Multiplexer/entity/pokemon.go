package entity

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string        `json:"name"`
	Order         int           `json:"order"`
	PastAbilities []interface{} `json:"past_abilities"`
	PastTypes     []interface{} `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (p *Pokemon) GetAbilities() []string {
	var abilities []string
	for _, ability := range p.Abilities {
		abilities = append(abilities, ability.Ability.Name)
	}
	return abilities
}

func (p *Pokemon) GetTypes() []string {
	var types []string
	for _, t := range p.Types {
		types = append(types, t.Type.Name)
	}
	return types
}

func (p *Pokemon) GetStats() map[string]int {
	stats := make(map[string]int)
	sum := 0
	for _, s := range p.Stats {
		stats[s.Stat.Name] = s.BaseStat
		sum += s.BaseStat
	}

	stats["effort"] = p.Stats[0].Effort
	stats["total"] = sum
	return stats
}

func (p *Pokemon) GetMoveSet() map[string]interface{} {
	var moves, tm, tutor []string

	for _, m := range p.Moves {
		for _, v := range m.VersionGroupDetails {
			name := m.Move.Name
			level := strconv.Itoa(v.LevelLearnedAt)
			learnMethod := v.MoveLearnMethod.Name

			if learnMethod == "level-up" {
				moves = append(moves, name+" (level "+level+") by "+learnMethod+" method")
				break
			}

			if learnMethod == "tutor" {
				tutor = append(tutor, name+" by "+learnMethod)
				break
			}

			if learnMethod == "machine" {
				tm = append(tm, name+" by TM")
				break
			}
		}
	}

	return map[string]interface{}{"moves": moves, "tm": tm, "tutor": tutor}
}

func (p *Pokemon) GetPokemonInformation() map[string]interface{} {
	return map[string]interface{}{
		"name":      p.Name,
		"height":    p.Height,
		"weight":    p.Weight,
		"abilities": p.GetAbilities(),
		"skills":    p.GetMoveSet(),
		"types":     p.GetTypes(),
		"stats":     p.GetStats(),
	}
}

func (p *Pokemon) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pokemon := request.URL.Query().Get("name")
	if pokemon == "" {
		http.Error(writer, "Nome do Pokemon não informado", http.StatusBadRequest)
		return
	}

	site := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	response, err := http.Get(site)

	if err != nil {
		http.Error(writer, "Erro ao buscar Pokemon", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(writer, "Erro ao ler o body do response", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	var pokemonJSON Pokemon
	err = json.Unmarshal(body, &pokemonJSON)

	if err != nil {
		http.Error(writer, "Erro ao fazer o unmarshal", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// Convertendo a map[string]interface{} para []byte
	json.NewEncoder(writer).Encode(pokemonJSON.GetPokemonInformation())
}
