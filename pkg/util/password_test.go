package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyPassowrd(t *testing.T) {
	password := RandomString(8)
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = VerifyPassword(password, hashedPassword)
	require.NoError(t, err)
}
