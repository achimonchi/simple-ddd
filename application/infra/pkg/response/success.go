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
			StatusCode: http.StatusOK,
			Message:    string(MSG_LoginSuccess),
		},
		Payload: payload,
	}
}
func GetProfileSuccess(payload interface{}) *CustomSuccess {
	return &CustomSuccess{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusOK,
			Message:    string(MSG_GetProfileSuccess),
		},
		Payload: payload,
	}
}
func CreateProfileSuccess() *CustomSuccess {
	return &CustomSuccess{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusCreated,
			Message:    string(MSG_CreateProfileSuccess),
		},
	}
}
