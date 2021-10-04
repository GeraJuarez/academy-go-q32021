package model

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NullPokemon returns a default Pokemon struct with invalid attributes
func NullPokemon() Pokemon {
	return Pokemon{ID: -1, Name: "Null"}
}
