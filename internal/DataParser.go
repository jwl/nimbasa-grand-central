package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Extract object from JSON
func GetObjectFromJSON(jsonObject string) string {
	testJson := `{"testkey":"testvalue"}`

	var parsedJson string

	json.Unmarshal([]byte(testJson), &parsedJson)

	return parsedJson

}

func GetJSONFromFile(path string) string {
	jsonFile, err := os.Open("stats.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened stats.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)

	return ""
}
