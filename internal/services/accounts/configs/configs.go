package configs

import (
	"fmt"
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	customizeginconfigs "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_gin/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"github.com/gin-contrib/cors"
	"github.com/samber/lo"
)

type App struct {
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

func (o *App) GetConnWebPort() string {
	return o.ConnWebPort
}

func (o *App) GetVersion() string {
	return o.Version
}

func (o *App) GetServiceName() string {
	return o.ServiceName
}

func (o *App) GetBasePath() string {
	if o.Version == "" {
		return o.ServiceName
	}
	return fmt.Sprintf("%s/%s", o.Version, o.ServiceName)
}

func (o *App) GetEnvironment() enums.Environment {
	return o.Env
}

func (o *App) GenerateCORSConfig() cors.Config {
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

func NewAppConfig(env enums.Environment) (customizeginconfigs.GinConfig, error) {
	typeName := reflect.GetTypeName[App]()
	app, err := configs.BindConfigByKey[App](typeName, env)
	if err != nil {
		return nil, err
	}
	app.Env = env
	return &app, nil
}
