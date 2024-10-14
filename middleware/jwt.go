package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func UserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TO DO: IMPLEMENT

		next.ServeHTTP(w, r)
	})
}

func ParseToken(tok string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// secret := os.Getenv("JWT_SECRET")
		secret := "JWT_SECRET"
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
