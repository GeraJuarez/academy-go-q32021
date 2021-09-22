package repository

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
)

var ErrorNoSuchKey = errors.New("no such key")

type PokemonRepository interface {
	Get(id int) (model.Pokemon, error)
}
