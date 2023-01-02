package param

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
