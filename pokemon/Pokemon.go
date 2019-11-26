package pokemon

import "../internal/Dataparser"

type Pokemon struct {
	name    string
	number  int
	hp      int
	attack  int
	defense int
	spAtk   int
	spDef   int
	speed   int
	type1   string
	type2   string
}

func populatePokemon() []Pokemon {
	GetObjectFromJSON("../pokemondata/stats.json")
}
