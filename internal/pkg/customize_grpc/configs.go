package customizegrpc

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"google.golang.org/grpc/keepalive"
)

type GRPCConfig interface {
	GetHost() string
	GetPort() string
	GetServiceName() string
	GetVersion() string
	GetBasePath() string
	GetEnvironment() enums.Environment

	GetKeepAliveParams() keepalive.ServerParameters
}
