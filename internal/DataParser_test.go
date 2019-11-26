package internal

import "testing"

func TestGetPokemonFromFile(t *testing.T) {
	GetPokemonFromFile("dummy.json")
	t.Error("TestGetPokemonFromFile is incomplete!")
}

func TestGetJSONFromFile(t *testing.T) {
	//GetJSONFromFile("test.json")
	t.Error("test unwritten!")
}

func TestGetObjectFromJSON(t *testing.T) {
	testJSON := `{"testkey": "testvalue"}`
	//GetObjectFromJSON(`{"testkey": "testvalue"}`)
	t.Error("TestGetObjectFromJSON test not yet written!")
}
