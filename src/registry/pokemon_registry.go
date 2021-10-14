package registry

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
	repoCSV "github.com/gerajuarez/wize-academy-go/usecase/repository/csv"
	repoAPI "github.com/gerajuarez/wize-academy-go/usecase/repository/extAPI"
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

func (r *registry) RegisterPokemonApiCon() controller.PokemonController {
	return controller.NewPokemonController(r.RegisterPokemonApiInter())
}

func (r *registry) RegisterPokemonApiInter() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.RegisterPokemonApiRepo())
}

func (r *registry) RegisterPokemonApiRepo() repository.PokemonRepository {
	return repoAPI.NewExtApiRepo(r.csvFile)
}
