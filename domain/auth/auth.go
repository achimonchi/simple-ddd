package auth

type Auth struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
