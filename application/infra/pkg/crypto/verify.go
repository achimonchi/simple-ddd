package crypto

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(passHash, plainPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(plainPass))
}
