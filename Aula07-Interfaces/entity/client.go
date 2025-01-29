package entity

import (
	"GoExpert/entity/valueObjects"
	"errors"
	"fmt"
	"time"
)

type AccountType string

const (
	BASIC   AccountType = "Basic"
	PLUS    AccountType = "Plus"
	PREMIUM AccountType = "Premium"
	VIP     AccountType = "VIP"
)

type Client struct {
	Base
	CPF         valueObjects.CPF
	AccountType AccountType
}

func CreateClient(name string, active bool, cpf valueObjects.CPF, email string, phone string, birthdate time.Time, accountType AccountType) (Client, error) {
	if cpf.ValidateCPF() == false {
		errorMessage := fmt.Sprintf("The provided CPF %s is invalid", cpf)
		return Client{}, errors.New(errorMessage)
	}

	return Client{
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
		CPF:         cpf,
		AccountType: accountType,
	}, nil
}

func (client *Client) Print() {
	println("-------------------------------------------------------")
	println("Client: ", client.Name)
	println("CPF: ", client.CPF)
	println("AccountType: ", client.AccountType)
	println("Email: ", client.Email)
	println("Phone: ", client.Phone)
	println("Birthdate: ", client.Birthdate.Format("02/01/2006"))
	println("Active: ", client.Active)
	println("Created_At: ", client.Created_At.Format("02/01/2006"))
	println("Updated_At: ", client.Updated_At.Format("02/01/2006"))
	println("Deleted_At: ", client.Deleted_At.Format("02/01/2006"))
}
