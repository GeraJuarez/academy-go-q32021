package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	pokeAPI "github.com/gerajuarez/wize-academy-go/infrastructure/poke_api"
)

type registry struct {
	csvFile    string
	pokeExtAPI pokeAPI.PokeAPIClient
}

// Registry resolves dependencies using constructor injection
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns a Registry interface for the Pokemon repository
func NewRegistry(filePath string, api pokeAPI.PokeAPIClient) Registry {
	return &registry{filePath, api}
}

// NewAppController starts the injection for al the respositories in the registry
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		HelloController: r.RegisterHello(),
		PokeCSV:         r.RegisterPokemonController(),
	}
}
