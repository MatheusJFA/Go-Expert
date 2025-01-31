package main

import "fmt"

type Number int

type Numeric interface {
	// ~ int -> Significa que podemos usar qualquer tipo de dado que seja int, incluindo o MyNumber, int, int8, int16, int32, int64
	~int | float32
}

func main() {
	// Declaração de interfaces vazias (aceitam qualquer tipo de dado e é a forma "primitiva" de generics em Go)
	var texto interface{} = "Matheus"
	var numero interface{} = 30

	fmt.Println("------------------------------")
	fmt.Println(texto.(string)) // Type assertion
	fmt.Println(numero.(int))   // Type assertion
	fmt.Println("------------------------------")

	// Convertendo um tipo de dado para outro
	fmt.Println("------------------------------")
	res, ok := texto.(float32)

	if !ok {
		fmt.Println("Erro ao converter o tipo de dado")
	} else {
		fmt.Println(res)
	}
	fmt.Println("------------------------------")

	DescribeType(texto)
	DescribeType(numero)
	fmt.Println("------------------------------")

	workers := map[string]int{
		"Matheus": 10_000,
		"Pâmella": 15_000,
		"Lucas":   5_000,
		"Gerson":  7_000,
		"Caio":    8_000,
	}

	employees := map[string]float32{
		"Matheus": 10_000.50,
		"Pâmella": 15_000.75,
		"Lucas":   5_000.25,
		"Gerson":  7_000.75,
		"Caio":    8_000.25,
	}

	cooworkers := map[string]Number{
		"Matheus": 10,
		"Pâmella": 15,
		"Lucas":   5,
		"Gerson":  7,
		"Caio":    8,
	}

	fmt.Printf("O somatório dos salários usando o método {Add} é R$ %v \n", Add(workers))
	fmt.Println("------------------------------")

	fmt.Printf("O somatório dos salários usando o método {AddGeneric} recebendo o worker é R$ %v e recebendo o employees R$ %v \n", AddGeneric(workers), AddGeneric(employees))
	fmt.Println("------------------------------")

	fmt.Printf("O somatório dos salários usando o método {AddWithConstraints} recebendo o worker é R$ %v e recebendo o employees R$ %v \n", AddWithConstraints(workers), AddWithConstraints(employees))
	fmt.Println("------------------------------")

	fmt.Printf("O somatório dos salários usando o método {AddWithConstraints} recebendo o cooworkers R$ %v \n", AddWithConstraints(cooworkers))
	fmt.Println("------------------------------")

	fmt.Printf("O resultado da comparação entre 10 e 10 é %v \n", Compare(10, 10))
	fmt.Printf("O resultado da comparação entre 10 e 20 é %v \n", Compare(10, 20.0))
	// fmt.Printf("O resultado da comparação entre 10 e 20 é %v \n", Compare(10, "20")) // Erro de compilação

}

func Compare[T comparable](a T, b T) bool {
	return a == b
}

func Add(list map[string]int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum
}

func AddGeneric[T int | float32](list map[string]T) T {
	var sum T

	for _, v := range list {
		sum += v
	}

	return sum
}

func AddWithConstraints[T Numeric](list map[string]T) T {
	var sum T

	for _, v := range list {
		sum += v
	}

	return sum
}

func DescribeType(i interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v \n", i, i)
}
