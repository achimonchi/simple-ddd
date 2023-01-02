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

func (c *CustomError) NotFound() *CustomError {
	c.StatusCode = http.StatusNotFound
	return c
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
