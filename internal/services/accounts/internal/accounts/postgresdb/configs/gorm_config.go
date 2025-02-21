package postgresdbconfigs

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type GormDB struct {
	dsn                                        string            `mapstructure:"dsn"`
	dbDisableForeignKeyConstraintWhenMigrating bool              `mapstructure:"dbDisableForeignKeyConstraintWhenMigrating"`
	dbMaxIdle                                  int               `mapstructure:"dbMaxIdle"`
	dbMaxOpen                                  int               `mapstructure:"dbMaxOpen"`
	dbMaxLifetimeMinute                        int               `mapstructure:"dbMaxLifetimeMinute"`
	env                                        enums.Environment `mapstructure:"-"`
}

func (o *GormDB) GetDSN() string {
	return o.dsn
}

func (o *GormDB) GetDBDisableForeignKeyConstraintWhenMigrating() bool {
	return o.dbDisableForeignKeyConstraintWhenMigrating
}

func (o *GormDB) GetDBMaxIdle() int {
	return o.dbMaxIdle
}

func (o *GormDB) GetDBMaxOpen() int {
	return o.dbMaxOpen
}

func (o *GormDB) GetDBMaxLifetimeMinute() int {
	return o.dbMaxLifetimeMinute
}

func (o *GormDB) GetEnvironment() enums.Environment {
	return o.env
}

func NewGormDBConfig(env enums.Environment) (*GormDB, error) {
	gormDB, err := configs.BindConfigByKey[GormDB](reflect.GetTypeName[GormDB](), env)
	if err != nil {
		return nil, err
	}
	return &gormDB, nil
}
