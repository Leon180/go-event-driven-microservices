package mongodb

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type MongoDbConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	UseAuth         bool   `mapstructure:"useAuth"`
	EnableTracing   bool   `mapstructure:"enableTracing"   default:"true"`
	ConnectTimeout  int    `mapstructure:"connectTimeout"`  // in seconds
	MaxConnIdleTime int    `mapstructure:"maxConnIdleTime"` // in seconds
	MinPoolSize     uint64 `mapstructure:"minPoolSize"`
	MaxPoolSize     uint64 `mapstructure:"maxPoolSize"`
}

func NewMongoDbConfig(env enums.Environment) (*MongoDbConfig, error) {
	typeName := reflect.GetTypeName[MongoDbConfig]()
	mongoDB, err := configs.BindConfigByKey[MongoDbConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	return &mongoDB, nil
}
