package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (c CEP) String() string {
	return fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
		c.Cep, c.Logradouro, c.Complemento, c.Bairro, c.Localidade, c.Uf, c.Estado, c.Regiao, c.Ibge, c.Gia, c.Ddd, c.Siafi)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o CEP: ")
	reader.Scan()
	cep := reader.Text()
	site := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	site = strings.TrimRight(site, "\n")

	request, err := http.Get(site)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close() // Atraso do recurso, ou seja, ele será chamado depois do print(string(response))

	body, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	var cepResponse CEP
	json.Unmarshal(body, &cepResponse)

	// Cria um arquivo chamado cep.txt
	filePath := "./public/cep-" + cep + ".txt"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Escreve no arquivo
	_, err = file.WriteString(cepResponse.String())

	if err != nil {
		panic(err)
	}

	fmt.Println(cepResponse)
}
