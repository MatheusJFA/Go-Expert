package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Address struct {
	Country string
	City    string
}

type Base struct {
	ID         int
	Active     bool
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type User struct {
	Audit     Base // Criando Variável
	Name      string
	Email     string
	Password  []byte
	Birthdate time.Time
	Address   // Compondo a struct
}

// Métodos exclusivos da struct User
func (user *User) Deactivate() {
	user.Audit.Active = false
	user.Audit.Updated_at = time.Now()
	user.Audit.Deleted_at = time.Now()
}

func (user *User) Activate() {
	user.Audit.Active = true
	user.Audit.Updated_at = time.Now()
	user.Audit.Deleted_at = time.Time{}
}

func main() {
	// Tem como fazer das duas formas.
	beloHorizonte := Address{Country: "Brazil", City: "Belo Horizonte"}
	saoPaulo := Address{Country: "Brazil", City: "São Paulo"}
	extrema := CreateAddress("Brazil", "Extrema")

	matheus := CreateUser(1, true, "Matheus", "matheus@outlook.com", []byte("123"), time.Date(1995, time.April, 04, 0, 0, 0, 0, time.Local), beloHorizonte)
	pamella := CreateUser(2, true, "Pâmella", "pamella@gmail.com", []byte("321"), time.Date(1998, time.December, 12, 0, 0, 0, 0, time.Local), saoPaulo)
	gerson := CreateUser(3, false, "Gerson", "gerson@yahoo.com", []byte("999"), time.Date(1991, time.April, 10, 0, 0, 0, 0, time.Local), beloHorizonte)
	lucas := CreateUser(4, true, "Lucas", "lucas@hotmail.com", []byte("129"), time.Date(1996, time.February, 05, 0, 0, 0, 0, time.Local), extrema)

	Print(matheus)
	Print(pamella)
	Print(gerson)

	Print(lucas)
	fmt.Println("Alterando nome do Lucas")
	lucas.Name = "Zin" // As stucts se comportam como métodos públicos
	Print(lucas)

	fmt.Println("Desativando o Matheus")
	matheus.Deactivate()

	Print(matheus)

	fmt.Println("Trocando a cidade da Pâmella")
	pamella.City = "Belo Horizonte" // Pela struct Address compor a struct User, consigo alterar diretamente a propriedade

	Print(pamella)

	fmt.Println("Trocando o ID do gerson")
	gerson.Audit.ID = 10 // Pela struct Audit ser uma variável dentro da struct User, preciso acessar ela antes de trocar a propriedade ID

	Print(gerson)
}

func CreateUser(id int, active bool, name string, email string, password []byte, birthdate time.Time, address Address) User {
	// No CreateUser ambas as structs devem ser adicionadas no respectivo struct
	return User{
		Audit: Base{
			ID:         id,
			Active:     active,
			Created_at: time.Now(),
			Updated_at: time.Now(),
			Deleted_at: Ternary(active, time.Now(), time.Time{}),
		},
		Address: Address{
			City:    address.City,
			Country: address.Country,
		},
		Name:      name,
		Email:     email,
		Password:  HashPassword(password),
		Birthdate: birthdate,
	}
}

func CreateAddress(country string, city string) Address {
	return Address{
		Country: country,
		City:    city,
	}
}

func Print(user User) {
	fmt.Println(
		"ID:", user.Audit.ID, "\n",
		"Name:", user.Name, "\n",
		"Active:", Ternary(user.Audit.Active, "Yes", "No"), "\n",
		"Created_at:", user.Audit.Created_at, "\n",
		"Deleted_at:", user.Audit.Deleted_at, "\n",
		"Age:", GetAge(user.Birthdate), "\n",
		"Country:", user.Address.Country, "\n",
		"City:", user.Address.City, "\n",
		"Email:", user.Email, "\n",
		"Password:", string(user.Password))
}

func GetAge(date time.Time) int {
	currentYear := time.Now().Year()
	return currentYear - date.Year()
}

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}

func HashPassword(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hash
}
