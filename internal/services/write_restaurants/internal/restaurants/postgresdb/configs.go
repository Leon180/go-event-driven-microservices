package postgresdb

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type GormDBConfig struct {
	DSN                                        string `mapstructure:"dsn"`
	DBDisableForeignKeyConstraintWhenMigrating bool   `mapstructure:"dbDisableForeignKeyConstraintWhenMigrating"`
	DBMaxIdle                                  int    `mapstructure:"dbMaxIdle"`
	DBMaxOpen                                  int    `mapstructure:"dbMaxOpen"`
	DBMaxLifetimeMinute                        int    `mapstructure:"dbMaxLifetimeMinute"`
}

func NewGormDBConfig(env enums.Environment) (*GormDBConfig, error) {
	typeName := reflect.GetTypeName[GormDBConfig]()
	gormDB, err := configs.BindConfigByKey[GormDBConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	return &gormDB, nil
}
