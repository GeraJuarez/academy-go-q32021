package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
	repoCSV "github.com/gerajuarez/wize-academy-go/usecase/repository/csv"
	repoAPI "github.com/gerajuarez/wize-academy-go/usecase/repository/extAPI"
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
		HelloController: r.NewHelloController(),
		PokeCSV:         r.NewPokemonController(),
		PokeAPI:         r.NewPokemonApiCon(),
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
	return repoCSV.NewPokemonCSVReader(r.csvFile)
}

func (r *registry) NewPokemonApiCon() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonApiInter())
}

func (r *registry) NewPokemonApiInter() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonApiRepo())
}

func (r *registry) NewPokemonApiRepo() repository.PokemonRepository {
	//pokeClient := pokeAPI.NewPokeAPIClient()
	return repoAPI.NewExtApiRepo(r.csvFile)
}
