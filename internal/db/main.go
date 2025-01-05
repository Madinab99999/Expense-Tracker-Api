package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/configs"

	_ "github.com/lib/pq"
)

type DB struct {
	logger *slog.Logger
	Pg     *sql.DB
}

func New(logger *slog.Logger, cfg *configs.Config) (*DB, error) {
	pgsql, err := NewPgSQL(cfg)
	if err != nil {
		return nil, err
	}

	return &DB{
		logger: logger,
		Pg:     pgsql,
	}, nil
}

func NewPgSQL(cfg *configs.Config) (*sql.DB, error) {
	host := cfg.DBHost
	port, err := strconv.Atoi(cfg.DBPort)
	if err != nil {
		return nil, err
	}
	user := cfg.DBUser
	password := cfg.DBPassword
	dbname := cfg.DBName

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Ping(ctx context.Context) error {
	err := db.Pg.PingContext(ctx)
	if err != nil {
		db.logger.ErrorContext(ctx, "failed to connect to database", "error", err)
		return err
	}

	db.logger.InfoContext(ctx, "success connected to database")
	return nil
}
