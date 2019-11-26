package main

import (
	"fmt"
	"github.com/jwl/nimbasa-grand-central/internal"
)

func main() {
	fmt.Println("Hello world")
	internal.GetPokemonFromFile("./pokemondata/stats.json")
}
