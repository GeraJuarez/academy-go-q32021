# Wizeline academy Go final project

## API DOC

URL `http://localhost:8080/api`

## Hello Service

`GET` Hello
The endpoint `GET /hello` is the health check service for the API

## Pokemon Service V1

### `GET` Pokemon
The endpoint `GET /v1/pokemon/{id}` returns a pokemon with the specified ID located in the CSV resource file.

### `GET` Pokemon
The endpoint `GET /v1/pokemon` returns all requested pokemons by quantity and type using concurrency

#### Query Params
* type: `odd` or `even` to select pokemons with odd IDs or even IDs respectively
* items: the number of pokemons to fetch
* items_per_workers: the max number of pokemon a worket can processes

### `POST` Pokemon
The endpoint `POST /v1/pokemon/{id}` returns a pokemon with the specified from the pokeAPI service (https://pokeapi.co/) and saves it into the CSV resource file.



## Notes

* The interactor layer is more usually called Service or Usecase; I was following manakuro's clean archicture code, so I called it Interactor.
