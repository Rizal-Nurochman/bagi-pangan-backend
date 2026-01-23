package migratons

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Rizal-Nurochman/Bagi-Pangan-Backend/database/entities"
	"gorm.io/gorm"
)

type MigrationRecord struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Migration string    `gorm:"type:varchar(255);not null"`
	Batch     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone"`
}

func (MigrationRecord) TableName() string {
	return "migrations"
}

var migrationEntities = []interface{}{
	&entities.User{},
	&entities.Category{},
	&entities.MitraProfile{},
	&entities.ReceiptProfile{},
	&entities.SurplusListing{},
	&entities.Transaction{},
	&entities.TransactionItem{},
}

var migrationNames = []string{
	"create_users_table",
	"create_categories_table",
	"create_mitra_profiles_table",
	"create_receipt_profiles_table",
	"create_surplus_listings_table",
	"create_transactions_table",
	"create_transaction_items_table",
}

var tableNames = []string{
	"users",
	"categories",
	"mitra_profiles",
	"receipt_profiles",
	"surplus_listings",
	"transactions",
	"transaction_items",
}

func RunMigration(db *gorm.DB) error {
	db.AutoMigrate(&MigrationRecord{})

	var lastBatch int
	db.Model(&MigrationRecord{}).Select("COALESCE(MAX(batch), 0)").Scan(&lastBatch)
	newBatch := lastBatch + 1

	for i, entity := range migrationEntities {
		var exists MigrationRecord
		if err := db.Where("migration = ?", migrationNames[i]).First(&exists).Error; err == gorm.ErrRecordNotFound {
			if err := db.AutoMigrate(entity); err != nil {
				return fmt.Errorf("failed to migrate %s: %v", migrationNames[i], err)
			}
			db.Create(&MigrationRecord{
				Migration: migrationNames[i],
				Batch:     newBatch,
				CreatedAt: time.Now(),
			})
			fmt.Printf("Migrated: %s\n", migrationNames[i])
		} else {
			fmt.Printf("Already migrated: %s\n", migrationNames[i])
		}
	}

	return nil
}

func RollbackMigration(db *gorm.DB) error {
	db.AutoMigrate(&MigrationRecord{})

	var lastBatch int
	db.Model(&MigrationRecord{}).Select("COALESCE(MAX(batch), 0)").Scan(&lastBatch)

	if lastBatch == 0 {
		fmt.Println("Nothing to rollback.")
		return nil
	}

	return rollbackBatch(db, lastBatch)
}

func RollbackMigrationBatch(db *gorm.DB, batch int) error {
	db.AutoMigrate(&MigrationRecord{})
	return rollbackBatch(db, batch)
}

func RollbackAllMigrations(db *gorm.DB) error {
	db.AutoMigrate(&MigrationRecord{})

	for i := len(tableNames) - 1; i >= 0; i-- {
		if db.Migrator().HasTable(tableNames[i]) {
			if err := db.Migrator().DropTable(tableNames[i]); err != nil {
				return fmt.Errorf("failed to drop table %s: %v", tableNames[i], err)
			}
			fmt.Printf("Dropped: %s\n", tableNames[i])
		}
	}

	db.Where("1 = 1").Delete(&MigrationRecord{})
	fmt.Println("All migrations rolled back.")
	return nil
}

func rollbackBatch(db *gorm.DB, batch int) error {
	var records []MigrationRecord
	db.Where("batch = ?", batch).Order("id DESC").Find(&records)

	if len(records) == 0 {
		fmt.Printf("No migrations found for batch %d.\n", batch)
		return nil
	}

	for _, record := range records {
		idx := -1
		for i, name := range migrationNames {
			if name == record.Migration {
				idx = i
				break
			}
		}

		if idx >= 0 && db.Migrator().HasTable(tableNames[idx]) {
			if err := db.Migrator().DropTable(tableNames[idx]); err != nil {
				return fmt.Errorf("failed to drop table %s: %v", tableNames[idx], err)
			}
			fmt.Printf("Rolled back: %s\n", record.Migration)
		}

		db.Delete(&record)
	}

	return nil
}

func MigrationStatus(db *gorm.DB) error {
	db.AutoMigrate(&MigrationRecord{})

	var records []MigrationRecord
	db.Order("batch ASC, id ASC").Find(&records)

	if len(records) == 0 {
		fmt.Println("No migrations have been run.")
		return nil
	}

	fmt.Println("+-------+----------------------------------------+-------+")
	fmt.Println("| Batch | Migration                              | Ran   |")
	fmt.Println("+-------+----------------------------------------+-------+")

	for _, record := range records {
		fmt.Printf("| %-5d | %-38s | Yes   |\n", record.Batch, record.Migration)
	}

	fmt.Println("+-------+----------------------------------------+-------+")
	return nil
}

func CreateMigration(name string) error {
	migrationsDir := "database/migrations/files"
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return fmt.Errorf("failed to create migrations directory: %v", err)
	}

	timestamp := time.Now().Format("20060102150405")
	filename := fmt.Sprintf("%s_%s.go", timestamp, name)
	filepath := filepath.Join(migrationsDir, filename)

	content := fmt.Sprintf(`package files

// Migration: %s
// Created at: %s

// TODO: Implement your migration logic here
// Add the entity to migrationEntities and migrationNames in migration.go
`, name, time.Now().Format("2006-01-02 15:04:05"))

	if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create migration file: %v", err)
	}

	fmt.Printf("Created migration: %s\n", filepath)
	return nil
}
