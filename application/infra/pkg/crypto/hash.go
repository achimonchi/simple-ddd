package crypto

import "golang.org/x/crypto/bcrypt"

const (
	COST = bcrypt.DefaultCost
)

func GeneratePassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), COST)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
