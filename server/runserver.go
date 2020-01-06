package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/jwl/nimbasa-grand-central/internal"
)

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonData := internal.GetPokemonFromFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemonData)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonData := internal.GetPokemonFromFile()
	w.Header().Set("Content-Type", "application/json")

	targetNumber, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("id passed to pokemon/{id} is not an int")
	}

	for _, singlePokemon := range pokemonData {
		if singlePokemon.Number == targetNumber {
			json.NewEncoder(w).Encode(singlePokemon)
			fmt.Printf("old method: %#v", singlePokemon)
			fmt.Printf("new method: %#v", internal.GetSinglePokemon(targetNumber))
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
