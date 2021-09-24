package repository

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
)

var ErrorKeyNotFound = errors.New("pokemon not found")

type PokemonRepository interface {
	Get(id int) (model.Pokemon, error)
}
