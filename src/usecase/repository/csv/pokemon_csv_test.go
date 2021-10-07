package repository

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gerajuarez/wize-academy-go/usecase/repository"

	"github.com/stretchr/testify/assert"
)

var pkmnCSVRepo = NewPokemonCSVReader("./resources/pokemons.csv")

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

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
