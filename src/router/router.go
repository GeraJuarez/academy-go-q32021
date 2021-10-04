package router

import (
	"github.com/gerajuarez/wize-academy-go/controller"

	"github.com/gorilla/mux"
)

// Start initializes the API routing
func Start(c controller.AppController) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	addHelloRouter(c, api)
	addPokemonRouter(c, api)

	return router
}

func addHelloRouter(c controller.AppController, api *mux.Router) {
	hello := api.PathPrefix("/hello").Subrouter()
	hello.HandleFunc("", controller.Hello).Methods("GET")
}

func addPokemonRouter(c controller.AppController, api *mux.Router) {
	pkmn := api.PathPrefix("/v1/pokemon").Subrouter()

	pkmn.HandleFunc("/{id}", c.PokemonController.GetValue).Methods("GET")
}
