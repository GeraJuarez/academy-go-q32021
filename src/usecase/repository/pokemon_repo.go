package repository

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
)

// ErrorKeyNotFound is returned when an ID is not found in the repository
var ErrorKeyNotFound = errors.New("pokemon not found")

// ErrorItemZeroParam is returned when the parameter items is lower than 0
var ErrorItemZeroParam = errors.New("item must be higher than 0")

// ErrorWorkerZeroParam returned when the parameter itemsPerWorker is lwer than 0
var ErrorWorkerZeroParam = errors.New("items per worker must be higher than 0")

const (
	ALL_ITEM_QUERY = -1
)

// PokemonRepository implements the direct usage of a Pokemon data source
type PokemonRepository interface {
	Get(id int) (model.Pokemon, error)
	GetAllValid(items int, itemsPerWorker int, isValid func(id int) bool) ([]model.Pokemon, error)
	PostById(id int) (model.Pokemon, error)
}

// IsEven validates if a number is divisible by the number two
func IsEven(num int) bool {
	return num%2 == 0
}

// IsOdd validates if a number is not divisible by the number two
func IsOdd(num int) bool {
	return !IsEven(num)
}
