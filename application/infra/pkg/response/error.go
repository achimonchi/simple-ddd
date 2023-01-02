package response

import (
	"net/http"
)

type CustomError struct {
	BaseResponse
	Error string `json:"error"`
}

func (c *CustomError) BadRequest() *CustomError {
	c.StatusCode = http.StatusBadRequest
	return c
}

func (c *CustomError) NotFound() *CustomError {
	c.StatusCode = http.StatusNotFound
	return c
}
func (c *CustomError) Unauthorized() *CustomError {
	c.StatusCode = http.StatusUnauthorized
	return c
}
func (c *CustomError) InternalServerError() *CustomError {
	c.StatusCode = http.StatusNotFound
	return c
}

func (c *CustomError) WithMessage(msg string) *CustomError {
	c.Error = msg
	return c
}
func (c *CustomError) WithInfo(usecase, info string) *CustomError {
	c.AdditionalInfo = &AdditionalInfo{
		Usecase: usecase,
		Info:    info,
	}
	return c
}

func GeneralError() *CustomError {
	return &CustomError{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    string(MSG_GeneralError),
		},
		Error: string(MSG_GeneralError),
	}
}

func RegisterError() *CustomError {
	return &CustomError{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    string(MSG_RegisterError),
		},
		Error: string(MSG_RegisterError),
	}
}

func LoginError() *CustomError {
	return &CustomError{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    string(MSG_LoginError),
		},
		Error: string(MSG_LoginError),
	}
}
func CreateProfileError() *CustomError {
	return &CustomError{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    string(MSG_CreateProfileError),
		},
		Error: string(MSG_LoginError),
	}
}
func GetProfileError() *CustomError {
	return &CustomError{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    string(MSG_GetProfileError),
		},
		Error: string(MSG_LoginError),
	}
}
