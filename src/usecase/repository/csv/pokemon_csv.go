package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	csvUtils "github.com/gerajuarez/wize-academy-go/infrastructure/csv_utils"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

type pokemonCSVReader struct {
	filePath string
}

// NewPokemonCSVReader return a Pokemon repository that
// manages data from a CSV file
func NewPokemonCSVReader(csvPath string) repository.PokemonRepository {
	pkmnCSV := &pokemonCSVReader{
		filePath: csvPath,
	}

	return pkmnCSV
}

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

func (pkmnCSV *pokemonCSVReader) GetAllValid(items int, itemsPerWorker int, isValid func(id int) bool) ([]model.Pokemon, error) {
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
	source := make(chan model.Pokemon)
	dests := Split(source, totalWorkers)

	go func() {
		for {
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

	filterPkmn := make([]model.Pokemon, 0)
	for pkmn := range Funnel(itemsPerWorker, isValid, dests...) {
		filterPkmn = append(filterPkmn, pkmn)
	}

	return filterPkmn, nil
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
