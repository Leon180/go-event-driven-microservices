package loggersconfigs

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logLevel      zapcore.Level `mapstructure:"logLevel"`
	callerEnabled bool          `mapstructure:"callerEnabled"`
	enableTracing bool          `mapstructure:"enableTracing" default:"true"`
}

func (o *Logger) GetLogLevel() zapcore.Level {
	return o.logLevel
}

func (o *Logger) GetCallerEnabled() bool {
	return o.callerEnabled
}

func (o *Logger) GetEnableTracing() bool {
	return o.enableTracing
}

func NewLoggerConfig(env enums.Environment) (*Logger, error) {
	logger, err := configs.BindConfigByKey[Logger](reflect.GetTypeName[Logger](), env)
	if err != nil {
		return nil, err
	}
	return &logger, nil
}
