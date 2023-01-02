package infrahttp

import (
	"net/http"
	"poc-ddd/application/domain/user/param"
	"poc-ddd/application/infra/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase Usecase
}

func NewUserHandler(usecase Usecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (u *UserHandler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (u *UserHandler) CreateProfile(ctx *gin.Context) {
	var req = new(param.CreateProfileRequest)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.RegisterError().BadRequest().WithMessage(err.Error()).WithInfo("Register", err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	authId := ctx.GetString("authId")
	req.AuthId = authId

	custErr := u.usecase.CreateProfile(ctx, req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreateProfileSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (u *UserHandler) GetProfile(ctx *gin.Context) {
	authId := ctx.GetString("authId")

	profile, custErr := u.usecase.GetProfile(ctx.Request.Context(), authId)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}
	resp := response.GetProfileSuccess(profile)
	ctx.JSON(resp.StatusCode, resp)
}
