package main

import (
	"GoExpert/utils"
	"fmt"
	"time"
)

// Declaração de constante (não pode ser alterada após a declaração)
const PI = 3.1415

// Declaração de variáveis globais
var (
	name       string    = "Matheus"
	birthdate  time.Time = time.Date(1995, time.April, 4, 0, 0, 0, 0, time.UTC)
	salary     float64   = 7_500.94
	isLearning bool      = true
)

func main() {
	// Declaração de variáveis locais
	girlfriend := "Pâmella"
	girlfriendBirthdate := time.Date(1998, time.December, 18, 0, 0, 0, 0, time.UTC)
	hasGirlfriend := utils.Ternary(len(girlfriend) > 0, "Yes", "No")

	learn := utils.Ternary(isLearning, "Yes", "No")
	age := utils.GetAge(birthdate)
	girlfriendAge := utils.GetAge(girlfriendBirthdate)

	fmt.Println("Hello, World!") // Não é bom esquecer disso :D
	fmt.Printf("I'm %s, I'm %d years old, I earn %.2f and I'm learning? %s \n", name, age, salary, learn)
	fmt.Printf("I have a girlfriend? %s, her name is %s and she is %d years old \n", hasGirlfriend, girlfriend, girlfriendAge)
	fmt.Printf("The value of PI is %.4f \n", PI)
}
