package postgresdb

import (
	"log"
	"time"

	gormlogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/grom_logger"
	gormconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/internal/accounts/postgresdb/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormiologger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormDB(cfg *gormconfigs.GormDB, logger gormlogger.GormCustomLogger) *gorm.DB {

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: cfg.GetDBDisableForeignKeyConstraintWhenMigrating(),
	}

	if logger == nil {
		gormConfig.Logger = gormiologger.Default.LogMode(gormiologger.Info)
	} else {
		gormConfig.Logger = logger
		gormConfig.Logger.LogMode(gormiologger.Info)
	}

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), gormConfig)
	if err != nil {
		log.Fatalf("error occur while connect to postgresql db: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error occur while connect to postgresql db: %s", err)
	}

	sqlDB.SetMaxIdleConns(cfg.GetDBMaxIdle())
	sqlDB.SetMaxOpenConns(cfg.GetDBMaxOpen())
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.GetDBMaxLifetimeMinute()) * time.Minute)

	return db
}
