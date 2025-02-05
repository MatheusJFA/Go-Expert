package main

import (
	"GoExpert/entity"
	"net/http"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/abilities", GetAbilities)
	mux.HandleFunc("/pokemon", GetPokemon)
	http.ListenAndServe(port, mux)

}

func GetAbilities(writer http.ResponseWriter, request *http.Request) {
	ability := &entity.Ability{}
	ability.ServeHTTP(writer, request)
}

func GetPokemon(writer http.ResponseWriter, request *http.Request) {
	pokemon := &entity.Pokemon{}
	pokemon.ServeHTTP(writer, request)
}
