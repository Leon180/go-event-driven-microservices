package postgresdb

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type GormDBConfig struct {
	DSN                                        string            `mapstructure:"dsn"`
	DBDisableForeignKeyConstraintWhenMigrating bool              `mapstructure:"dbDisableForeignKeyConstraintWhenMigrating"`
	DBMaxIdle                                  int               `mapstructure:"dbMaxIdle"`
	DBMaxOpen                                  int               `mapstructure:"dbMaxOpen"`
	DBMaxLifetimeMinute                        int               `mapstructure:"dbMaxLifetimeMinute"`
	Env                                        enums.Environment `mapstructure:"-"`
}

func (o *GormDBConfig) GetDSN() string {
	return o.DSN
}

func (o *GormDBConfig) GetDBDisableForeignKeyConstraintWhenMigrating() bool {
	return o.DBDisableForeignKeyConstraintWhenMigrating
}

func (o *GormDBConfig) GetDBMaxIdle() int {
	return o.DBMaxIdle
}

func (o *GormDBConfig) GetDBMaxOpen() int {
	return o.DBMaxOpen
}

func (o *GormDBConfig) GetDBMaxLifetimeMinute() int {
	return o.DBMaxLifetimeMinute
}

func (o *GormDBConfig) GetEnvironment() enums.Environment {
	return o.Env
}

func NewGormDBConfig(env enums.Environment) (*GormDBConfig, error) {
	typeName := reflect.GetTypeName[GormDBConfig]()
	gormDB, err := configs.BindConfigByKey[GormDBConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	gormDB.Env = env
	return &gormDB, nil
}
