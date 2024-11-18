package types

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(createUser CreateUser) (*User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 10)

	if err != nil {
		return nil, err
	}

	return &User{
		Email:        createUser.Email,
		PasswordHash: string(hashPassword),
	}, nil
}
