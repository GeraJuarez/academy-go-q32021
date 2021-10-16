package repository

import (
	"testing"

	"github.com/gerajuarez/wize-academy-go/usecase/repository"

	"github.com/stretchr/testify/assert"
)

var pkmnCSVRepo = NewPokemonCSVReader("./resources/pokemons.csv")

func TestCSVGet(t *testing.T) {
	cases := []struct {
		testName string
		id       int
		err      error
	}{
		{
			"PkmnCSV Repo OK",
			1,
			nil,
		},
		{
			"PkmnCSV Repo NotFound",
			-1,
			repository.ErrorKeyNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			_, err := pkmnCSVRepo.Get(c.id)
			assert.Equal(t, c.err, err, "Error should be equal")

		})
	}
}

func TestCSVGetAll(t *testing.T) {
	cases := []struct {
		testName       string
		items          int
		itemsPerWorker int
		validation     func(n int) bool
		expected       int
		err            error
	}{
		{
			"Even ID pokemons",
			10,
			5,
			repository.IsEven,
			10,
			nil,
		},
		{
			"Odd ID pokemons",
			10,
			5,
			repository.IsOdd,
			10,
			nil,
		},
		{
			"Odd ID pokemons EOF, 900 total, 450 odd pkmns",
			900,
			10,
			repository.IsOdd,
			450,
			nil,
		},
		{
			"Odd ID pokemons EOF, same workers, same items",
			900,
			900,
			repository.IsOdd,
			450,
			nil,
		},
		{
			"Odd ID pokemons, same workers, same items",
			200,
			200,
			repository.IsOdd,
			200,
			nil,
		},
		{
			"Error zero items",
			0,
			200,
			repository.IsOdd,
			0,
			repository.ErrorItemZeroParam,
		},
		{
			"Error zero items per worker",
			1,
			0,
			repository.IsOdd,
			0,
			repository.ErrorWorkerZeroParam,
		},
		{
			"Odd ID, items per worker larger",
			5,
			10,
			repository.IsOdd,
			5,
			nil,
		},
		{
			"Items not divisible per worker",
			5,
			2,
			repository.IsOdd,
			5,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			pkmns, err := pkmnCSVRepo.GetAllValid(c.items, c.itemsPerWorker, c.validation)
			got := len(pkmns)
			assert.Equal(t, c.expected, got, "lens should be equal")
			for _, poke := range pkmns {
				assert.True(t, c.validation(poke.ID), "poke ID %d no valid", poke.ID)
			}
			assert.Equal(t, c.err, err, "Error should be equal")
		})
	}
}
