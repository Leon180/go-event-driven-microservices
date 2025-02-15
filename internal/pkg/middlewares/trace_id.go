package middlewares

import (
	"bytes"
	"context"
	"io"
	"slices"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type traceIDMiddleware struct {
	uuidGenerator utilities.UUIDGenerator
}

func NewTraceIDMiddleware(uuidGenerator utilities.UUIDGenerator) *traceIDMiddleware {
	return &traceIDMiddleware{uuidGenerator: uuidGenerator}
}

func (middleware *traceIDMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID := middleware.uuidGenerator.GenerateUUID()
		c.Set(enums.TraceIDKey.ToString(), eventID)
		if contextType := c.Request.Header.Get(enums.RequestHeaderContentType.ToString()); slices.Contains(enums.ContextTypeGroupDefault.GetSlice().ToStringSlice(), contextType) {
			buf, err := io.ReadAll(c.Request.Body)
			if err != nil {
				utilities.LogError(c.Request.Context(), err)
				c.Next()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(buf))
			utilities.LogInfo(c.Request.Context(), "Http Request: %+v, EventID: %s, Body: %s", c.Request, eventID, string(buf))
		} else {
			utilities.LogInfo(c.Request.Context(), "Http Request: %+v, EventID: %s", c.Request, eventID)
		}
		c.Next()
	}
}

type traceIDInterceptor struct {
	uuidGenerator utilities.UUIDGenerator
}

func NewTraceIDInterceptor(uuidGenerator utilities.UUIDGenerator) *traceIDInterceptor {
	return &traceIDInterceptor{uuidGenerator: uuidGenerator}
}

func (interceptor *traceIDInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		traceID := interceptor.uuidGenerator.GenerateUUID()
		ctx = context.WithValue(ctx, enums.TraceIDKey, traceID)
		utilities.LogInfo(ctx, "Http Request: %+v, EventID: %s", req, traceID)
		return handler(ctx, req)
	}
}
