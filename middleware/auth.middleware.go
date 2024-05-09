package mw

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/manish-neemnarayan/assignment/types"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Printf("authorization header not found")
			http.Error(w, "invalid auth header", http.StatusBadRequest)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[1] != "Bearer" {
			fmt.Printf("invalid authorization header format")
			http.Error(w, "invalid auth header", http.StatusBadRequest)
		}

		claims, err := validateToken(parts[0])
		if err != nil {
			// handler.WriteJson(w, http.)
			http.Error(w, "unauthorized access", http.StatusUnauthorized)
			return
		}

		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		// Check token expiration
		if time.Now().Unix() > expires {
			http.Error(w, "token expired", http.StatusUnauthorized)
		}
		role := claims["role"].(string)

		r.Header.Set("role", role)
		next.ServeHTTP(w, r)
	})
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized error")
		}
		return []byte(types.Secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unauthorized error")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized error")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized error")
	}
	return claims, nil
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "admin" {
			http.Error(w, "unauthorized access", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
