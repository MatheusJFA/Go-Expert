package utils

import "time"

// Função que calcula a idade
func GetAge(birthdate time.Time) int {
	var currentYear = time.Now()
	return int(currentYear.Sub(birthdate).Hours() / 8760)
}
