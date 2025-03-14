package rabbitmq

import (
	"fmt"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/configs"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/reflect"
)

type RabbitMQConfig struct {
	DeliveryMode uint8  `mapstructure:"deliveryMode" default:"2"`
	Persisted    bool   `mapstructure:"persisted"    default:"true"`
	AppId        string `mapstructure:"appId"`
	AutoStart    bool   `mapstructure:"autoStart"    default:"true"`

	HostName    string `mapstructure:"hostName"`
	VirtualHost string `mapstructure:"virtualHost"`
	Port        int    `mapstructure:"port"`
	HttpPort    int    `mapstructure:"httpPort"`
	UserName    string `mapstructure:"userName"`
	Password    string `mapstructure:"password"`

	RetryAttempts  int `mapstructure:"retryAttempts"`
	RetryDelay     int `mapstructure:"retryDelay"`     // in milliseconds
	ReconnectDelay int `mapstructure:"reconnectDelay"` // in milliseconds
}

func (c *RabbitMQConfig) AmqpEndPoint() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", c.UserName, c.Password, c.HostName, c.Port)
}

func (c *RabbitMQConfig) HttpEndPoint() string {
	return fmt.Sprintf("http://%s:%d", c.HostName, c.HttpPort)
}

func NewRabbitMQConfig(environment enums.Environment) (*RabbitMQConfig, error) {
	typeName := reflect.GetTypeName[RabbitMQConfig]()
	cfg, err := configs.BindConfigByKey[RabbitMQConfig](typeName, environment)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
