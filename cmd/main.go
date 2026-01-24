package main

import (
	"log"

	"github.com/Rizal-Nurochman/Bagi-Pangan-Backend/database/migrations"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	migrations.MigrationsDatabase()
}

