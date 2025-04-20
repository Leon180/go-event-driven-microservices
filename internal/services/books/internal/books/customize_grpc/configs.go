package customizegrpc

import (
	"time"

	"google.golang.org/grpc/keepalive"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	reflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type GRPCConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	ServiceName string `mapstructure:"serviceName"`
	Version     string `mapstructure:"version"`
	BasePath    string `mapstructure:"basePath"`

	MaxConnectionIdle     int `mapstructure:"maxConnectionIdle"`
	MaxConnectionAge      int `mapstructure:"maxConnectionAge"`
	MaxConnectionAgeGrace int `mapstructure:"maxConnectionAgeGrace"`
	Time                  int `mapstructure:"time"`
	Timeout               int `mapstructure:"timeout"`

	Environment enums.Environment `mapstructure:"-"`
}

func (o *GRPCConfig) GetHost() string {
	return o.Host
}

func (o *GRPCConfig) GetPort() string {
	return o.Port
}

func (o *GRPCConfig) GetServiceName() string {
	return o.ServiceName
}

func (o *GRPCConfig) GetVersion() string {
	return o.Version
}

func (o *GRPCConfig) GetBasePath() string {
	return o.BasePath
}

func (o *GRPCConfig) GetEnvironment() enums.Environment {
	return o.Environment
}

func (o *GRPCConfig) GetKeepAliveParams() keepalive.ServerParameters {
	return keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(o.MaxConnectionIdle) * time.Second,
		MaxConnectionAge:      time.Duration(o.MaxConnectionAge) * time.Second,
		MaxConnectionAgeGrace: time.Duration(o.MaxConnectionAgeGrace) * time.Second,
		Time:                  time.Duration(o.Time) * time.Second,
		Timeout:               time.Duration(o.Timeout) * time.Second,
	}
}

func NewGRPCConfig(env enums.Environment) (customizegrpc.GRPCConfig, error) {
	typeName := reflect.GetTypeName[GRPCConfig]()
	app, err := configs.BindConfigByKey[GRPCConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	app.Environment = env
	return &app, nil
}
