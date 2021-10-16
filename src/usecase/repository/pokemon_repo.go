package repository

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
)

var ErrorKeyNotFound = errors.New("pokemon not found")
var ErrorItemZeroParam = errors.New("item must be higher than 0")
var ErrorWorkerZeroParam = errors.New("items per worker must be higher than 0")

// PokemonRepository implements the direct usage of a Pokemon data source
type PokemonRepository interface {
	Get(id int) (model.Pokemon, error)
	GetAllValid(items int, itemsPerWorker int, isValid func(id int) bool) ([]model.Pokemon, error)
	PostById(id int) (model.Pokemon, error)
}

func IsEven(num int) bool {
	return num%2 == 0
}

func IsOdd(num int) bool {
	return !IsEven(num)
}
