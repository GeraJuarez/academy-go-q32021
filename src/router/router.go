package router

import (
	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gorilla/mux"
)

func Start(c controller.AppController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	addPokemonRouter(c, api)

	return router
}

func addPokemonRouter(c controller.AppController, api *mux.Router) {
	pkmn := api.PathPrefix("/v1/pokemon").Subrouter()

	pkmn.HandleFunc("/{id}", c.PokemonController.GetValue).Methods("GET")
}
