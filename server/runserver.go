package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"

  "github.com/jwl/nimbasa-grand-central/internal"
  //"github.com/jwl/nimbasa-grand-central/pokemon"
)

// Pokemon : Pokemon custom
//type Pokemon struct {
  //number  string `json:"id"`
  //name    string `json:"name"`
  //hp      int `json:"hp"`
  //attack  int `json:"attack"`
  //defense int `json:"defense"`
  //spAtk   int `json:"spAtk"`
  //spDef   int `json:"spDef"`
  //speed   int `json:"speed"`
  //type1   string `json:"type1"`
  //type2   string `json:"type2"`
//}


func getAllPokemon(w http.ResponseWriter, r *http.Request) {
  pokemonData := internal.GetPokemonFromFile("./pokemondata/stats.json")

  //fmt.Println("Returning all pokemon:")
  //for key, pokemon := range pokemonData {
    //fmt.Println("\tkey:", key, "/ number:", pokemon["number"], "/ Name:", pokemon["name"])
  //}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(pokemonData)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
  pokemonData := internal.GetPokemonFromFile("./pokemondata/stats.json")
  w.Header().Set("Content-Type", "application/json")

  targetNumber := mux.Vars(r)["id"]
  fmt.Println("searching for id:", targetNumber)
  //var targetPokemon pokemon.Pkmn


  for _, singlePokemon := range pokemonData {
    fmt.Println("singlePokemon.Number is:", singlePokemon.Number, "targetNumber:", targetNumber)
    //if pokemon["number"] == targetNumber {
    if singlePokemon.Number == targetNumber {
      fmt.Println("target found!")
      //targetPokemon = singlePokemon
      json.NewEncoder(w).Encode(singlePokemon)
    }
  }
  //if targetPokemon != nil {
    //json.NewEncoder(w).Encode(targetPokemon)
  //} else {
    //json.NewEncoder(w).Encode("'':''")
  //}

  //params := mux.Vars(r)

  //for _, item := range pokemonList {
    //if item.Number == params["id"] {
      //json.NewEncoder(w).Encode(item)
      //return
    //}
  //}
}

func updatePokemon(w http.ResponseWriter, r *http.Request) {
  // implement later
}

func deletePokemon(w http.ResponseWriter, r *http.Request) {
  // implement later
}

func main() {
  // populate pokemonList
  //pokemonList = append(
    //pokemonList,
    //Pokemon{Number: "1", Name: "Bulbasaur"},
    //Pokemon{Number: "4", Name: "Charmander"},
    //Pokemon{Number: "7", Name: "Squirtle"},
  //)

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
