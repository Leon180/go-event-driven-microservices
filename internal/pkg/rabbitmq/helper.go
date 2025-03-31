package rabbitmq

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConvertDeliveryToPublishing(d *amqp.Delivery, newTimeStamp bool) amqp.Publishing {
	if d == nil {
		return amqp.Publishing{}
	}
	return amqp.Publishing{
		Headers:         d.Headers,
		ContentType:     d.ContentType,
		ContentEncoding: d.ContentEncoding,
		DeliveryMode:    d.DeliveryMode,
		Priority:        d.Priority,
		CorrelationId:   d.CorrelationId,
		ReplyTo:         d.ReplyTo,
		Expiration:      d.Expiration,
		MessageId:       d.MessageId,
		Timestamp: func() time.Time {
			if newTimeStamp {
				return time.Now()
			}
			return d.Timestamp
		}(),
		Type:   d.Type,
		UserId: d.UserId,
		AppId:  d.AppId,
		Body:   d.Body,
	}
}

type PublishingHeaderSetter interface {
	SetError(err error) PublishingHeaderSetter
	SetExchange(exchange string) PublishingHeaderSetter
	SetRoutingKey(routingKey string) PublishingHeaderSetter
	SetQueue(queue string) PublishingHeaderSetter
	SetRetryCount(retryCount int) PublishingHeaderSetter
	Build() *amqp.Publishing
}

func NewPublishingHeaderSetter(publishing *amqp.Publishing) PublishingHeaderSetter {
	return &publishingHeaderSetter{
		publishing: publishing,
	}
}

type publishingHeaderSetter struct {
	publishing *amqp.Publishing
}

func (p *publishingHeaderSetter) SetError(err error) PublishingHeaderSetter {
	p.publishing.Headers[enums.DeliveryHeaderError.ToString()] = err.Error()
	return p
}

func (p *publishingHeaderSetter) SetExchange(exchange string) PublishingHeaderSetter {
	p.publishing.Headers[enums.DeliveryHeaderExchange.ToString()] = exchange
	return p
}

func (p *publishingHeaderSetter) SetRoutingKey(routingKey string) PublishingHeaderSetter {
	p.publishing.Headers[enums.DeliveryHeaderRoutingKey.ToString()] = routingKey
	return p
}

func (p *publishingHeaderSetter) SetQueue(queue string) PublishingHeaderSetter {
	p.publishing.Headers[enums.DeliveryHeaderQueue.ToString()] = queue
	return p
}

func (p *publishingHeaderSetter) SetRetryCount(retryCount int) PublishingHeaderSetter {
	p.publishing.Headers[enums.DeliveryHeaderRetryCount.ToString()] = retryCount
	return p
}

func (p *publishingHeaderSetter) Build() *amqp.Publishing {
	return p.publishing
}
