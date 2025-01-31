package main

func main() {
	var numeros [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var itens = map[string]int{
		"item1": 1,
		"item2": 2,
		"item3": 3,
		"item4": 4,
	}

	// Percorrendo um array for tradicional
	for i := 0; i < len(numeros); i++ {
		println(numeros[i])
	}

	println("-------------------------------------------------")

	// Percorrendo um array com range
	for index, valor := range numeros {
		println(index, ":", valor)
	}

	println("-------------------------------------------------")

	// Percorrendo um map
	for chave, valor := range itens {
		println(chave, ":", valor)
	}

	println("-------------------------------------------------")

	// Ignorando o índice (Blank Identifier)
	for _, valor := range numeros {
		println(valor)
	}

	println("-------------------------------------------------")

	// While
	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	// for {
	// 	println("Loop infinito")
	// }

}
