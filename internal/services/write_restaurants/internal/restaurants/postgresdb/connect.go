package postgresdb

import (
	"log"
	"time"

	gormcustomizelogger "github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers/gorm_customize_logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewGormDB(cfg *GormDBConfig, logger gormcustomizelogger.GormCustomLogger) *gorm.DB {
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: cfg.DBDisableForeignKeyConstraintWhenMigrating,
	}

	if logger == nil {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	} else {
		gormConfig.Logger = logger
		gormConfig.Logger.LogMode(gormlogger.Info)
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN), gormConfig)
	if err != nil {
		log.Fatalf("error occur while connect to postgresql db: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error occur while connect to postgresql db: %s", err)
	}

	sqlDB.SetMaxIdleConns(cfg.DBMaxIdle)
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBMaxLifetimeMinute) * time.Minute)

	return db
}
