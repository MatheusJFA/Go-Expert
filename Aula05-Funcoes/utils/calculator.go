package utils

import (
	"errors"
)

// func -> função
// Add -> nome da função
// (a int, b int) -> parâmetros da função
// int -> tipo de retorno da função

// Se os arametros forem o mesmo, posso unir os parâmetros e colocar o tipo depois
// Exemplo: func Concat(a, b String) String
func Add(a, b int) int {
	return a + b
}

// Funções variádicas
func AddMany(numbers ...int) int {
	total := 0

	// _ fica sendo o index no "array"
	for _, valor := range numbers {
		total += valor
	}

	return total
}

// Se preferir posso também colocar cada parâmetro com seu tipo
// Exemplo: func Concat(a String, b String) String
func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

// Nesse caso, tenho dois retornos, um int ou um error
// nil é a mesma coisa de null
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}

	return a / b, nil
}
