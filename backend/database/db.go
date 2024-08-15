package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("DATABASE_URL environment variable not set")
	}

	db, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database!")
	return db, nil
}

func GetDB() *pgxpool.Pool {
	return DB
}