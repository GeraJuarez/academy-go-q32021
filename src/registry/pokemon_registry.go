package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
	repoCSV "github.com/gerajuarez/wize-academy-go/usecase/repository/csv"
)

func (r *registry) RegisterPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.RegisterPokemonInter())
}

func (r *registry) RegisterPokemonInter() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.RegisterPokemonRepo())
}

func (r *registry) RegisterPokemonRepo() repository.PokemonRepository {
	return repoCSV.NewPokemonCSVReader(r.csvFile)
}
