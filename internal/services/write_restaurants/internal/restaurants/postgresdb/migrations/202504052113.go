package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202504052113Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.OutboxMessage{}); err != nil {
			log.Println("AutoMigrate OutboxMessage failed", "error", err)
			return err
		}
		return nil
	})
}

var v202504052113 = &gormigrate.Migration{
	ID:      "v202504052113",
	Migrate: v202504052113Migration,
}
