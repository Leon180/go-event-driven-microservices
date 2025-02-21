package customizeginendpoints

import (
	"fmt"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Endpoint interface {
	Handle(c *gin.Context)
	MapEndpoint(router *gin.RouterGroup)
}

func FxTagEndpoint(handler interface{}) interface{} {
	return fx.Annotate(
		handler,
		fx.As(new(Endpoint)),
		fx.ResultTags(fmt.Sprintf(`group:"%s"`, enums.FxGroupEndpoints.ToString())),
	)
}
