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

	pkmn.HandleFunc("/{id}", c.PokemonController.GetValue).Methods(http.MethodGet)
}
