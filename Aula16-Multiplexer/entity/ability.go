package entity

import (
	"encoding/json"
	"io"
	"net/http"
)

type Ability struct {
	EffectChanges []interface{} `json:"effect_changes"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokemon []struct {
		IsHidden bool `json:"is_hidden"`
		Pokemon  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		Slot int `json:"slot"`
	} `json:"pokemon"`
}

func (a *Ability) GetAbilities() []string {
	var abilities []string
	for _, ability := range a.Pokemon {
		abilities = append(abilities, ability.Pokemon.Name)
		break
	}
	return abilities
}

func (a *Ability) GetEffectEntries() []string {
	var effectEntries []string
	language := "en"
	for _, effectEntry := range a.EffectEntries {
		if effectEntry.Language.Name == language {
			effectEntries = append(effectEntries, effectEntry.Effect)
			break
		}
	}
	return effectEntries
}

func (a *Ability) GetFlavorTextEntries() []string {
	var flavorTextEntries []string
	language := "en"
	for _, flavorTextEntry := range a.FlavorTextEntries {
		if flavorTextEntry.Language.Name == language {
			flavorTextEntries = append(flavorTextEntries, flavorTextEntry.FlavorText)
			break
		}
	}
	return flavorTextEntries
}

func (a *Ability) GetNames() []string {
	var names []string
	language := "en"
	for _, name := range a.Names {
		if name.Language.Name == language {
			names = append(names, name.Name)
			break
		}
	}
	return names
}

func (a *Ability) GetPokemon() []string {
	var pokemons []string
	for _, pokemon := range a.Pokemon {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}
	return pokemons
}

func (a *Ability) GetAbilitiesInformation() map[string]interface{} {
	abilitiesInformation := make(map[string]interface{})
	abilitiesInformation["names"] = a.GetNames()
	abilitiesInformation["effect_entries"] = a.GetEffectEntries()
	abilitiesInformation["flavor_text_entries"] = a.GetFlavorTextEntries()
	abilitiesInformation["pokemons"] = a.GetPokemon()
	return abilitiesInformation
}

func (a *Ability) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	habilidade := request.URL.Query().Get("name")
	if habilidade == "" {
		http.Error(writer, "Habilidade do Pokemon não informado", http.StatusBadRequest)
		return
	}

	site := "https://pokeapi.co/api/v2/ability/" + habilidade

	response, err := http.Get(site)

	if err != nil {
		http.Error(writer, "Erro ao buscar Habilidade", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(writer, "Erro ao ler o body do response", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	var abilitiesJSON Ability
	err = json.Unmarshal(body, &abilitiesJSON)

	if err != nil {
		http.Error(writer, "Erro ao fazer o unmarshal", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// Convertendo a map[string]interface{} para []byte
	json.NewEncoder(writer).Encode(abilitiesJSON.GetAbilitiesInformation())
}
