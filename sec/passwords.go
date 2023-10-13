package sec

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPasswordString(pwd string, cost int) (string, error) {
	return HashAndSaltPassword([]byte(pwd), cost)
}

// HashAndSaltPassword returns password hash with provided cost (values between 4 & 31).
func HashAndSaltPassword(pwd []byte, cost int) (string, error) {
	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	// Use GenerateFromPassword to hash & salt the password.
	//
	// DefaultCost = 10, but can be set up to 31 (which is the max).
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("[HashAndSaltPassword] %w", err)
	}
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
