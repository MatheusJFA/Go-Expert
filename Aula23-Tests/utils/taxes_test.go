package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
* Para rodar os testes, execute o comando:
		go test
* Para rodar os testes e ver o output de cada teste, execute o comando:
		go test -v
* Para rodar um teste e verificar o coverage do código, execute o comando:
		go test -coverprofile=coverage.out
		go tool cover -html=coverage.out
* Para rodar os testes de benchmarking, execute o comando:
		go test -bench .
* Para rodar fuzzy testing, execute o comando:
		go test -fuzz .
*/

func Test_CalculateTaxes(t *testing.T) {
	type Result struct {
		value    float64
		expected float64
	}

	results := []Result{
		{value: -1000, expected: 0},
		{value: 0, expected: 0},
		{value: 1000, expected: 50},
		{value: 3000, expected: 300},
		{value: 6000, expected: 900},
		{value: 10000, expected: 1500},
		{value: 20000, expected: 3000},
	}

	for _, result := range results {
		taxes, err := CalculateTaxes(result.value)
		if err != nil {
			assert.Error(t, err)
		}
		if taxes != result.expected {
			assert.NotEqual(t, "Expected %f, got %f", result.expected, taxes)
		}
	}
}

func Benchmark_CalculateTaxes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTaxes(1000)
	}
}

func Fuzz_CalculateTaxes(f *testing.F) {
	seed := []float64{-1000, -250, -100, 0, 250, 1000}

	for _, value := range seed {
		f.Add(value)
	}

	f.Fuzz(func(t *testing.T, value float64) {
		taxes, err := CalculateTaxes(value)
		if err != nil {
			assert.Error(t, err)
		}

		if taxes == 0 {
			assert.Contains(t, err.Error(), "value cannot be 0 or negative")
		}

	})
}

type TaxRepositoryMock struct {
	mock.Mock
}

func (mock *TaxRepositoryMock) SaveTax(value float64) error {
	args := mock.Called(value)
	return args.Error(0)
}

func Test_CalculateTaxesAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 300.0).Return(nil)
	repository.On("SaveTax", 10.0).Return(nil)

	err := CalculateTaxAndSave(3000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0, repository)
	assert.Error(t, err, errors.New("value cannot be 0 or negative"))

	repository.AssertExpectations(t)
}
