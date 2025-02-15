package postgresdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	accountconfigs "github.com/Leon180/go-event-driven-microservices/internal/services/accounts/configs"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitGormDB(config accountconfigs.Config, zapLogger *zap.Logger) *gorm.DB {

	var (
		connectionString string
		err              error
		gormConfig       = &gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
			DisableForeignKeyConstraintWhenMigrating: config.DBDisableForeignKeyConstraintWhenMigrating,
		}
		db    *gorm.DB
		sqlDB *sql.DB
	)

	if config.Environment == "development" {
		connectionString = config.DSNTest
	} else {
		connectionString = config.DSNDeployment
	}

	if zapLogger == nil {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.Logger = utilities.NewGormZapLogger(zapLogger)
		gormConfig.Logger.LogMode(logger.Info)
	}

	db, err = gorm.Open(postgres.Open(connectionString), gormConfig)
	if err != nil {
		utilities.LogFatal(context.Background(), "error occur while connect to postgresql db: %s", err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		utilities.LogFatal(context.Background(), "error occur while connect to postgresql db: %s", err)
	}

	sqlDB.SetMaxIdleConns(config.DBMaxIdle)
	sqlDB.SetMaxOpenConns(config.DBMaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(config.DBMaxLifetimeMinute) * time.Minute)

	return db
}
