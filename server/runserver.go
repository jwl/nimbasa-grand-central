package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Pokemon : Pokemon custom
type Pokemon struct {
	Number string `json:"id"`
	Name   string `json:"name"`
}

var pokemonList []Pokemon

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Returning all pokemon: %s", pokemonList)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonList)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range pokemonList {
		if item.Number == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updatePokemon(w http.ResponseWriter, r *http.Request) {
	// implement later
}

func deletePokemon(w http.ResponseWriter, r *http.Request) {
	// implement later
}

func main() {
	// populate pokemonList
	pokemonList = append(
		pokemonList,
		Pokemon{Number: "1", Name: "Bulbasaur"},
		Pokemon{Number: "4", Name: "Charmander"},
		Pokemon{Number: "7", Name: "Squirtle"},
	)

	r := mux.NewRouter()

	// endpionts
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API response!\nNow on multiple lines.")
	})
	r.HandleFunc("/pokemon", getAllPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{id}", getPokemon).Methods("GET")
	r.HandleFunc("/pokemon/{id}", updatePokemon).Methods("POST")
	r.HandleFunc("/pokemon/{id}", deletePokemon).Methods("DELETE")

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":5000", r))
}
