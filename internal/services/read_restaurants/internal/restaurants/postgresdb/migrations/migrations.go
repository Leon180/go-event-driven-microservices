package postgresdbmigrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrationsV2 = []*gormigrate.Migration{
	v202503052348, // add restaurant, branch, address, price range, branch category relation, category, table, table available table, book
	v202503052351, // add category
}

var migrateOptionsV2 = &gormigrate.Options{
	TableName:                 "go_migrations",
	IDColumnName:              "id",
	IDColumnSize:              255,
	UseTransaction:            false,
	ValidateUnknownMigrations: false,
}

// MigrateDB performs database migrations for the dbmodels
func MigrateDB(db *gorm.DB) error {
	log.Println("Starting database migration...")

	if err := gormigrate.New(db, migrateOptionsV2, migrationsV2).Migrate(); err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
