package main

import (
	"GoExpert/utils"
	"fmt"
)

func main() {
	var a = 10
	var b = 20
	var c = 0

	// Closures (Função Anônima)
	valor := func() int {
		return utils.Add(5, 1)
	}()

	fmt.Println("O valor da função anônima é: ", valor)

	fmt.Println("Adição de", a, "com", b, "é igual a:", utils.Add(a, b))
	fmt.Println("Subtração de", a, "com", b, "é igual a:", utils.Subtract(a, b))
	fmt.Println("Multiplicação de", a, "com", b, "é igual a:", utils.Multiply(a, b))

	var _, err = utils.Divide(a, c)

	var result, _ = utils.Divide(a, b)

	fmt.Println("Divisão de", a, "com", b, "é igual a: ", result)
	fmt.Println("Divisão de", a, "com", c, "é igual a: ", err)

	fmt.Println("-------------------------------------------------------")

	fmt.Println("O resultado do somatório é:", utils.AddMany(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 44))
}
