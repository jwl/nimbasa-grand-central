package main

import (
  "fmt"
  "github.com/jwl/nimbasa-grand-central/internal"
)

func main() {
  fmt.Println("Hello world")
  pokemonData := internal.GetPokemonFromFile("./pokemondata/stats.json")
  for key, pokemon := range pokemonData {
    fmt.Println("Reading Value for Key :", key)
    fmt.Println("\tnumber :", pokemon["number"], "/ Name :", pokemon["name"])
  }

}
