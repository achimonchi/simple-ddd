package infrahttp

import (
	"poc-ddd/application/infra/pkg/response"
	"poc-ddd/application/infra/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) ValidateAuth(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	tokenString := strings.Split(auth, "Bearer ")
	if len(tokenString) != 2 {
		resp := response.GeneralError().
			Unauthorized().
			WithMessage("try to get token string from header").
			WithInfo("MiddlewareValidateAuth", "len token must be 2")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	tok, err := token.ValidateToken(tokenString[1])

	if err != nil {
		resp := response.GeneralError().
			Unauthorized().
			WithMessage("try to validate token").
			WithInfo("MiddlewareValidateAuth", err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	ctx.Set("authId", tok.Id)
	ctx.Next()
}
