package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Erreur ouverture DB: %v", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Erreur driver migration: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migration/sqlite",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Erreur initialisation migration: %v", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Erreur application migration: %v", err)
	}

	return db
}
