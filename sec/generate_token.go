package sec

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	err := assertAvailablePRNG()
	if err != nil {
		return "", fmt.Errorf("[GenerateRandomString]%w", err)
	}

	const symbols = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-`
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", fmt.Errorf("[GenerateRandomString] %w", err)
		}
		ret[i] = symbols[num.Int64()]
	}

	return string(ret), nil
}

func assertAvailablePRNG() error {
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return fmt.Errorf("[assertAvailablePRNG] crypto/rand is unavailable: %w", err)
	}

	return nil
}
