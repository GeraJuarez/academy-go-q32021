# Wizeline academy Go final project

## API DOC

URL `http://localhost:8080/api`

### Hello Service

`GET` Hello
The endpoint `GET /hello` is the health check service for the API

### Pokemon Service V1

`GET` Pokemon
The endpoint `GET /v1/pokemon/{id}` returns a pokemon with the specified ID located in the CSV resource file.

### Pokemon Service V2

`GET` Pokemon
The endpoint `GET /v2/pokemon/{id}` returns a pokemon with the specified from the pokeAPI service (https://pokeapi.co/) and saves it into the CSV resource file.

### Pokemon Service V3

WIP

## Notes

* The interactor layer is more usually called Service or Usecase; I was following manakuro's clean archicture code, so I called it Interactor.
* The pokemon version 2 controller should be instead placed inside the version 1 controller but as POST method, the reason this is in another controler is because I wanted to separate the project deliveries in controllers.