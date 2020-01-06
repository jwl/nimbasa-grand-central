package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jwl/nimbasa-grand-central/pokemon"
)

var allPokemonReferenceFilepath = "./pokemondata/stats.json"

// GetPokemonFromFile retrieves all pokemon from reference file
func GetPokemonFromFile() []pokemon.Pkmn {
	jsonFile, err := os.Open(allPokemonReferenceFilepath)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var results []pokemon.Pkmn

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(byteValue, &results)

	return results
}

// GetSinglePokemon retrieves pokemon by number
func GetSinglePokemon(number int) pokemon.Pkmn {
	allPokemon := GetPokemonFromFile()
	var targetPokemon pokemon.Pkmn
	for _, singlePokemon := range allPokemon {
		if singlePokemon.Number == number {
			targetPokemon = singlePokemon
			break
		}
	}
	return targetPokemon
}
