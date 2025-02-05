package router

import (
	"net/http"

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
	hello.HandleFunc("", c.HelloController.HelloWorld).Methods(http.MethodGet)
}

func addPokemonRouter(c controller.AppController, api *mux.Router) {
	pkmn := api.PathPrefix("/v1/pokemon").Subrouter()

	pkmn.HandleFunc("/{id}", c.PokeCSV.GetValue).Methods(http.MethodGet)
	pkmn.HandleFunc("", c.PokeCSV.GetAll).
		Methods(http.MethodGet).
		Queries(
			"type", "{type:odd|even}",
			"items", "{items:[0-9]+}",
			"items_per_workers", "{items_per_workers:[0-9]+}")
	pkmn.HandleFunc("", c.PokeCSV.GetAllNoItems).
		Methods(http.MethodGet).
		Queries(
			"type", "{type:odd|even}",
			"items_per_workers", "{items_per_workers:[0-9]+}")
	pkmn.HandleFunc("/{id}", c.PokeCSV.PostByID).Methods(http.MethodPost)
}
