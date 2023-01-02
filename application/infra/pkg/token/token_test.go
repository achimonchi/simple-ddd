package token

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var id = "3709480d-ca06-460a-9ac1-1c235605714d"

func TestGenerateToken(t *testing.T) {
	tok, err := GenerateToken(id)
	log.Println("token :", tok)
	require.Nil(t, err)
	require.NotNil(t, tok)
}

func TestVerifyToken(t *testing.T) {
	tokString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7IklkIjoiMzcwOTQ4MGQtY2EwNi00NjBhLTlhYzEtMWMyMzU2MDU3MTRkIiwiRXhwaXJlZCI6IjIwMjMtMDEtMDJUMTE6Mjg6MzEuNjY3NDMrMDc6MDAifX0.lEQj8GsbhtzhIU4IkuNhRYwFiWVTOQYuI4p-hqSapCE"

	tok, err := ValidateToken(tokString)
	require.Nil(t, err)
	require.NotNil(t, tok)

	require.Equal(t, id, tok.Id)
}
