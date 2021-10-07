package controller_test

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"
	"github.com/stretchr/testify/mock"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

type MockPockemonInteractor struct {
	mock.Mock
}

func (m *MockPockemonInteractor) Get(id int) (model.Pokemon, error) {
	args := m.Called(id)
	return args.Get(0).(model.Pokemon), args.Error(1)
}

func TestControllerStatus(t *testing.T) {
	cases := []struct {
		testName       string
		request        string
		httpMethod     string
		route          string
		handler        func(http.ResponseWriter, *http.Request)
		expectedStatus int
	}{
		{
			"hello Controller OK",
			"/hello",
			http.MethodGet,
			"/hello",
			controller.NewHelloController().HelloWorld,
			http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			req, err := http.NewRequest(c.httpMethod, c.request, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := mux.NewRouter()
			handler.HandleFunc(c.route, c.handler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != c.expectedStatus {
				t.Errorf("Expected %d,got %d", c.expectedStatus, status)
			}
		})
	}
}

func TestPkmnController_GetValue(t *testing.T) {
	cases := []struct {
		testName       string
		request        string
		httpMethod     string
		route          string
		expectedStatus int
		err            error
	}{
		{
			"Pokemon CSV Controller OK",
			"/v1/pokemon/1",
			http.MethodGet,
			"/v1/pokemon/{id}",
			http.StatusOK,
			nil,
		},
		{
			"Pokemon CSV Controller Not found",
			"/v1/pokemon/1",
			http.MethodGet,
			"/v1/pokemon/{id}",
			http.StatusNotFound,
			repository.ErrorKeyNotFound,
		},
		{
			"Pokemon CSV Controller Server Error",
			"/v1/pokemon/1",
			http.MethodGet,
			"/v1/pokemon/{id}",
			http.StatusInternalServerError,
			errors.New("test error"),
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			req, err := http.NewRequest(c.httpMethod, c.request, nil)
			if err != nil {
				t.Fatal(err)
			}

			mock := new(MockPockemonInteractor)
			mock.On("Get", 1).Return(model.NullPokemon(), c.err)

			rr := httptest.NewRecorder()
			handler := mux.NewRouter()
			handler.HandleFunc(c.route, controller.NewPokemonController(mock).GetValue)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != c.expectedStatus {
				t.Errorf("Expected %d,got %d", c.expectedStatus, status)
			}
		})
	}

}
