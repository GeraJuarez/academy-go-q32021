package main

import (
	"log"
	"net/http"
	"time"

	pokeAPI "github.com/gerajuarez/wize-academy-go/infrastructure/poke_api"
	"github.com/gerajuarez/wize-academy-go/registry"
	"github.com/gerajuarez/wize-academy-go/router"
)

const (
	PORT          = "8080"
	PATH_PKMN_CSV = "./usecase/repository/csv/resources/pokemons.csv"
)

func main() {
	portEnv := PORT

	extAPI := pokeAPI.NewPokeAPIClient()
	registry := registry.NewRegistry(PATH_PKMN_CSV, extAPI)
	router := router.Start(registry.NewAppController())

	srv := &http.Server{
		Addr:         ":" + portEnv,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen:%+s\n", err)
	}
}
