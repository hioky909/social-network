package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"social-network/backend/pkg/db"
	"social-network/backend/pkg/structure"

	"github.com/golang-jwt/jwt"
)

func isEmailValid(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func RegisterHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		var form structure.RegisterForm
		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}

		// Validation basique
		if !isEmailValid(form.Email) ||
			len(form.Username) < 3 ||
			len(form.Password) < 6 ||
			len(form.FirstName) < 1 ||
			len(form.LastName) < 1 ||
			len(form.DateOfBirth) < 4 {
			http.Error(w, "Champs invalides", http.StatusBadRequest)
			return
		}

		if err := db.InsertUser(dbConn, form); err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Inscription réussie"})
	}
}

var jwtSecret = []byte("super_secret_key") // À stocker dans une variable d'env en prod

func LoginHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}
		var form structure.LoginForm
		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}
		userID, err := db.AuthenticateUser(dbConn, form)
		if err != nil {
			http.Error(w, "Identifiants invalides", http.StatusUnauthorized)
			return
		}
		// Générer le token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
