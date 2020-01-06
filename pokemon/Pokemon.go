package pokemon

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
