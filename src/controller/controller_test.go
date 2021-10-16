package controller_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gerajuarez/wize-academy-go/controller"
	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/repository"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

const (
	HELLO_PATH = "/hello"
	PKMN_PATH  = "/v1/pokemon"
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

func (m *MockPockemonInteractor) GetAllByType(typeStr string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	args := m.Called(typeStr, items, itemsPerWorker)
	return args.Get(0).([]model.Pokemon), args.Error(1)
}

func (m *MockPockemonInteractor) PostById(id int) (model.Pokemon, error) {
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
			HELLO_PATH,
			http.MethodGet,
			HELLO_PATH,
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
			PKMN_PATH + "/1",
			http.MethodGet,
			PKMN_PATH + "/{id}",
			http.StatusOK,
			nil,
		},
		{
			"Pokemon CSV Controller Not found",
			PKMN_PATH + "/1",
			http.MethodGet,
			PKMN_PATH + "/{id}",
			http.StatusNotFound,
			repository.ErrorKeyNotFound,
		},
		{
			"Pokemon CSV Controller Server Error",
			PKMN_PATH + "/1",
			http.MethodGet,
			PKMN_PATH + "/{id}",
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

func TestPkmnController_GetAll(t *testing.T) {
	cases := []struct {
		testName       string
		request        string
		typeParam      string
		itemsParam     string
		workerParams   string
		expectedStatus int
		err            error
	}{
		{
			"Pokemon CSV Controller OK",
			PKMN_PATH,
			"odd",
			"1",
			"1",
			http.StatusOK,
			nil,
		},
		{
			"Pokemon CSV Bad request",
			PKMN_PATH,
			"odd",
			"0",
			"1",
			http.StatusBadRequest,
			repository.ErrorItemZeroParam,
		},
		{
			"Pokemon CSV unknown error",
			PKMN_PATH,
			"odd",
			"0",
			"1",
			http.StatusInternalServerError,
			errors.New("new error"),
		},
		{
			"Pokemon CSV invalid items param",
			PKMN_PATH,
			"odd",
			"o",
			"1",
			http.StatusBadRequest,
			nil,
		},
		{
			"Pokemon CSV invalid worker param",
			PKMN_PATH,
			"odd",
			"1",
			"two",
			http.StatusBadRequest,
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?type=%s&items=%s&items_per_workers=%s", c.request, c.typeParam, c.itemsParam, c.workerParams), nil)
			if err != nil {
				t.Fatal(err)
			}

			mockInter := new(MockPockemonInteractor)
			mockInter.On("GetAllByType", c.typeParam, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]model.Pokemon{}, c.err)

			rr := httptest.NewRecorder()
			handler := mux.NewRouter()
			handler.HandleFunc(c.request, controller.NewPokemonController(mockInter).GetAll)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != c.expectedStatus {
				t.Errorf("Expected %d,got %d", c.expectedStatus, status)
			}
		})
	}

}
