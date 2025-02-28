package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202503010242Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.Customer{}); err != nil {
			log.Println("AutoMigrate Customer failed", "error", err)
			return err
		}
		return nil
	})
}

var v202503010242 = &gormigrate.Migration{
	ID:      "v202503010242",
	Migrate: v202503010242Migration,
}
