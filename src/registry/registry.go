package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type registry struct {
	pkmn_repo repository.PokemonRepository
}

// Registry resolves dependencies using constructor injection
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns a Registry interface for the Pokemon repository
func NewRegistry(pkmn_repo repository.PokemonRepository) Registry {
	return &registry{pkmn_repo}
}

// NewAppController starts the injection for al the respositories in the registry
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		PokemonController: r.NewPokemonController(),
		HelloController:   r.NewHelloController(),
	}
}

func (r *registry) NewHelloController() controller.HelloController {
	return controller.NewHelloController()
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
