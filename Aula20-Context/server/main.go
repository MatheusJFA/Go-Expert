package main

import (
	"log"
	"net/http"
	"time"
)

var TWO_SECONDS = 2 * time.Second

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(writter http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	FIVE_SECONDS := 5 * time.Second

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
	case <-time.After(FIVE_SECONDS):
		log.Println("Request finalizada com sucesso")
		writter.Write([]byte("Request finalizada"))
	}
}
