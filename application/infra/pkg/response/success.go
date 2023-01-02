package response

import "net/http"

type CustomSuccess struct {
	BaseResponse
	Payload interface{} `json:"payload,omitempty"`
}

func RegisterSuccess() *CustomSuccess {
	return &CustomSuccess{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusCreated,
			Message:    string(MSG_RegisterSuccess),
		},
	}
}
func LoginSuccess(payload interface{}) *CustomSuccess {
	return &CustomSuccess{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusCreated,
			Message:    string(MSG_LoginSuccess),
		},
		Payload: payload,
	}
}
