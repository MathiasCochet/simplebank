package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordIsCorrect(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotNil(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)
}

func TestPasswordIsWrong(t *testing.T) {
	password := RandomString(6)
	hashedPassword, err := HashPassword(password)
	wrongPassword := RandomString(6)

	err = CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}

func TestPasswordIsUnique(t *testing.T) {
	password := RandomString(6)
	hashedPassword1, _ := HashPassword(password)
	hashedPassword2, _ := HashPassword(password)

	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
