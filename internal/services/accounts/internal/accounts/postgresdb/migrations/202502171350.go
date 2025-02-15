package migrations

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/dbmodels"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202502171350Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&dbmodels.Account{}); err != nil {
			utilities.LogError(context.Background(), "AutoMigrate Account failed", "error", err)
			return err
		}
		return nil
	})
}

var v202502171350 = &gormigrate.Migration{
	ID:      "v202502171350",
	Migrate: v202502171350Migration,
}
