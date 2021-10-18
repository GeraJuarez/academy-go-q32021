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

// PokemonController implements the interaction with the Pokemon resource
type PokemonController interface {
	GetValue(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllNoItems(w http.ResponseWriter, r *http.Request)
	PostByID(w http.ResponseWriter, r *http.Request)
}

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

// NewPokemonController creates a PokemonController using the interactor inter
func NewPokemonController(inter interactor.PokemonInteractor) PokemonController {
	return &pokemonController{inter}
}

// GetValue
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

// GetAll handles the GET pokemons endpoint, it reads the query params items, type, and items_per_worker
func (c *pokemonController) GetAll(w http.ResponseWriter, r *http.Request) {
	paramType := r.URL.Query().Get("type")
	paramItems := r.URL.Query().Get("items")
	paramPerWorker := r.URL.Query().Get("items_per_workers")

	intItems, err := strconv.Atoi(paramItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intPerWorker, err := strconv.Atoi(paramPerWorker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	values, err := c.pokemonInteractor.GetItemsByType(paramType, intItems, intPerWorker)
	if errors.Is(err, repository.ErrorItemZeroParam) || errors.Is(err, repository.ErrorWorkerZeroParam) || errors.Is(err, interactor.ErrorInvalidTypeParam) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAll handles the GET pokemons endpoint, it reads the query params type, and items_per_worker
func (c *pokemonController) GetAllNoItems(w http.ResponseWriter, r *http.Request) {
	paramType := r.URL.Query().Get("type")
	paramPerWorker := r.URL.Query().Get("items_per_workers")

	intPerWorker, err := strconv.Atoi(paramPerWorker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	values, err := c.pokemonInteractor.GetAllByType(paramType, intPerWorker)
	if errors.Is(err, repository.ErrorItemZeroParam) || errors.Is(err, repository.ErrorWorkerZeroParam) || errors.Is(err, interactor.ErrorInvalidTypeParam) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PostByID handles the POST pokemon using an ID
func (c *pokemonController) PostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqId := vars["id"]

	id, err := strconv.Atoi(reqId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := c.pokemonInteractor.PostById(id)

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
