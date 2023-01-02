package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePassword(t *testing.T) {
	pass := "Hello"
	hash, err := GeneratePassword(pass)
	require.Nil(t, err)
	require.NotNil(t, hash)
}

func TestVerifyPassword(t *testing.T) {
	pass := "Hello"
	hash, err := GeneratePassword(pass)
	require.Nil(t, err)
	require.NotNil(t, hash)

	err = VerifyPassword(hash, pass)
	require.Nil(t, err)

}
