package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type registry struct {
	pkmn_repo repository.PokemonRepository
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(pkmn_repo repository.PokemonRepository) Registry {
	return &registry{pkmn_repo}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		PokemonController: r.NewPokemonController(),
	}
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() repository.PokemonRepository {
	return r.pkmn_repo
}
