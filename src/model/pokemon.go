package model

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NullPokemon() Pokemon {
	return Pokemon{ID: -1, Name: "Null"}
}
