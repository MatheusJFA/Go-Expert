package entity

import (
	"fmt"
	"time"
)

type Role string

const (
	ADMINISTRATOR Role = "Administrator"
	MANAGER       Role = "Manager"
	EMPLOYEE      Role = "Employee"
)

type Worker struct {
	Base
	Company string
	Role    Role
	Salary  float64
}

func CreateWorker(name string, active bool, email string, phone string, birthdate time.Time, company string, role Role, salary float64) Worker {
	return Worker{
		Base: Base{
			Name:      name,
			Active:    active,
			Email:     email,
			Phone:     phone,
			Birthdate: birthdate,
			Audit: Audit{
				Created_At: time.Now(),
				Updated_At: time.Now(),
				Deleted_At: time.Time{},
			},
		},
		Company: company,
		Role:    role,
		Salary:  salary,
	}
}

func (worker *Worker) Promote(role Role) {
	worker.Role = role
}

func (worker *Worker) Print() {
	println("-------------------------------------------------------")
	println("Worker: ", worker.Name)
	println("Company: ", worker.Company)
	println("Role: ", worker.Role)
	println("Salary: ", formatSalary(worker.Salary))
	println("Email: ", worker.Email)
	println("Phone: ", worker.Phone)
	println("Birthdate: ", worker.Birthdate.Format("02/01/2006"))
	println("Active: ", worker.Active)
	println("Created_At: ", worker.Created_At.Format("02/01/2006"))
	println("Updated_At: ", worker.Updated_At.Format("02/01/2006"))
	println("Deleted_At: ", worker.Deleted_At.Format("02/01/2006"))
}

func formatSalary(salary float64) string {
	return fmt.Sprintf("R$ %.2f", salary)
}
