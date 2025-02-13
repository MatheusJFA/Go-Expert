package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	// Client é uma estrutura que contém um cliente HTTP
	client := http.Client{}

	// NewRequest cria uma nova requisição HTTP
	request, err := http.NewRequest("GET", "http://www.google.com", nil)

	request.Header.Add("User-Agent", "Golang")       // Adiciona um cabeçalho à requisição
	request.Header.Add("Accept", "application/json") // Adiciona um cabeçalho à requisição

	if err != nil {
		panic(err)
	}

	// Do envia uma requisição HTTP e retorna uma resposta
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	// Fecha o corpo da resposta
	defer response.Body.Close()

	// Copia o corpo da resposta para o stdout
	io.CopyBuffer(os.Stdout, response.Body, nil)

	// Cria um buffer de bytes com um JSON
	jsonVar := bytes.NewBuffer([]byte(`{"nome":"Matheus"}`))

	// Cria uma nova requisição HTTP
	request, err = http.NewRequest("POST", "http://www.google.com", jsonVar)

	// Provavelmente você não terá permissão para fazer um POST no google.com, então é esperado um retorno 405 (Method Not Allowed)
	if err != nil {
		panic(err)
	}

	response, err = client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	io.CopyBuffer(os.Stdout, response.Body, nil)
}
