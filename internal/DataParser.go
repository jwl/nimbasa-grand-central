package common

import "encoding/json"

func getObjectFromJSON(path string) string {
  testJson := `{"testkey":"testvalue"}`

  var parsedJson string

  json.Unmarshal([]byte(testJson), &parsedJson)

  return parsedJson

}
