package customizegin

import (
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	Handle(c *gin.Context)
	MapEndpoint(router *gin.RouterGroup)
}
