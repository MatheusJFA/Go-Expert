package valueObjects

import (
	"regexp"
)

const (
	symbolsRegex = "[^0-9]"
)

type CPF string

func sameNumber(cpf string) bool {
	for i := 0; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}

	return true
}

func (value CPF) ValidateCPF() bool {
	cpf := regexp.MustCompile(symbolsRegex).ReplaceAllString(string(value), "")

	if len(cpf) != 11 || sameNumber(cpf) || cpf == "" {
		return false
	}

	sum := 0

	digits := cpf[9:]

	for i := 1; i <= 10; i++ {
		sum += int(cpf[i-1]-'0') * (11 - i)
	}

	rest := (sum * 10) % 11
	if rest == 10 || rest == 11 {
		rest = 0
	}

	if rest != int(digits[0]-'0') {
		return false
	}

	sum = 0

	for i := 1; i <= 10; i++ {
		sum += int(cpf[i-1]-'0') * (12 - i)
	}

	rest = (sum * 10) % 11

	if rest == 10 || rest == 11 {
		rest = 0
	}

	if rest != int(digits[1]-'0') {
		return false
	}

	return true
}
