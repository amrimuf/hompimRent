package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")

    var err error
    DB, err = pgxpool.Connect(context.Background(), databaseUrl)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    log.Println("Connected to the database!")
}

func GetDB() *pgxpool.Pool {
	return DB
}