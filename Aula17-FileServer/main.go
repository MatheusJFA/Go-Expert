package main

import (
	"log"
	"net/http"
)

func main() {
	// Neste momento, tudo que estiver na pasta public, seja um html, uma imagem fica sendo externalizado para o cliente da aplicação, ele só necessita do caminho até lá
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()

	mux.Handle("/", fileServer) // Logo, neste ponto será chamado o fileServer

	// No ponto abaixo será chamado a função anônima se o cliente acessar o http://localhost:8080/Informacao 
	mux.HandleFunc("/Informacao", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Essa aplicação foi escrita por Matheus José durante um treinamento no curso GoExpert"))
	})

	port := ":8080"
	log.Fatal(http.ListenAndServe(port, mux))

}
