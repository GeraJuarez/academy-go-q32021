package controller

import "net/http"

// HelloController implements the health check for the API
type HelloController interface {
	HelloWorld(w http.ResponseWriter, r *http.Request)
}

type helloController struct{}

// NewHelloController creates a HelloController
func NewHelloController() HelloController {
	return &helloController{}
}

// HelloWorld handles the health check endpoint
func (c *helloController) HelloWorld(w http.ResponseWriter, r *http.Request) {
	value := "Hello wizeline academy 2021."

	w.Write([]byte(value))
}
