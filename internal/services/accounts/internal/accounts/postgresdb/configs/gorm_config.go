package postgresdbconfigs

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type GormDB struct {
	DSN                                        string            `mapstructure:"dsn"`
	DBDisableForeignKeyConstraintWhenMigrating bool              `mapstructure:"dbDisableForeignKeyConstraintWhenMigrating"`
	DBMaxIdle                                  int               `mapstructure:"dbMaxIdle"`
	DBMaxOpen                                  int               `mapstructure:"dbMaxOpen"`
	DBMaxLifetimeMinute                        int               `mapstructure:"dbMaxLifetimeMinute"`
	Env                                        enums.Environment `mapstructure:"-"`
}

func (o *GormDB) GetDSN() string {
	return o.DSN
}

func (o *GormDB) GetDBDisableForeignKeyConstraintWhenMigrating() bool {
	return o.DBDisableForeignKeyConstraintWhenMigrating
}

func (o *GormDB) GetDBMaxIdle() int {
	return o.DBMaxIdle
}

func (o *GormDB) GetDBMaxOpen() int {
	return o.DBMaxOpen
}

func (o *GormDB) GetDBMaxLifetimeMinute() int {
	return o.DBMaxLifetimeMinute
}

func (o *GormDB) GetEnvironment() enums.Environment {
	return o.Env
}

func NewGormDBConfig(env enums.Environment) (*GormDB, error) {
	typeName := reflect.GetTypeName[GormDB]()
	gormDB, err := configs.BindConfigByKey[GormDB](typeName, env)
	if err != nil {
		return nil, err
	}
	gormDB.Env = env
	return &gormDB, nil
}
