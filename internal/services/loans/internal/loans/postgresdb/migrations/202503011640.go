package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/loans/internal/loans/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202503011640Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.Loan{}); err != nil {
			log.Println("AutoMigrate Loan failed", "error", err)
			return err
		}
		return nil
	})
}

var v202503011640 = &gormigrate.Migration{
	ID:      "v202503011640",
	Migrate: v202503011640Migration,
}
