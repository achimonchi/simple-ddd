package domainhttp

import (
	"poc-ddd/application/domain/auth"
	"poc-ddd/application/domain/auth/repository"
	"poc-ddd/application/domain/auth/usecase"

	"github.com/gin-gonic/gin"
)

type RouterHttp struct {
	router *gin.RouterGroup
	auth   *AuthHandler
}

// RegisterRoute implements auth.Auth
func (r *RouterHttp) RegisterRoute() {
	route := r.router.Group("auth")
	{
		route.GET("/health", r.auth.Ping)
		route.POST("/register", r.auth.Register)
		route.POST("/login", r.auth.Login)
	}
}

func NewRouterHttp(r *gin.RouterGroup) auth.Auth {
	repo := repository.NewRepoFactory().Create(repository.REPO_Memory)
	usecase := usecase.NewUsecase(repo)
	auth := NewAuthHandler(usecase)
	return &RouterHttp{
		router: r,
		auth:   auth,
	}
}
