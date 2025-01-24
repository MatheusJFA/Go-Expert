package main

import "fmt"

func main() {
	// Arrays

	// Array estático
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// array[5] = 6 Error: index out of range [5] with length 5

	fmt.Println("Valores no array:", array)
	fmt.Println("Tamanho do array:", len(array))
	fmt.Println("Capacidade do array:", cap(array))

	fmt.Println("Primeiro valor do array:", array[0])
	fmt.Println("Último valor do array:", array[len(array)-1])

	fmt.Println("--------------------------------------------------")

	// Iterando sobre o array
	for i := 0; i < len(array); i++ {
		fmt.Println("Valor do array na posição", i+1, ":", array[i])
	}

	fmt.Println("--------------------------------------------------")

	// Invertendo o array com o range
	for i, v := range array {
		fmt.Println("Valor do array na posição", len(array)-i, ":", v)
	}

	fmt.Println("--------------------------------------------------")

	// Slices
	// [x:y] -> x é o índice inicial e y é o índice final
	// [x:] -> x é o índice inicial e o final é o tamanho do array
	// [:y] -> o índice inicial é 0 e o final é y
	// [:] -> o índice inicial é 0 e o final é o tamanho do array

	fmt.Println("Array completo:", array[:])
	fmt.Println("Posições intermediárias do Array :", array[1:4])
	fmt.Println("Primeiras 3 posições do Array :", array[:3])
	fmt.Println("Últimas 2 posições do Array :", array[3:])

	fmt.Println("--------------------------------------------------")

	// Aumentando o tamanho do array
	fmt.Println("Adicionando um novo valor ao array")

	// Array dinâmico
	a := []int{10, 20, 30, 40, 50}

	fmt.Println("Array completo:", a)

	// Adicionando um novo valor ao array
	a = append(a, 60)
	a = append(a, 80)

	fmt.Println("Novo array completo:", a)

	fmt.Println("--------------------------------------------------")

}
