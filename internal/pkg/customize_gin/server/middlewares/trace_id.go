package customizeginmiddlewares

import (
	"bytes"
	"io"
	"slices"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"

	"github.com/gin-gonic/gin"
)

func NewTraceIDMiddleware(uuidGenerator uuid.UUIDGenerator, logger loggers.Logger) GinMiddleware {
	return &traceIDMiddleware{uuidGenerator: uuidGenerator, logger: logger}
}

type traceIDMiddleware struct {
	uuidGenerator uuid.UUIDGenerator
	logger        loggers.Logger
}

func (middleware *traceIDMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID := middleware.uuidGenerator.GenerateUUID()
		c.Set(enums.ContextKeyTraceID.ToString(), eventID)
		if contextType := c.Request.Header.Get(enums.RequestHeaderContentType.ToString()); slices.Contains(enums.ContextTypeGroupDefault.GetSlice().ToStringSlice(), contextType) {
			buf, err := io.ReadAll(c.Request.Body)
			if err != nil {
				middleware.logger.Err("error while reading request body", err)
				c.Next()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(buf))
			middleware.logger.Info("Http Request: %+v, EventID: %s, Body: %s", c.Request, eventID, string(buf))
		} else {
			middleware.logger.Info("Http Request: %+v, EventID: %s", c.Request, eventID)
		}
		c.Next()
	}
}
