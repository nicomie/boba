package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPin(t *testing.T) {
	pin := RandomString(4)

	hashedPin, err := HashPin(pin)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPin)

	err = CheckPin(pin, hashedPin)

	require.NoError(t, err)

	wrongPin := RandomString(4)
	err = CheckPin(wrongPin, hashedPin)

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPin2, err := HashPin(pin)
	require.NoError(t, err)
	require.NotEqual(t, hashedPin, hashedPin2)

}
