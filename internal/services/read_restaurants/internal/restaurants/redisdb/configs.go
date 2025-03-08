package redisdb

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type RedisConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Password        string `mapstructure:"password"`
	Database        int    `mapstructure:"database"`
	PoolSize        int    `mapstructure:"poolSize"`
	EnableTracing   bool   `mapstructure:"enableTracing"   default:"true"`
	MaxRetries      int    `mapstructure:"maxRetries"      default:"5"`
	MinRetryBackoff int    `mapstructure:"minRetryBackoff" default:"300"` // in milliseconds
	MaxRetryBackoff int    `mapstructure:"maxRetryBackoff" default:"500"` // in milliseconds
	DialTimeout     int    `mapstructure:"dialTimeout"     default:"5"`   // in seconds
	ReadTimeout     int    `mapstructure:"readTimeout"     default:"5"`   // in seconds
	WriteTimeout    int    `mapstructure:"writeTimeout"    default:"3"`   // in seconds
	MinIdleConns    int    `mapstructure:"minIdleConns"    default:"20"`
	PoolTimeout     int    `mapstructure:"poolTimeout"     default:"6"` // in seconds
}

func NewRedisConfig(env enums.Environment) (*RedisConfig, error) {
	typeName := reflect.GetTypeName[RedisConfig]()
	redisConfig, err := configs.BindConfigByKey[RedisConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	return &redisConfig, nil
}
