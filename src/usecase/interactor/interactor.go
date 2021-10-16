package interactor

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

var ErrorInvalidTypeParam = errors.New("invalid type paramter")

type pokemonInteractor struct {
	repo repository.PokemonRepository
}

// PokemonInteractor implements the usage of the Pokemon repository
// applying specific application business rules
type PokemonInteractor interface {
	Get(id int) (model.Pokemon, error)
	GetAllByType(typeStr string, items int, itemsPerWorker int) ([]model.Pokemon, error)
	PostById(id int) (model.Pokemon, error)
}

// NewPokemonInteractor returns a PokemonInteractor with the given repo
func NewPokemonInteractor(repo repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{repo}
}

func (inter *pokemonInteractor) Get(id int) (model.Pokemon, error) {
	val, err := inter.repo.Get(id)

	// NOTE: if something is processed in this layer
	// do not forget to check for errors
	// do not "mask" errors from previous layers

	return val, err
}

func (inter *pokemonInteractor) GetAllByType(typeStr string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	validationFunc, err := getValidationFunction(typeStr)
	if err != nil {
		return nil, err
	}

	values, err := inter.repo.GetAllValid(items, itemsPerWorker, validationFunc)

	return values, err
}

func (inter *pokemonInteractor) PostById(id int) (model.Pokemon, error) {
	val, err := inter.repo.PostById(id)

	return val, err
}

func getValidationFunction(filter string) (func(id int) bool, error) {
	switch filter {
	case "odd":
		return repository.IsOdd, nil
	case "even":
		return repository.IsEven, nil
	default:
		return nil, ErrorInvalidTypeParam
	}
}
