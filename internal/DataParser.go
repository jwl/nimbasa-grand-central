package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jwl/nimbasa-grand-central/pokemon"
)

func GetPokemonFromFile(path string) []pokemon.Pkmn {
	jsonFile, err := os.Open(path)

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
