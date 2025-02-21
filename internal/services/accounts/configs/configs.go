package configs

import (
	"strings"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"github.com/gin-contrib/cors"
)

type App struct {
	connWebPort string `mapstructure:"connWebPort"`
	serviceName string `mapstructure:"serviceName"`
	// TokenSymmetricKey                          string        `mapstructure:"tokenSymmetricKey"`
	// AccessTokenDuration                        time.Duration `mapstructure:"accessTokenDuration"`
	// RefreshTokenDuration                       time.Duration `mapstructure:"refreshTokenDuration"`
	// RefreshDuration                            time.Duration `mapstructure:"refreshDuration"`
	maxAge           int               `mapstructure:"maxAge"`
	allowAllOrigins  bool              `mapstructure:"allowAllOrigins"`
	allowCredentials bool              `mapstructure:"allowCredentials"`
	allowMethods     string            `mapstructure:"allowMethods"`
	allowHeaders     string            `mapstructure:"allowHeaders"`
	exposeHeaders    string            `mapstructure:"exposeHeaders"`
	env              enums.Environment `mapstructure:"-"`
}

func (o *App) GetConnWebPort() string {
	return o.connWebPort
}

func (o *App) GetServiceName() string {
	return o.serviceName
}

func (o *App) GetMaxAge() int {
	return o.maxAge
}

func (o *App) GetAllowAllOrigins() bool {
	return o.allowAllOrigins
}

func (o *App) GetAllowCredentials() bool {
	return o.allowCredentials
}

func (o *App) GetAllowMethods() string {
	return o.allowMethods
}

func (o *App) GetAllowHeaders() string {
	return o.allowHeaders
}

func (o *App) GetExposeHeaders() string {
	return o.exposeHeaders
}

func (o *App) GetEnvironment() enums.Environment {
	return o.env
}

func (o *App) GenerateCORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = o.allowAllOrigins
	corsConfig.AllowCredentials = o.allowCredentials
	corsConfig.AllowMethods = strings.Split(o.allowMethods, ",")
	corsConfig.AllowHeaders = strings.Split(o.allowHeaders, ",")
	corsConfig.ExposeHeaders = strings.Split(o.exposeHeaders, ",")
	return corsConfig
}

func NewAppConfig(env enums.Environment) (*App, error) {
	app, err := configs.BindConfigByKey[App](reflect.GetTypeName[App](), env)
	if err != nil {
		return nil, err
	}
	app.env = env
	return &app, nil
}
