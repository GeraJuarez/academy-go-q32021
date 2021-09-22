package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/gorilla/mux"
)

type AppController struct {
	PokemonController PokemonController
}

type PokemonController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
}

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

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
	if errors.Is(err, interactor.ErrorKeyNotFound) {
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
