package main

import (
	"fmt"
	"strings"
	"time"
)

type Person[T any] struct {
	ID        int
	Name      string
	Birthdate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Employee struct {
	Details  Person[Employee]
	Position string
	Salary   float64
}

type Manager struct {
	Details      Person[Manager]
	Department   string
	Subordinates []Employee
}

type Worker interface {
	Employee | Manager
}

func main() {
	employee := CreateEmployee(1, "John Doe", time.Date(1995, time.April, 12, 0, 0, 0, 0, time.UTC), "Developer", 7_500.94)
	employee2 := CreateEmployee(2, "Jane Doe", time.Date(1998, time.December, 27, 0, 0, 0, 0, time.UTC), "Developer", 7_500.94)
	employee3 := CreateEmployee(3, "Alice Doe", time.Date(1997, time.January, 15, 0, 0, 0, 0, time.UTC), "Developer", 7_500.94)
	manager := CreateManager(4, "Bob Doe", time.Date(1990, time.March, 5, 0, 0, 0, 0, time.UTC), "Development", []Employee{employee, employee2, employee3})

	pessoa := CreatePerson(5, "Matheus", time.Date(1995, time.April, 4, 0, 0, 0, 0, time.UTC))

	PrintInformation(employee)
	PrintInformation(employee2)
	PrintInformation(employee3)

	PrintInformation(manager)

	// A função PrintInformation é genérica e seu comportamento muda de acordo com o tipo de parâmetro passado, no caso, Pessoa não é aceito, logo, irá retornar "Tipo não reconhecido"
	PrintInformation(pessoa)
}

// Função que cria uma pessoa
func CreatePerson(id int, name string, birthdate time.Time) Person[any] {
	return Person[any]{
		ID:        id,
		Name:      name,
		Birthdate: birthdate,
	}
}

// Função que cria um empregado
func CreateEmployee(id int, name string, birthdate time.Time, position string, salary float64) Employee {
	return Employee{
		Details: Person[Employee]{
			ID:        id,
			Name:      name,
			Birthdate: birthdate,
		},
		Position: position,
		Salary:   salary,
	}
}

// Função que cria um gerente
func CreateManager(id int, name string, birthdate time.Time, department string, subordinates []Employee) Manager {
	return Manager{
		Details: Person[Manager]{
			ID:        id,
			Name:      name,
			Birthdate: birthdate,
		},
		Department:   department,
		Subordinates: subordinates,
	}
}

// Função genérica que imprime informações de qualquer tipo Pessoa, Empregado ou Gerente
func PrintInformation[T any](tipo T) {
	// Verificando o tipo de T e imprimindo informações específicas
	switch worker := any(tipo).(type) {
	case Employee:
		fmt.Printf("Empregado: ID=%d, \n Nome=%s,\n Nascimento=%s,\n Idade: %d,\n Cargo=%s,\n Salario=R$%.2f \n", worker.Details.ID, worker.Details.Name, worker.Details.Birthdate.Format("02/01/2006"), GetAge(worker.Details.Birthdate), worker.Position, worker.Salary)
	case Manager:
		fmt.Printf("Gerente: ID=%d, \n Nome=%s,\n Nascimento=%s,\n Idade: %d,\n Departamento=%s,\n Subordinados=%s\n", worker.Details.ID, worker.Details.Name, worker.Details.Birthdate.Format("02/01/2006"), GetAge(worker.Details.Birthdate), worker.Department, GetEmployeesNames(worker.Subordinates))
	default:
		fmt.Println("Tipo não reconhecido")
	}
}

// Função que retorna os nomes dos empregados separados por vírgula
func GetEmployeesNames(employees []Employee) string {
	var names []string
	for _, employee := range employees {
		names = append(names, employee.Details.Name)
	}
	return strings.Join(names, ", ")
}

// Função que calcula a idade
func GetAge(birthdate time.Time) int {
	var currentYear = time.Now()
	const hoursInYear = 8760
	return int(currentYear.Sub(birthdate).Hours() / hoursInYear)
}
