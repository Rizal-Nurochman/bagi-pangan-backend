package seeders

import (
	"fmt"

	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) error {
	fmt.Println("Running seeders...")

	// Add your seeders here
	// Example:
	// if err := seedCategories(db); err != nil {
	// 	return err
	// }

	fmt.Println("All seeders completed!")
	return nil
}
