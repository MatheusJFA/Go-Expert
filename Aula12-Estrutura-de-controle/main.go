package main

func main() {
	primeiroNumero := 10
	segundoNumero := 20

	// Comparações em Go
	// == -> Igual
	// != -> Diferente
	// > -> Maior
	//  < -> Menor
	// >= -> Maior ou igual
	// <= -> Menor ou igual

	// Operadores lógicos
	// && -> E
	// || -> OU
	// ! -> Negação

	// Operador ternário
	// Não existe operador ternário em Go

	// If padrão
	if primeiroNumero > segundoNumero {
		println("O primeiro número é maior que o segundo")
	} else if primeiroNumero < segundoNumero {
		println("O primeiro número é menor que o segundo")
	} else {
		println("Os números são iguais")
	}

	// Switch

	// O switch em Go não precisa de um break para cada case
	println("-------------------------------------------------")
	switch {
	case primeiroNumero > segundoNumero:
		println("O primeiro número é maior que o segundo")
	case primeiroNumero < segundoNumero:
		println("O primeiro número é menor que o segundo")
	default:
		println("Os números são iguais")
	}

	// Switch com variável
	println("-------------------------------------------------")
	switch primeiroNumero {
	case 10:
		println("O número é 10")
	case 20:
		println("O número é 20")
	default:
		println("O número não é 10 nem 20")
	}
}
