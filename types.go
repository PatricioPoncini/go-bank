package main

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand/v2"
	"time"
)

type LoginRequest struct {
	Number   int    `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TransferRequest struct {
	ToAccount int64 `json:"toAccount"`
	Amount    int   `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (a *Account) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)) == nil
}

func NewAccount(firstName, lastName, password string) (*Account, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Number:            int64(rand.IntN(10000000)),
		EncryptedPassword: string(encryptedPassword),
		Balance:           0,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
