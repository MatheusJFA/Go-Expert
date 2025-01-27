package main

import "fmt"

func main() {
	// Inicializando um map
	var dictionary = make(map[string]int)

	// Adicionando valores ao map
	dictionary["um"] = 1
	dictionary["dois"] = 2
	dictionary["três"] = 3
	dictionary["quatro"] = 4
	dictionary["cinco"] = 5

	fmt.Println("Valores no map:", dictionary)
	fmt.Println("Tamanho do map:", len(dictionary))

	fmt.Println("--------------------------------------------------")

	// Iterando sobre o map
	for k, v := range dictionary {
		fmt.Println("Chave:", k, "Valor:", v)
	}

	fmt.Println("--------------------------------------------------")

	// Deletando um valor do map
	delete(dictionary, "dois")

	fmt.Println("Valores no map:", dictionary)
	fmt.Println("Tamanho do map:", len(dictionary))

	fmt.Println("--------------------------------------------------")

	// Verificando se uma chave existe no map
	var exist = exists("quatro", dictionary)

	fmt.Println("A chave 'quatro' existe no map?", exist)

	fmt.Println("--------------------------------------------------")

	// Instanciando um map com valores
	var letras = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println("Valores no map:", letras)
	fmt.Println("Tamanho do map:", len(letras))

	fmt.Println("--------------------------------------------------")

	// Utilizando o comando make para criar um map
	// O make é utilizado para criar mapas, slices e channels
	var numeros = make(map[string]int)

	numeros["um"] = 1
	numeros["dois"] = 2
	numeros["três"] = 3

	fmt.Println("Valores no map:", numeros)
}

func exists(key string, dictionary map[string]int) bool {
	// _ é uma variável que ignora o valor retornado também chamado de blank identifier
	_, ok := dictionary[key]
	return ok
}
