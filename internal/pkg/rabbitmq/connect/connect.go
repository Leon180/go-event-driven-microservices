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
	IsConnect() bool
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
	logger loggers.Logger
	*amqp.Connection
	config           *rabbitmq.RabbitMQConfig
	isConnect        bool
	reconnectedEvent chan struct{}
	errorEvent       chan error
}

func (c *amqpConnectionImpl) IsConnect() bool {
	return c.isConnect
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
	conn, err := c.connectRabbitMQ()
	if err != nil {
		return err
	}
	c.Connection = conn
	c.isConnect = true

	// register connection close event
	connClose := c.Connection.NotifyClose(make(chan *amqp.Error))
	go func() {
		for amqpErr := range connClose {
			if amqpErr == nil {
				continue
			}
			c.logger.Errorf("rabbitmq connection closed, error: %v", amqpErr)
			c.isConnect = false
			c.errorEvent <- amqpErr

			// Implement exponential backoff for reconnection
			backoff := time.Duration(c.config.ReconnectDelay) * time.Millisecond
			maxBackoff := 30 * time.Second
			for {
				c.logger.Info("attempting to reconnect to rabbitmq...")
				conn, err = c.connectRabbitMQ()
				if err != nil {
					c.logger.Error("failed to reconnect to rabbitmq: %v", err)
					// Exponential backoff with max limit
					backoff = time.Duration(float64(backoff) * 1.5)
					backoff = min(backoff, maxBackoff)
					time.Sleep(backoff)
					continue
				}

				c.logger.Info("successfully reconnected to rabbitmq")
				c.Connection = conn
				c.isConnect = true
				c.reconnectedEvent <- struct{}{}
				break
			}
		}
	}()

	return nil
}

func (c *amqpConnectionImpl) connectRabbitMQ() (*amqp.Connection, error) {
	// Add connection timeout
	dialConfig := amqp.Config{
		Dial:      amqp.DefaultDial(30 * time.Second),
		Heartbeat: 10 * time.Second, // Send heartbeat every 10 seconds
	}

	conn, err := amqp.DialConfig(c.config.AmqpEndPoint(), dialConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}

	return conn, nil
}
