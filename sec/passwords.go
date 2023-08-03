package sec

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPasswordString(pwd string) (string, error) {
	return HashAndSaltPassword([]byte(pwd))
}

func HashAndSaltPassword(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided if it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("[HashAndSaltPassword] %w", err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), err
}

// ComparePasswords compares hash with password
//
// It is named this way for compatibility with my old projects
func ComparePasswords(hashedPwd string, plainPwd string) (bool, error) {
	return ComparePasswordBytes([]byte(hashedPwd), []byte(plainPwd))
}

// ComparePasswordBytes compares hash with password
func ComparePasswordBytes(byteHash []byte, plainPwdBytes []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwdBytes)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, fmt.Errorf("[ComparePasswords] %w", err)
	}

	return true, nil
}
