package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type pokemonCSVReader struct {
	filePath string
	pokemons map[int]model.Pokemon
}

func NewPokemonCSVReader() repository.PokemonRepository {
	pkmnCSV := &pokemonCSVReader{filePath: "./resources/pokemons.csv"}

	csvLines, err := ReadCSV(pkmnCSV.filePath)
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
		return model.Pokemon{}, repository.ErrorNoSuchKey
	}

	return pkmn, nil
}

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
