package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	csvUtils "github.com/gerajuarez/wize-academy-go/infrastructure/csv_utils"
	pokeAPI "github.com/gerajuarez/wize-academy-go/infrastructure/poke_api"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type pokemonCSVReader struct {
	filePath   string
	pokeClient pokeAPI.PokeAPIClient
}

// NewPokemonCSVReader return a Pokemon repository that
// manages data from a CSV file
func NewPokemonCSVReader(csvPath string, httpClient pokeAPI.PokeAPIClient) repository.PokemonRepository {
	pkmnCSV := &pokemonCSVReader{
		filePath:   csvPath,
		pokeClient: httpClient,
	}

	return pkmnCSV
}

// Get returns a pokemon from the CSV fil resource that matches the ID
func (pkmnCSV *pokemonCSVReader) Get(id int) (model.Pokemon, error) {
	csvLines, err := csvUtils.ReadCSV(pkmnCSV.filePath)
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

// GetAllValid returns a slice of pokemons which id is validated by the given function
// items are the maximum lenght of the slice to be returned, if the CSV does not have enough records, the number may be lower than expected.
// itemsPerWorker determines the concurrency level
func (pkmnCSV *pokemonCSVReader) GetAllValid(items int, itemsPerWorker int, isValid func(id int) bool) ([]model.Pokemon, error) {
	if items == repository.ALL_ITEM_QUERY {
		fileCount, err := csvUtils.CountCSVLines(pkmnCSV.filePath)
		if err != nil {
			return []model.Pokemon{}, err
		}
		items = fileCount
	}
	if items <= 0 {
		return []model.Pokemon{}, repository.ErrorItemZeroParam
	}
	if itemsPerWorker <= 0 {
		return []model.Pokemon{}, repository.ErrorWorkerZeroParam
	}
	if items < itemsPerWorker {
		itemsPerWorker = items
	}

	f, err := os.Open(pkmnCSV.filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	totalWorkers := items / itemsPerWorker
	totalWorkers += items % itemsPerWorker
	source := make(chan model.Pokemon)
	dests := Split(source, totalWorkers)
	filterPkmn := make([]model.Pokemon, 0)

	go func() {
		for {
			if len(filterPkmn) >= items {
				break
			}

			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				//return nil, err
			}

			csvID := line[0]
			csvName := line[1]
			pkmnId, err := strconv.Atoi(csvID)
			if err != nil {
				//return nil, err
			}
			source <- model.Pokemon{ID: pkmnId, Name: csvName}

		}

		close(source)
	}()

	for pkmn := range Funnel(itemsPerWorker, isValid, dests...) {
		filterPkmn = append(filterPkmn, pkmn)
		if len(filterPkmn) == items {
			break
		}
	}

	return filterPkmn, nil
}

// PostById calls an external source to fetch the pokemon with the specified ID and it is added to the CSV resource file.
func (pkmnCSV *pokemonCSVReader) PostById(id int) (model.Pokemon, error) {
	body, statusCode, err := pkmnCSV.pokeClient.GetPokemonByID(id)
	if err != nil {
		return model.NullPokemon(), err
	}

	switch statusCode {
	case http.StatusOK:
		var pkmn model.Pokemon
		json.Unmarshal(body, &pkmn)
		return pkmnCSV.Post(pkmn)
	case http.StatusNotFound:
		return model.NullPokemon(), repository.ErrorKeyNotFound
	default:
		return model.NullPokemon(), fmt.Errorf("PokeAPI Error: %s", string(body))
	}
}

// Post adds a pokemon to the CSV resource file.
func (pkmnCSV *pokemonCSVReader) Post(pkmn model.Pokemon) (model.Pokemon, error) {
	var data [][]string
	row := []string{strconv.Itoa(pkmn.ID), pkmn.Name}
	data = append(data, row)

	err := csvUtils.AppendCSV(pkmnCSV.filePath, data)
	if err != nil {
		return model.NullPokemon(), err
	}

	return pkmn, nil
}

// Split implements a Fan-Out pattern that splits the channel into multiple pnes
func Split(source <-chan model.Pokemon, n int) []<-chan model.Pokemon {
	dests := make([]<-chan model.Pokemon, 0)

	for i := 0; i < n; i++ {
		ch := make(chan model.Pokemon)
		dests = append(dests, ch)

		go func() {
			defer close(ch)
			for val := range source {
				ch <- val
			}
		}()
	}

	return dests
}

// Funnel implements a Fan-In pattern and filters by ID type
func Funnel(maxItems int, isValid func(id int) bool, sources ...<-chan model.Pokemon) <-chan model.Pokemon {
	dest := make(chan model.Pokemon)
	var wg sync.WaitGroup
	wg.Add(len(sources))

	for id, ch := range sources {
		go func(c <-chan model.Pokemon, idx int) {
			count := 0
			defer wg.Done()
			for n := range c {
				if isValid(n.ID) {
					dest <- n
					count++
				}
				if count == maxItems {
					break
				}
			}
		}(ch, id)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}
