package utils

import (
	"strconv"
	"strings"
)

func VerificaModulo11(valor string) bool {
	valor = padLeft(valor, 11, "0")

	var soma int = 0
	var peso int = 2

	if len(valor) != 11 || onlyZero(valor) {
		return false
	}

	ultimoDigito, err := strconv.Atoi(valor[len(valor)-1:])

	if err != nil {
		return false
	}

	for i := len(valor) - 2; i >= 0; i-- {
		soma += int(valor[i]-'0') * peso

		maiorQueNove := peso > 9

		if maiorQueNove {
			peso = 2
		} else {
			peso++
		}
	}

	resultado := soma % 11

	var digitoVerificador int

	if resultado == 0 || resultado == 1 {
		digitoVerificador = 0
	} else {
		digitoVerificador = 11 - resultado
	}

	return digitoVerificador == ultimoDigito
}

func padLeft(s string, n int, c string) string {
	s = strings.Repeat(c, n) + s
	return s[len(s)-n:]
}

func onlyZero(s string) bool {
	for _, v := range s {
		if v != '0' {
			return false
		}
	}
	return true
}
