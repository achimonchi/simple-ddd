package infrahttp

import (
	authRoute "poc-ddd/application/domain/auth/http"
	userRoute "poc-ddd/application/domain/user/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	port   string
	router *gin.Engine
}

// Run implements infra.Infra
func (r *Router) Run() {
	v1 := r.router.Group("v1")
	middleware := NewMiddleware()

	// register all routes
	rAuth := authRoute.NewRouterHttp(v1)
	rAuth.RegisterRoute()

	rUser := userRoute.NewRouterHttp(v1, middleware)
	rUser.RegisterRoute()

	r.router.Run(r.port)
}

func NewRouter(port string) *Router {
	router := gin.Default()

	return &Router{
		port:   port,
		router: router,
	}
}
