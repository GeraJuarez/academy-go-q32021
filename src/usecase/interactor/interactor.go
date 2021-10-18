package interactor

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

// ErrorInvalidTypeParam is thrown when the typeStr is not supported
var ErrorInvalidTypeParam = errors.New("invalid type paramter")

type pokemonInteractor struct {
	repo repository.PokemonRepository
}

// PokemonInteractor implements the usage of the Pokemon repository
// applying specific application business rules
type PokemonInteractor interface {
	Get(id int) (model.Pokemon, error)
	GetAllByType(typeStr string, itemsPerWorker int) ([]model.Pokemon, error)
	GetItemsByType(typeStr string, items int, itemsPerWorker int) ([]model.Pokemon, error)
	PostById(id int) (model.Pokemon, error)
}

// NewPokemonInteractor returns a PokemonInteractor with the given repo
func NewPokemonInteractor(repo repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{repo}
}

// Get returns a pokemon with the specified ID
func (inter *pokemonInteractor) Get(id int) (model.Pokemon, error) {
	val, err := inter.repo.Get(id)

	// NOTE: if something is processed in this layer
	// do not forget to check for errors
	// do not "mask" errors from previous layers

	return val, err
}

// GetItemsByType returns a slice of pokemons of lenght items from the specified typeStr.
func (inter *pokemonInteractor) GetItemsByType(typeStr string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	validationFunc, err := getValidationFunction(typeStr)
	if err != nil {
		return nil, err
	}

	values, err := inter.repo.GetAllValid(items, itemsPerWorker, validationFunc)

	return values, err
}

// GetAllByType returns a slice of all pokemons from the specified typeStr.
func (inter *pokemonInteractor) GetAllByType(typeStr string, itemsPerWorker int) ([]model.Pokemon, error) {
	validationFunc, err := getValidationFunction(typeStr)
	if err != nil {
		return nil, err
	}

	values, err := inter.repo.GetAllValid(repository.ALL_ITEM_QUERY, itemsPerWorker, validationFunc)

	return values, err
}

// PostById creates a record of a pokemon using only the ID and returns it.
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
