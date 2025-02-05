package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func CEPHandler(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	cep := request.URL.Query().Get("cep")

	if path != "/" {
		http.NotFound(writer, request)
		return
	}

	if cep == "" {
		http.Error(writer, "Informe um CEP", http.StatusBadRequest)
		return
	}

	cepJSON, err := BuscaCEP(cep)

	if err != nil {
		http.Error(writer, "Erro ao buscar CEP", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	json, err := json.Marshal(cepJSON)

	if err != nil {
		http.Error(writer, "Erro ao buscar CEP", http.StatusInternalServerError)
		return
	}

	// json.NewEncoder(writer).Encode(cepJSON) => Versão mais simples sem necessidade de realizar o json.Marshal
	writer.Write(json)
}

func main() {
	http.HandleFunc("/", CEPHandler)

	port := ":8080"
	http.ListenAndServe(port, nil)
}

func BuscaCEP(cep string) (*CEP, error) {
	site := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	site = strings.TrimRight(site, "\n")

	response, err := http.Get(site)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close() // Atraso do recurso, ou seja, ele será chamado depois do print(string(response))

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var cepJSON CEP
	err = json.Unmarshal(body, &cepJSON)

	if err != nil {
		return nil, err
	}

	return &cepJSON, nil
}

func (c CEP) String() string {
	return fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
		c.Cep, c.Logradouro, c.Complemento, c.Bairro, c.Localidade, c.Uf, c.Estado, c.Regiao, c.Ibge, c.Gia, c.Ddd, c.Siafi)
}
