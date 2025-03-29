package postgresdbmigrations

import (
	"log"

	"github.com/Leon180/go-event-driven-microservices/internal/services/read_restaurants/internal/restaurants/entities"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func v202503052348Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&entities.Restaurant{}); err != nil {
			log.Println("AutoMigrate Restaurant failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Branch{}); err != nil {
			log.Println("AutoMigrate Branch failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Address{}); err != nil {
			log.Println("AutoMigrate Address failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.PriceRange{}); err != nil {
			log.Println("AutoMigrate PriceRange failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.BranchCategoryRelation{}); err != nil {
			log.Println("AutoMigrate BranchCategoryRelation failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Category{}); err != nil {
			log.Println("AutoMigrate Category failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Table{}); err != nil {
			log.Println("AutoMigrate Table failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Available{}); err != nil {
			log.Println("AutoMigrate Available failed", "error", err)
			return err
		}
		if err := tx.AutoMigrate(&entities.Book{}); err != nil {
			log.Println("AutoMigrate Book failed", "error", err)
			return err
		}
		return nil
	})
}

var v202503052348 = &gormigrate.Migration{
	ID:      "v202503052348",
	Migrate: v202503052348Migration,
}
