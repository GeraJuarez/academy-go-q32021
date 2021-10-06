package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gerajuarez/wize-academy-go/model"
)

const BASE_URL_PKMN = "http://pokeapi.co/api/v2/pokemon"

type pokeAPIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewPokeAPIClient return a client that calls pokeAPI
func NewPokeAPIClient() *pokeAPIClient {
	return &pokeAPIClient{
		BaseURL:    BASE_URL_PKMN,
		HTTPClient: &http.Client{Timeout: time.Minute},
	}
}

func (c *pokeAPIClient) GetPokemonByID(id int) (model.Pokemon, error) {
	res, err := http.Get(fmt.Sprintf("%s/%d", c.BaseURL, id))
	if err != nil {
		return model.NullPokemon(), err
	}

	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return model.NullPokemon(), fmt.Errorf("error requesting pokeAPI, status code: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.NullPokemon(), err
	}

	var pokemon model.Pokemon
	json.Unmarshal(body, &pokemon)
	return pokemon, nil

}
