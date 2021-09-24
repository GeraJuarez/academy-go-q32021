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

	if err != nil {
		return model.Pokemon{}, err
	}

	return val, nil
}

// todo:
// csv read per request
