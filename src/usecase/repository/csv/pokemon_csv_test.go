package repository

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gerajuarez/wize-academy-go/usecase/repository"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestCSVReadOK(t *testing.T) {
	pkmnCSVRepo := NewPokemonCSVReader()
	expected := "bulbasaur"
	result, _ := pkmnCSVRepo.Get(1)

	if expected != result.Name {
		t.Errorf("wrong type: got %v want %v", result, expected)
	}
}

func TestCSVReadErr(t *testing.T) {
	pkmnCSVRepo := NewPokemonCSVReader()
	expected := repository.ErrorNoSuchKey
	_, err := pkmnCSVRepo.Get(999)

	if !errors.Is(expected, err) {
		t.Errorf("wrong type: got %v want %v", err, expected)
	}
}
