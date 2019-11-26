package pokemon

type Pokemon struct {
	Name    string
	Number  int
	HP      int
	Attack  int
	Defense int
	SpAtk   int
	SpDef   int
	Speed   int
	Type1   string
	Type2   string
}

type PokemonArray struct {
	PokemonArray []Pokemon `json:"pokemon"`
}

//func populatePokemon() []Pokemon {
//GetObjectFromJSON("../pokemondata/stats.json")
//}
