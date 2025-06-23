package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"

	"social-network/backend/pkg/db"
	"social-network/backend/pkg/structure"
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
