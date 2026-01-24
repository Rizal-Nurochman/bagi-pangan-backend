package migrations

import (
	"log"

	"github.com/Rizal-Nurochman/Bagi-Pangan-Backend/database/entities"
)

func MigrationsDatabase() {
	db := SetUpDatabaseConnection()

	err := db.AutoMigrate(
		&entities.User{},
		&entities.MitraProfile{},
		&entities.ReceiptProfile{},
		&entities.Category{},
		&entities.SurplusListing{},
		&entities.TransactionItem{},
		&entities.Transaction{},
		&entities.Review_Mitra{},
	)
	if err != nil {
		log.Println("Database migration failed:", err)
		return
	}
	log.Println("Database migration completed.")
}