package configs

import (
	"fmt"
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	customizegin "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	reflect "github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"github.com/gin-contrib/cors"
	"github.com/samber/lo"
)

func NewAppConfig(env enums.Environment) (customizegin.GinConfig, error) {
	typeName := reflect.GetTypeName[AppConfig]()
	app, err := configs.BindConfigByKey[AppConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	app.Env = env
	return &app, nil
}

type AppConfig struct {
	ConnWebPort string `mapstructure:"connWebPort"`
	ServiceName string `mapstructure:"serviceName"`
	Version     string `mapstructure:"version"`
	// TokenSymmetricKey                          string        `mapstructure:"tokenSymmetricKey"`
	// AccessTokenDuration                        time.Duration `mapstructure:"accessTokenDuration"`
	// RefreshTokenDuration                       time.Duration `mapstructure:"refreshTokenDuration"`
	// RefreshDuration                            time.Duration `mapstructure:"refreshDuration"`
	MaxAge           int               `mapstructure:"maxAge"`
	AllowAllOrigins  bool              `mapstructure:"allowAllOrigins"`
	AllowCredentials bool              `mapstructure:"allowCredentials"`
	AllowMethods     string            `mapstructure:"allowMethods"`
	AllowHeaders     string            `mapstructure:"allowHeaders"`
	ExposeHeaders    string            `mapstructure:"exposeHeaders"`
	Env              enums.Environment `mapstructure:"-"`
}

func (o *AppConfig) GetConnWebPort() string {
	return o.ConnWebPort
}

func (o *AppConfig) GetVersion() string {
	return o.Version
}

func (o *AppConfig) GetServiceName() string {
	return o.ServiceName
}

func (o *AppConfig) GetBasePath() string {
	if o.Version == "" {
		return o.ServiceName
	}
	return fmt.Sprintf("%s/%s", o.Version, o.ServiceName)
}

func (o *AppConfig) GetEnvironment() enums.Environment {
	return o.Env
}

func (o *AppConfig) GenerateCORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = o.AllowAllOrigins
	corsConfig.AllowCredentials = o.AllowCredentials
	allowMethods := strings.Split(o.AllowMethods, ",")
	allowMethods = lo.Map(allowMethods, func(item string, _ int) string {
		return strings.TrimSpace(item)
	})
	allowHeaders := strings.Split(o.AllowHeaders, ",")
	allowHeaders = lo.Map(allowHeaders, func(item string, _ int) string {
		return strings.TrimSpace(item)
	})
	exposeHeaders := strings.Split(o.ExposeHeaders, ",")
	exposeHeaders = lo.Map(exposeHeaders, func(item string, _ int) string {
		return strings.TrimSpace(item)
	})
	corsConfig.AllowMethods = allowMethods
	corsConfig.AllowHeaders = allowHeaders
	corsConfig.ExposeHeaders = exposeHeaders

	return corsConfig
}
