package connect

import (
	"fmt"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPConnection interface {
	IsClosed() bool
	ReconnectedEvent() chan struct{} // channel store if the connection is reconnected
	ErrorEvent() chan error          // channel store if the connection is closed
	NewChannel() (*amqp.Channel, error)
	Close() error
	NotifyClose(receiver chan *amqp.Error) chan *amqp.Error
	Raw() *amqp.Connection
}

func NewRabbitMQConnection(config *rabbitmq.RabbitMQConfig, logger loggers.Logger) (AMQPConnection, error) {
	if config == nil {
		return nil, fmt.Errorf("rabbitmq host options is nil")
	}

	c := &amqpConnectionImpl{
		logger:           logger,
		config:           config,
		reconnectedEvent: make(chan struct{}),
		errorEvent:       make(chan error),
	}

	err := c.connect()
	if err != nil {
		return nil, err
	}

	return c, nil
}

type amqpConnectionImpl struct {
	*amqp.Connection
	logger           loggers.Logger
	config           *rabbitmq.RabbitMQConfig
	reconnectedEvent chan struct{}
	errorEvent       chan error
	connClose        chan *amqp.Error
}

func (c *amqpConnectionImpl) ReconnectedEvent() chan struct{} {
	return c.reconnectedEvent
}

func (c *amqpConnectionImpl) ErrorEvent() chan error {
	return c.errorEvent
}

func (c *amqpConnectionImpl) Close() error {
	c.logger.Info("closing rabbitmq connection")
	return c.Connection.Close()
}

func (c *amqpConnectionImpl) NewChannel() (*amqp.Channel, error) {
	c.logger.Info("creating new rabbitmq channel")
	return c.Connection.Channel()
}

func (c *amqpConnectionImpl) Raw() *amqp.Connection {
	return c.Connection
}

func (c *amqpConnectionImpl) connect() error {
	var err error
	if c.Connection, err = c.connectRabbitMQ(); err != nil {
		return err
	}
	// register connection close event
	c.connClose = c.Connection.NotifyClose(make(chan *amqp.Error))
	go c.reConnect()
	return nil
}

func (c *amqpConnectionImpl) reConnect() {
	var err error
	for amqpErr := range c.connClose {
		// only reconnect if the error is not nil
		if amqpErr == nil {
			continue
		}
		c.logger.Errorf("rabbitmq connection closed, error: %v", amqpErr)
		c.errorEvent <- amqpErr

		backoff := time.Duration(c.config.ReconnectDelay) * time.Millisecond
		maxBackoff := 30 * time.Second
		for {
			c.logger.Info("attempting to reconnect to rabbitmq...")
			if c.Connection, err = c.connectRabbitMQ(); err != nil {
				c.logger.Errorf("failed to reconnect to rabbitmq: %v", err)
				backoff = min(time.Duration(float64(backoff)*1.5), maxBackoff)
				time.Sleep(backoff)
				continue
			}
			c.logger.Info("successfully reconnected to rabbitmq")
			c.reconnectedEvent <- struct{}{}
			break
		}
	}
}

func (c *amqpConnectionImpl) connectRabbitMQ() (*amqp.Connection, error) {
	dialConfig := amqp.Config{
		Dial:      amqp.DefaultDial(30 * time.Second),
		Heartbeat: 10 * time.Second,
	}
	conn, err := amqp.DialConfig(c.config.AmqpEndPoint(), dialConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
