package internal

import (
	"encoding/json"
	"fmt"
	"github.com/jwl/nimbasa-grand-central/pokemon"
	"io/ioutil"
	"os"
	"strconv"
)

// Extract Pokemon data from stats.json

// Extract object from JSON
func GetObjectFromJSON(jsonObject string) string {
	testJson := `{"testkey":"testvalue"}`

	var parsedJson string

	json.Unmarshal([]byte(testJson), &parsedJson)

	return parsedJson

}

func GetPokemonFromFile(path string) string {
	//jsonFile, err := os.Open("./pokemondata/stats.json")
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		fmt.Println("Successfully opened stats.json")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pokemonFromFile pokemon.PokemonArray

	json.Unmarshal(byteValue, &pokemonFromFile)

	fmt.Println("json unmarshaled, about to print contents of pokemonFromFile")
	fmt.Println("len(pokemonFromFile.PokemonArray) is: " + strconv.Itoa(len(pokemonFromFile.PokemonArray)))

	for i := 0; i < len(pokemonFromFile.PokemonArray); i++ {
		fmt.Println("Pokemon Name: " + pokemonFromFile.PokemonArray[i].Name)
		fmt.Println("Pokemon Number: " + strconv.Itoa(pokemonFromFile.PokemonArray[i].Number))
	}

	return ""
}
