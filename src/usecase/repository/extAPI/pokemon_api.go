package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gerajuarez/wize-academy-go/common"
	pokeAPI "github.com/gerajuarez/wize-academy-go/infrastructure/poke_api"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type extApiRepo struct {
	filePath string
}

// NewExtApiRepo creates the repository that uses the pokeAPI client
// it uses the given csvPath to save the responses obtained from the external API
func NewExtApiRepo(csvPath string) repository.PokemonRepository {
	return &extApiRepo{
		filePath: csvPath,
	}
}

func (api *extApiRepo) Get(id int) (model.Pokemon, error) {
	pokeClient := pokeAPI.NewPokeAPIClient()
	body, statusCode, err := pokeClient.GetPokemonByID(id)
	if err != nil {
		return model.NullPokemon(), err
	}

	if statusCode == http.StatusNotFound {
		return model.NullPokemon(), repository.ErrorKeyNotFound
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		return model.NullPokemon(), fmt.Errorf("PokeAPI Error: %s", string(body))
	}

	var pkmn model.Pokemon
	json.Unmarshal(body, &pkmn)

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
