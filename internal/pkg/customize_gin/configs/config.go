package customizeginconfigs

import (
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/gin-contrib/cors"
)

type GinConfig interface {
	GetConnWebPort() string
	GenerateCORSConfig() cors.Config
	GetServiceName() string
	GetVersion() string
	GetBasePath() string
	GetEnvironment() enums.Environment
}
