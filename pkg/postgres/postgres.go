package postgres

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewPostgres(cfg DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("NewPostgres: %w", err)
	}
	return database, nil
}
