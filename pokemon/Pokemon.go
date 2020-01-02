package pokemon

//type Pokemon struct {
//name    string
//number  int
//hp      int
//attack  int
//defense int
//spAtk   int
//spDef   int
//speed   int
//type1   string
//type2   string
//}

type Pkmn struct {
	Number  int    `json:"number"`
	Name    string `json:"name"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	SpAtk   int    `json:"spAtk"`
	SpDef   int    `json:"spDef"`
	Speed   int    `json:"speed"`
	Type1   string `json:"type1"`
	Type2   string `json:"type2"`
}

type Array struct {
	PokemonArray []Pkmn `json:"pokemon"`
}

//func populatePokemon() []Pokemon {
//GetObjectFromJSON("../pokemondata/stats.json")
//}
