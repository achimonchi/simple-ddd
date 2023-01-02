package infrahttp

import "github.com/gin-gonic/gin"

type Middleware interface {
	ValidateAuth(ctx *gin.Context)
}
