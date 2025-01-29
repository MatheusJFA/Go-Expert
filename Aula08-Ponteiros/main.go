package main

import "fmt"

type Client struct {
	nome  string
	saldo float64
}

func (c Client) Deposit(valor float64) {
	println("----------------")
	fmt.Printf("O saldo anterior de %v é de: R$%.2f \n", c.nome, c.saldo)
	c.saldo += valor
	fmt.Printf("O saldo atual de %v é de: R$%.2f \n", c.nome, c.saldo)
}

func (c *Client) DepositByReference(valor float64) {
	println("----------------")
	fmt.Printf("O saldo anterior de %v é de: R$%.2f \n", c.nome, c.saldo)
	c.saldo += valor
	fmt.Printf("O saldo atual de %v é de: R$%.2f \n", c.nome, c.saldo)
}

// Função para criar um cliente, que retorna um ponteiro para o cliente
func CreateClient(nome string) *Client {
	return &Client{nome: nome, saldo: 0}
}

func main() {
	// Ponteiros são variáveis que armazenam o endereço de memória de outra variável
	// Ou seja: Variável => Ponteiro => Endereço de memória => Valor da Variável

	// Declaração de variável
	a := 10

	println("O valor da variável é de: ", a) // Aqui será impresso o valor da variável a que no caso é 10

	// Declaração de ponteiro
	var ponteiro *int = &a

	println("Essa variável está apontando para: ", ponteiro) // Aqui será impresso o endereço de memória da variável a

	// Valor da variável que o ponteiro está apontando
	// O nome que damos para essa operação é desreferenciar
	println("O Valor que consta nesse endereço de memória é: ", *ponteiro) // Aqui será impresso o valor da variável a que no caso é 10

	// Alterando o valor da variável a através do ponteiro
	*ponteiro = 50

	println("O valor da variável é de: ", a) // Aqui será impresso o valor da variável a que no caso é 50

	// Ao chamarmos uma função, o que é passado é uma cópia do valor da variável e não a variável em si
	// Para passar a variável em si, devemos passar o endereço de memória dela

	primeiroValor := 10
	segundoValor := 20

	println("----------------------------------------------------------------------------------")
	result := Add(primeiroValor, segundoValor)

	println("O resultado da soma é: ", result)                      // Aqui será impresso o resultado da soma que é 30
	println("O valor da variável primeiroValor é: ", primeiroValor) // Aqui será impresso o valor da variável primeiroValor que é 10
	println("O valor da variável segundoValor é: ", segundoValor)   // Aqui será impresso o valor da variável segundoValor que é 20

	println("----------------------------------------------------------------------------------")
	// Agora vamos utilizar a função Add2 e vermos o que acontece
	result2 := Add2(primeiroValor, segundoValor)
	println("O resultado da soma é: ", result2)                     // Aqui será impresso o resultado da soma que é 120
	println("O valor da variável primeiroValor é: ", primeiroValor) // Aqui será impresso o valor da variável primeiroValor que é 10
	println("O valor da variável segundoValor é: ", segundoValor)   // Aqui será impresso o valor da variável segundoValor que é 20

	println("----------------------------------------------------------------------------------")

	// Agora vamos utilizar a função AddByReference e vermos o que acontece
	result3 := AddByReference(&primeiroValor, &segundoValor)
	println("O resultado da soma é: ", result3)                     // Aqui será impresso o resultado da soma que é 120
	println("O valor da variável primeiroValor é: ", primeiroValor) // Aqui será impresso o valor da variável primeiroValor que é 100
	println("O valor da variável segundoValor é: ", segundoValor)   // Aqui será impresso o valor da variável segundoValor que é 20

	println("----------------------------------------------------------------------------------")

	// Vamos criar um cliente e depositar um valor
	joao := CreateClient("João")
	joao.Deposit(100)
	joao.Deposit(400)
	joao.Deposit(500)

	println("----------------")
	fmt.Printf("O saldo do João é de: R$%.2f \n", joao.saldo) // Aqui será impresso o saldo do cliente que é 0, pois, a função Depositar não altera o saldo do cliente
	println("----------------")

	// Vamos criar um cliente e depositar um valor
	maria := CreateClient("Maria")
	maria.DepositByReference(100)
	maria.DepositByReference(100)
	maria.DepositByReference(400)
	maria.DepositByReference(500)

	println("----------------")
	fmt.Printf("O saldo da Maria é de: R$%.2f \n", maria.saldo) // Aqui será impresso o saldo do cliente que é 1100, pois, a função DepositarReferencia altera o saldo do cliente
	println("----------------")
}

func Add(a, b int) int {
	return a + b
}

func Add2(a, b int) int {
	a = 100
	return a + b
}

func AddByReference(a, b *int) int {
	*a = 100
	return *a + *b
}
