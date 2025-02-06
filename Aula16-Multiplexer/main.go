package main

import (
	"GoExpert/entity"
	"net/http"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()

	ability := &entity.Ability{}
	pokemon := &entity.Pokemon{}

	mux.Handle("/abilities", ability)
	mux.Handle("/pokemon", pokemon)
	http.ListenAndServe(port, mux)

}