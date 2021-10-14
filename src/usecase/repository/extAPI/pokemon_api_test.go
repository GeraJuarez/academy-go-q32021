package repository

import (
	"log"
	"os"
	"testing"

	"github.com/gerajuarez/wize-academy-go/usecase/repository"

	"github.com/stretchr/testify/assert"
)

const CSV_TEST_PATH = "./test.csv"

var pkmnCSVRepo = NewExtApiRepo(CSV_TEST_PATH)

func TestMain(m *testing.M) {
	code := m.Run()
	err := os.Remove(CSV_TEST_PATH)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
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
