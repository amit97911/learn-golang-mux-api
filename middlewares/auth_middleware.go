package middlewares

import (
	"learn-golang-mux-api/config"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string) (string, error) {
	secret := config.GetEnv("JWT_SECRET", "my_secret_key")
	jwtTtl := config.GetEnv("JWT_TTL", "120")
	ttlMinutes, err := strconv.Atoi(jwtTtl)
	if err != nil {
		log.Println("Failed to convert jwt ttl to integer")
		panic(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Duration(ttlMinutes) * time.Minute).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, "Bearer ")
		if len(authHeaderParts) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenStr := authHeaderParts[1]
		secret := config.GetEnv("JWT_SECRET", "my_secret_key")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
