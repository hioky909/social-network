package handler

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var JwtSecret = []byte("super_secret_key") // Utilise la même clé que pour la génération

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token manquant", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Token invalide", http.StatusUnauthorized)
			return
		}
		// Tu peux ajouter ici la récupération des claims si besoin
		next.ServeHTTP(w, r)
	})
}
