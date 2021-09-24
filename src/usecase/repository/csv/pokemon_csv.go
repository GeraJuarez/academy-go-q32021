package repository

import (
	"fmt"
	"strconv"

	"github.com/gerajuarez/wize-academy-go/common"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type pokemonCSVReader struct {
	filePath string
	// pokemons map[int]model.Pokemon
}

func NewPokemonCSVReader(csvPath string) repository.PokemonRepository {
	pkmnCSV := &pokemonCSVReader{
		filePath: csvPath,
		// pokemons: make(map[int]model.Pokemon),
	}

	return pkmnCSV
}

func (pkmnCSV *pokemonCSVReader) Get(id int) (model.Pokemon, error) {
	csvLines, err := common.ReadCSV(pkmnCSV.filePath)
	if err != nil {
		fmt.Println(err)
		return model.NullPokemon(), err
	}

	for _, line := range csvLines {
		csvID := line[0]
		csvName := line[1]
		pkmnId, err := strconv.Atoi(csvID)
		if err != nil {
			fmt.Println(err)
			return model.NullPokemon(), err
		}

		if pkmnId == id {
			pkmn := model.Pokemon{
				ID:   pkmnId,
				Name: csvName,
			}

			return pkmn, nil
		}
	}

	return model.NullPokemon(), repository.ErrorKeyNotFound
}
