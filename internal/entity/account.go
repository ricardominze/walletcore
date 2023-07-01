package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	ClientID  string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Client    *Client
}

func NewAccount(client *Client) *Account {

	if client == nil {
		return nil
	}

	account := &Account{
		ID:        uuid.New().String(),
		ClientID:  client.ID,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Client:    client,
	}

	return account
}

func (a *Account) Credit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()

}
