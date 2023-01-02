package response

type BaseResponse struct {
	StatusCode     int             `json:"status_code"`
	Message        string          `json:"message"`
	AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
}

type AdditionalInfo struct {
	Usecase string `json:"usecase"`
	Info    string `json:"info"`
}

type Message string

const (
	MSG_CreateSuccess Message = "CREATE SUCCESS"
	MSG_CreateError   Message = "CREATE ERROR"

	MSG_RegisterSuccess Message = "REGISTER SUCCESS"
	MSG_RegisterError   Message = "REGISTER ERROR"

	MSG_LoginSuccess Message = "LOGIN SUCCESS"
	MSG_LoginError   Message = "LOGIN ERROR"
)
