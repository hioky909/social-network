package db

import (
	"database/sql"
	"errors"
	"strings"

	"social-network/backend/pkg/structure"

	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, form structure.RegisterForm) error {
	// Vérifier unicité email/username
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? OR username = ?", form.Email, form.Username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("Email ou nom d'utilisateur déjà utilisé")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        INSERT INTO users (email, username, password, first_name, last_name, date_of_birth, avatar, nickname, about_me)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		strings.TrimSpace(form.Email),
		strings.TrimSpace(form.Username),
		string(hashed),
		strings.TrimSpace(form.FirstName),
		strings.TrimSpace(form.LastName),
		form.DateOfBirth,
		strings.TrimSpace(form.Avatar),
		strings.TrimSpace(form.Nickname),
		strings.TrimSpace(form.AboutMe),
	)
	return err
}

func AuthenticateUser(db *sql.DB, form structure.LoginForm) (int, error) {
    var id int
    var hashed string
    err := db.QueryRow(
        `SELECT id, password FROM users WHERE email = ? OR username = ?`,
        form.EmailOrUsername, form.EmailOrUsername,
    ).Scan(&id, &hashed)
    if err != nil {
        return 0, errors.New("Utilisateur non trouvé")
    }
    if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(form.Password)) != nil {
        return 0, errors.New("Mot de passe incorrect")
    }
    return id, nil
}