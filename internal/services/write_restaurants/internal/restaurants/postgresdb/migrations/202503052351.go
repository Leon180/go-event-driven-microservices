package postgresdbmigrations

import (
	"log"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func v202503052351Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		uuidGenerator := uuid.NewUUIDGenerator()
		timeNow := time.Now()
		categories := lo.Map(enums.CategoryCodes, func(code enums.CategoryCode, _ int) entities.Category {
			return entities.Category{
				ID:           uuidGenerator.GenerateUUID(),
				CategoryCode: code,
				CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
			}
		})

		if err := tx.Create(&categories).Error; err != nil {
			log.Println("Create categories failed", "error", err)
			return err
		}
		return nil
	})
}

var v202503052351 = &gormigrate.Migration{
	ID:      "v202503052351",
	Migrate: v202503052351Migration,
}
