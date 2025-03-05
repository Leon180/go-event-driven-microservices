package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202502171350Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.Account{}); err != nil {
			log.Println("AutoMigrate Account failed", "error", err)
			return err
		}
		return nil
	})
}

var v202502171350 = &gormigrate.Migration{
	ID:      "v202502171350",
	Migrate: v202502171350Migration,
}
