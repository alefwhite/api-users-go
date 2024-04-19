package database

import (
	"database/sql"
	"github.com/alefwhite/api-users-go/config/env"
	_ "github.com/lib/pq"
	"log/slog"
)

func NewDBConnection() (*sql.DB, error) {
	postgresURI := env.Env.DatabaseURL

	db, err := sql.Open("postgres", postgresURI)
	if err != nil {

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	slog.Info("database connected", slog.String("package", "database"))

	return db, nil
}
