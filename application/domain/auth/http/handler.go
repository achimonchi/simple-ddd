package domainhttp

import (
	"net/http"
	"poc-ddd/application/domain/auth/param"
	"poc-ddd/application/infra/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase Usecase
}

func NewAuthHandler(usecase Usecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (a *AuthHandler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var req = new(param.RegisterRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.RegisterError().BadRequest().WithMessage(err.Error()).WithInfo("Register", err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	custErr := a.usecase.RegisterAccount(ctx.Request.Context(), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.RegisterSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var req = new(param.LoginRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.LoginError().BadRequest().WithMessage(err.Error()).WithInfo("Register", err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	respLogin, custErr := a.usecase.Login(ctx.Request.Context(), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.LoginSuccess(respLogin)
	ctx.JSON(resp.StatusCode, resp)
}
