package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	migratons "github.com/Rizal-Nurochman/Bagi-Pangan-Backend/database/migrations"
	"github.com/Rizal-Nurochman/Bagi-Pangan-Backend/database/seeders"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	db := migratons.SetUpDatabaseConnection()
	defer migratons.CloseDatabaseConnection(db)

	args := os.Args[1:]

	if len(args) > 0 {
		shouldExit := false

		for i, arg := range args {
			switch {
			case arg == "--migrate:run":
				fmt.Println("Running migrations...")
				if err := migratons.RunMigration(db); err != nil {
					log.Fatalf("Migration failed: %v", err)
				}
				fmt.Println("Migrations completed successfully!")
				shouldExit = true

			case arg == "--migrate:rollback":
				if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
					batch, err := strconv.Atoi(args[i+1])
					if err == nil {
						fmt.Printf("Rolling back batch %d...\n", batch)
						if err := migratons.RollbackMigrationBatch(db, batch); err != nil {
							log.Fatalf("Rollback failed: %v", err)
						}
						shouldExit = true
						continue
					}
				}
				fmt.Println("Rolling back last migration batch...")
				if err := migratons.RollbackMigration(db); err != nil {
					log.Fatalf("Rollback failed: %v", err)
				}
				fmt.Println("Rollback completed!")
				shouldExit = true

			case arg == "--migrate:rollback:all":
				fmt.Println("Rolling back all migrations...")
				if err := migratons.RollbackAllMigrations(db); err != nil {
					log.Fatalf("Rollback all failed: %v", err)
				}
				shouldExit = true

			case arg == "--migrate:status":
				if err := migratons.MigrationStatus(db); err != nil {
					log.Fatalf("Failed to get migration status: %v", err)
				}
				shouldExit = true

			case strings.HasPrefix(arg, "--migrate:create:"):
				name := strings.TrimPrefix(arg, "--migrate:create:")
				if err := migratons.CreateMigration(name); err != nil {
					log.Fatalf("Failed to create migration: %v", err)
				}
				shouldExit = true

			case arg == "--seed":
				if err := seeders.RunSeeders(db); err != nil {
					log.Fatalf("Seeding failed: %v", err)
				}
				shouldExit = true
			}
		}

		if shouldExit {
			return
		}
	}

	server := gin.Default()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	fmt.Printf("Server running on port %s\n", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

