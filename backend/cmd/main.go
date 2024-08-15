package main

import (
	"log"
	"os"

	"github.com/amrimuf/hompimRent/bootstrap"
)

func main() {
    app := bootstrap.NewApp()
    app.Bootstrap()

	port := getPort()
	log.Printf("Server starting on port %s\n", port)
    if err := app.Start(":" + port); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}