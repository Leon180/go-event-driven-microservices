package customizegin

import "github.com/gin-gonic/gin"

type GinMiddleware interface {
	Handle() gin.HandlerFunc
}
