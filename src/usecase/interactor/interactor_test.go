package interactor_test

import (
	"testing"

	"github.com/gerajuarez/wize-academy-go/model"
	"github.com/gerajuarez/wize-academy-go/usecase/interactor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPockemonRepo struct {
	mock.Mock
}

func (m *MockPockemonRepo) Get(id int) (model.Pokemon, error) {
	args := m.Called(id)
	return args.Get(0).(model.Pokemon), args.Error(1)
}

func (m *MockPockemonRepo) GetAllValid(items int, itemsPerWorker int, isValid func(id int) bool) ([]model.Pokemon, error) {
	args := m.Called(items, itemsPerWorker, isValid)
	return args.Get(0).([]model.Pokemon), args.Error(1)
}

func TestPkmnController_GetAllValid(t *testing.T) {
	cases := []struct {
		testName     string
		typeParam    string
		itemsParam   int
		workerParams int
		err          error
	}{
		{
			"Pokemon Interactor Odd param OK",
			"odd",
			1,
			1,
			nil,
		},
		{
			"Pokemon Interactor Even param OK",
			"even",
			1,
			1,
			nil,
		},
		{
			"Pokemon Interactor invalid param",
			"not valid",
			1,
			1,
			interactor.ErrorInvalidTypeParam,
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			mockRepo := new(MockPockemonRepo)
			mockRepo.On("GetAllValid", c.itemsParam, c.workerParams, mock.AnythingOfType("func(int) bool")).Return([]model.Pokemon{}, c.err)
			pokeInter := interactor.NewPokemonInteractor(mockRepo)
			_, err := pokeInter.GetAllByType(c.typeParam, c.itemsParam, c.workerParams)
			assert.Equal(t, c.err, err, "Error should be equal")
		})
	}

}
