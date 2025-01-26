package entity

import (
	"time"
)

const (
	CLIENT      = "client"
	CLIENT_PLUS = "client_plus"
	PREMIUM     = "premium"
	VIP         = "vip"
)

type Client struct {
	Base
	cpf  string
	Type string
}

func (client *Client) Deactivate() {
	client.Ativo = false
	client.UpdatedAt = time.Now()
	client.DeletedAt = time.Now()

}

func (client *Client) Activate() {
	client.Ativo = true
	client.UpdatedAt = time.Now()
	client.DeletedAt = time.Time{}
}

func CreateClient(name string, cpf CPF, email string, phone string, birthdate time.Time) Client {

	return Client{
		Base: Base{
			Name:      name,
			Ativo:     true,
			Email:     email,
			Phone:     phone,
			Birthdate: birthdate,
			Audit: Audit{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Time{},
			},
		},
	}
}
