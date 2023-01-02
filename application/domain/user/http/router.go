package infrahttp

import (
	"net/http"
	"poc-ddd/application/domain/user"

	"github.com/gin-gonic/gin"
)

type RouterHttp struct {
	router *gin.RouterGroup
}

// RegisterRoute implements user.User
func (r *RouterHttp) RegisterRoute() {
	route := r.router.Group("users")
	{
		route.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}
}

func NewRouterHttp(r *gin.RouterGroup) user.User {
	return &RouterHttp{
		router: r,
	}
}
