package migrations

import (
	"context"
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrationsV2 = []*gormigrate.Migration{
	v202502171350, // add user, session, place and favorite table
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
		utilities.LogError(context.Background(), "Error during migration", "error", err)
		return err
	}

	utilities.LogInfo(context.Background(), "Database migration completed successfully")
	return nil
}
