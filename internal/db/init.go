package db

import (
	"context"
)

func (db *DB) Init(ctx context.Context) error {
	log := db.logger.With("method", "Init")

	if err := db.InitUser(ctx); err != nil {
		return err
	}

	if err := db.InitExpense(ctx); err != nil {
		return err
	}

	log.InfoContext(ctx, "success create tables for expense tracker")
	return nil
}

func (db *DB) InitUser(ctx context.Context) error {
	log := db.logger.With("table", "category")
	stmt := `
	CREATE TABLE IF NOT EXISTS users_
	(
		id serial PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
    	password_hash TEXT NOT NULL,
    	salt TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.Pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create user table", "error", err)
		return err
	}

	return nil
}

func (db *DB) InitExpense(ctx context.Context) error {
	log := db.logger.With("table", "category")
	stmt := `
		CREATE TABLE IF NOT EXISTS expense
	(
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users_ (id) ON DELETE CASCADE NOT NULL,
		amount INT NOT NULL,
		date_expense DATE NOT NULL,
		category TEXT CHECK (category IN ('Groceries', 'Leisure', 'Electronics', 'Utilities', 'Clothing', 'Health', 'Transport', 'Education', 'Credits', 'Others')) NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);`

	if _, err := db.Pg.Exec(stmt); err != nil {
		log.ErrorContext(ctx, "failed to create expense table", "error", err)
		return err
	}

	return nil
}
