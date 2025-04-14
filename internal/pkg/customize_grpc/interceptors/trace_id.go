package customizegrpcinterceptors

import (
	"context"

	customizegrpc "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_grpc"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"google.golang.org/grpc"
)

func NewTraceIDUnaryInterceptor(
	uuidGenerator uuid.UUIDGenerator,
	logger loggers.Logger,
) customizegrpc.GRPCUnaryInterceptor {
	return &traceIDInterceptor{
		uuidGenerator: uuidGenerator,
		logger:        logger,
	}
}

type traceIDInterceptor struct {
	uuidGenerator uuid.UUIDGenerator
	logger        loggers.Logger
}

func (middleware *traceIDInterceptor) Handle() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		eventID := middleware.uuidGenerator.GenerateUUID()
		ctx = context.WithValue(ctx, enums.ContextKeyTraceID, eventID)
		middleware.logger.Info("GRPC Request: %+v, EventID: %s", req, eventID)
		return handler(ctx, req)
	}
}
