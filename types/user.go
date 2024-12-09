package types

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 12)
	if err != nil {
		return nil, err
	}
	return &User{
		Email:        createUser.Email,
		PasswordHash: string(hashPassword),
	}, nil
}

func ValidatePassword(passwordHash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}

func CreateToken(user *User) (string, error) {
	now := time.Now()
	validUntil := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Printf("Error signing token: %v", err)
	}
	return tokenStr, nil
}
