package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202502271845Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.CreditCard{}); err != nil {
			log.Println("AutoMigrate CreditCard failed", "error", err)
			return err
		}
		return nil
	})
}

var v202502271845 = &gormigrate.Migration{
	ID:      "v202502271845",
	Migrate: v202502271845Migration,
}
