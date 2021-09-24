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
	pokemons map[int]model.Pokemon
}

func NewPokemonCSVReader(csvPath string) repository.PokemonRepository {
	pkmnCSV := &pokemonCSVReader{
		filePath: csvPath,
		pokemons: make(map[int]model.Pokemon),
	}

	csvLines, err := common.ReadCSV(pkmnCSV.filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, line := range csvLines {
		csvID := line[0]
		csvName := line[1]
		id, err := strconv.Atoi(csvID)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		pkmn := model.Pokemon{
			ID:   id,
			Name: csvName,
		}

		pkmnCSV.pokemons[id] = pkmn
	}

	return pkmnCSV
}

func (pkmnCSV *pokemonCSVReader) Get(id int) (model.Pokemon, error) {
	pkmn, ok := pkmnCSV.pokemons[id]

	if !ok {
		return model.Pokemon{}, repository.ErrorKeyNotFound
	}

	return pkmn, nil
}
