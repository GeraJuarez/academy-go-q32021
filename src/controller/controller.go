package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"

	"github.com/gorilla/mux"
)

type AppController struct {
	HelloController
	PokemonController
}

// HelloController implements the health check for the API
type HelloController interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}

type helloController struct{}

// NewHelloController creates a HelloController
func NewHelloController() HelloController {
	return &helloController{}
}

func (c *helloController) HelloWorld(w http.ResponseWriter, r *http.Request) {
	value := "Hello wizeline academy 2021."

	w.Write([]byte(value))
}

// PokemonController implements the interaction with the Pokemon resource
type PokemonController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
}

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

// NewPokemonController creates a PokemonController using the interactor inter
func NewPokemonController(inter interactor.PokemonInteractor) PokemonController {
	return &pokemonController{inter}
}

func (c *pokemonController) GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqId := vars["id"]

	id, err := strconv.Atoi(reqId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := c.pokemonInteractor.Get(id)

	if errors.Is(err, repository.ErrorKeyNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
