package interactor

import (
	"errors"

	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

var ErrorKeyNotFound = errors.New("key not found")

type pokemonInteractor struct {
	repo repository.PokemonRepository
}

type PokemonInteractor interface {
	Get(id int) (model.Pokemon, error)
}

func NewPokemonInteractor(repo repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{repo}
}

func (inter *pokemonInteractor) Get(id int) (model.Pokemon, error) {
	val, err := inter.repo.Get(id)
	if errors.Is(err, repository.ErrorNoSuchKey) {
		return model.Pokemon{}, ErrorKeyNotFound
	}
	if err != nil {
		return model.Pokemon{}, err
	}

	return val, nil
}
