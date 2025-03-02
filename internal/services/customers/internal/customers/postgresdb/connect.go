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
		DisableForeignKeyConstraintWhenMigrating: cfg.GetDBDisableForeignKeyConstraintWhenMigrating(),
	}

	if logger == nil {
		gormConfig.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	} else {
		gormConfig.Logger = logger
		gormConfig.Logger.LogMode(gormlogger.Info)
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
