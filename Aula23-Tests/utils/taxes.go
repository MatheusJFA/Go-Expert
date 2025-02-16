package utils

import "errors"

type Repository interface {
	SaveTax(value float64) error
}

func CalculateTaxes(value float64) (float64, error) {
	// Definindo as faixas de imposto

	if value <= 0 {
		return 0, errors.New("value cannot be 0 or negative")
	} else if value <= 1000 {
		return value * 0.01, nil // 1%
	} else if value <= 2000 {
		return value * 0.05, nil // 5%
	} else if value <= 5000 {
		return value * 0.10, nil // 10%
	} else {
		return value * 0.15, nil // 15%
	}
}

func CalculateTaxAndSave(value float64, repository Repository) error {
	tax, err := CalculateTaxes(value)
	if err != nil {
		return err
	}
	return repository.SaveTax(tax)
}
