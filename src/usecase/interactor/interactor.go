package interactor

import (
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

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

	// NOTE: if something is processed in this layer
	// do not forget to check for errors
	// do not "mask" errors from previous layers

	return val, err
}
