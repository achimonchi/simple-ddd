package auth

import (
	"errors"

	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Id       ulid.ULID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

var (
	ErrEmailEmpty     = errors.New("email cannot be empty")
	ErrPasswordEmpty  = errors.New("password cannot be empty")
	ErrPasswordLength = errors.New("password length must be greater than 6")

	ErrEmailNotFound   = errors.New("email not found")
	ErrInvalidPassword = errors.New("invalid password")
)

func NewAuth() Auth {
	return Auth{}
}

func (a *Auth) EncryptPassword() (err error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(encrypted)
	return
}

func (a Auth) FromRegister(req register) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) <= 6 {
		return a, ErrPasswordLength
	}

	a.Email = req.Email
	a.Password = req.Password
	a.Id = ulid.Make()
	return a, nil
}

func (a Auth) FromLogin(req login) (Auth, error) {
	if req.Email == "" {
		return a, ErrEmailEmpty
	}

	if req.Password == "" {
		return a, ErrPasswordEmpty
	}

	if len(req.Password) <= 6 {
		return a, ErrPasswordLength
	}

	a.Email = req.Email
	a.Password = req.Password
	return a, nil
}

func (a Auth) ValidatePasswordFromPlainText(plainText string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plainText))
	if err != nil {
		return ok, ErrInvalidPassword
	}
	ok = true
	return
}
