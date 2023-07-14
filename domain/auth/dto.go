package auth

type register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
