package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	p := RandomString(6)
	h, err := HashPassword(p)
	require.NoError(t, err)
	require.NotEmpty(t, h)

	err = CheckPassword(p, h)
	require.NoError(t, err)

	//wrongPassword
	wP := RandomString(6)
	err = CheckPassword(wP, h)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}