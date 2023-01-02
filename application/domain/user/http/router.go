package infrahttp

import (
	"net/http"
	authRepo "poc-ddd/application/domain/auth/repository"
	"poc-ddd/application/domain/user"
	"poc-ddd/application/domain/user/repository"
	"poc-ddd/application/domain/user/usecase"
	"poc-ddd/application/infra/adapter"

	"github.com/gin-gonic/gin"
)

type RouterHttp struct {
	router     *gin.RouterGroup
	user       *UserHandler
	middleware Middleware
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

		route.Use(r.middleware.ValidateAuth)
		route.POST("/profile", r.user.CreateProfile)
		route.GET("/profile", r.user.GetProfile)
	}
}

func NewRouterHttp(r *gin.RouterGroup, middleware Middleware) user.User {
	authRepo := authRepo.NewRepoFactory().Create(authRepo.REPO_Memory)
	authAdapter := adapter.NewAuthAdapter(authRepo)
	repo := repository.NewRepoFactory().Create(repository.REPO_Memory)
	usecase := usecase.NewUsecase().SetRepo(repo).SetAuthAdapter(authAdapter).Build()
	user := NewUserHandler(usecase)
	return &RouterHttp{
		router:     r,
		user:       user,
		middleware: middleware,
	}
}
