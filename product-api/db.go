package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connectDB() (*pgxpool.Pool, error) {
	databaseURL := "postgres://postgres:postgres123@localhost:5432/product_db"

	db, err := pgxpool.New(
		context.Background(),
		databaseURL,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}
