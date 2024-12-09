package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goschool/crud/types"
)

type key string

const userKey key = "user"

func WithUser(ctx context.Context, user *types.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *types.User {
	user, ok := ctx.Value(userKey).(*types.User)
	if !ok {
		return nil
	}
	return user
}

func UserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		token := headerParts[1]

		claims, err := ParseToken(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		expires := int64(claims["exp"].(float64))
		today := time.Now().Unix()
		if today > expires {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		var user types.User
		user.ID = claims["id"].(string)
		user.Email = claims["email"].(string)
		ctx := r.Context()
		ctx = WithUser(ctx, &user)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ParseToken(tok string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		// secret := "JWT_SECRET"
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
