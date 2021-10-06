package repository

import (
	"strconv"

	"github.com/gerajuarez/wize-academy-go/common"
	pokeAPI "github.com/gerajuarez/wize-academy-go/infrastructure/poke_api"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type extApiRepo struct {
	filePath string
}

func NewExtApiRepo(csvPath string) repository.PokemonRepository {
	return &extApiRepo{
		filePath: csvPath,
	}
}

func (api *extApiRepo) Get(id int) (model.Pokemon, error) {
	pokeClient := pokeAPI.NewPokeAPIClient()
	pkmn, err := pokeClient.GetPokemonByID(id)
	if err != nil {
		return model.NullPokemon(), err
	}

	return api.Post(pkmn)
}

func (api *extApiRepo) Post(pkmn model.Pokemon) (model.Pokemon, error) {
	var data [][]string
	row := []string{strconv.Itoa(pkmn.ID), pkmn.Name}
	data = append(data, row)

	err := common.AppendCSV(api.filePath, data)
	if err != nil {
		return model.NullPokemon(), err
	}

	return pkmn, nil
}