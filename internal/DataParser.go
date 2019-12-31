package internal

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "strconv"

  "github.com/jwl/nimbasa-grand-central/pokemon"
)

// Extract Pokemon data from stats.json

// GetObjectFromJSON extracts a data object from a json object
func GetObjectFromJSON(jsonObject string) string {
  testJSON := `{"testkey":"testvalue"}`

  var parsedJSON string

  json.Unmarshal([]byte(testJSON), &parsedJSON)

  return parsedJSON

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

  fmt.Println("byteValue is: " + string(byteValue))

  // Declared an empty interface of type Array
  var results []map[string]interface{}

  // Unmarshal or Decode the JSON to the interface.
  json.Unmarshal(byteValue, &results)

  for key, result := range results {
    fmt.Println("Reading Value for Key :", key)
    fmt.Println("\tnumber :", result["number"], "/ Name :", result["name"])
  }

  //json.Unmarshal(byteValue, &pokemonFromFile)

  //for key, pokemon := range pokemonFromFile {
    //fmt.Println("reading value for key: ", key)
    //fmt.Println("\tnumber: ", pokemon["number"], "/ Name: ", pokemon["name"])
  //}

  //err = json.Unmarshal(byteValue, &pokemonFromFile)
  //if err != nil {
  //fmt.Println("error:", err)
  //}

  //fmt.Println("json unmarshaled, about to print contents of pokemonFromFile")
  //fmt.Println("pokemonFromFile is: " + pokemonFromFile)
  //fmt.Println(json.MarshalIndent(pokemonFromFile, "", "    "))
  //fmt.Println("len(pokemonFromFile.PokemonArray) is: " + strconv.Itoa(len(pokemonFromFile.PokemonArray)))
  //fmt.Println("pokemonFromFile: " + pokemonFromFile)

  //for i := 0; i < len(pokemonFromFile.PokemonArray); i++ {
  //fmt.Println("Pokemon Name: " + pokemonFromFile.PokemonArray[i].Name)
  //fmt.Println("Pokemon Number: " + strconv.Itoa(pokemonFromFile.PokemonArray[i].Number))
  //}

  return ""
}
