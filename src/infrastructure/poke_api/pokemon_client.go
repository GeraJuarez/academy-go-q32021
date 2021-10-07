package pokeAPI

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// GetPokemonByID requests a GetPokemon by ID from pokeAPI
// and retuns a body, response http status, and error
func (c *pokeAPIClient) GetPokemonByID(id int) ([]byte, int, error) {
	res, err := http.Get(fmt.Sprintf("%s/%d", c.BaseURL, id))
	if err != nil {
		return nil, res.StatusCode, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}
