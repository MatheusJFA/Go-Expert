package main

import (
	"GoExpert/entity"
	"fmt"
	"time"
)

func main() {
	client, _ := entity.CreateClient("John Doe", true, "123.456.789-09", "johndoe@gmail.com", "(11) 99999-9999", time.Date(1990, time.January, 01, 0, 0, 0, 0, time.Local), entity.BASIC)
	worker := entity.CreateWorker("Jane Doe", true, "janedoe@gmail.com", "(11) 88888-8888", time.Date(1990, time.August, 16, 0, 0, 0, 0, time.Local), "Company", entity.ADMINISTRATOR, 10000)
	employee := entity.CreateWorker("Jack Doe", true, "jackdoe@gmail.com", "(11) 77777-7777", time.Date(1990, time.August, 16, 0, 0, 0, 0, time.Local), "Company", entity.EMPLOYEE, 3000)

	client.Print()

	worker.Print()

	employee.Print()

	// This will panic (invalid CPF)
	fmt.Println("-------------------------------------------------------")
	_, error := entity.CreateClient("Emilly Doe", true, "123.456.789-00", "emillyDoe@gmail.com", "(22) 88888-8888", time.Date(1990, time.January, 01, 0, 0, 0, 0, time.Local), entity.PLUS)
	fmt.Println(error)
}
