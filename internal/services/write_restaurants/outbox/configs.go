package outbox

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type OutboxConfig struct {
	Interval   time.Duration `mapstructure:"interval"`
	MaxRetries int           `mapstructure:"maxRetries"`
}

func NewOutboxConfig(env enums.Environment) (*OutboxConfig, error) {
	typeName := reflect.GetTypeName[OutboxConfig]()
	outbox, err := configs.BindConfigByKey[OutboxConfig](typeName, env)
	if err != nil {
		return nil, err
	}
	return &outbox, nil
}
