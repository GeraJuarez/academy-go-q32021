package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
)

type registry struct {
	csvFile string
}

// Registry resolves dependencies using constructor injection
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns a Registry interface for the Pokemon repository
func NewRegistry(filePath string) Registry {
	return &registry{filePath}
}

// NewAppController starts the injection for al the respositories in the registry
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		HelloController: r.RegisterHello(),
		PokeCSV:         r.RegisterPokemonController(),
	}
}
