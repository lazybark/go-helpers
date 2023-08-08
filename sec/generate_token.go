package sec

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/lazybark/go-helpers/gen"
)

// GenerateRandomString uses digits and english letters to generate cryptographically secure random string
func GenerateRandomString(n int) (string, error) {
	return GenerateRandomStringFromSet(n, []byte(gen.DigitsAndEnglish))
}

// GenerateRandomStringFromSet uses provided set of characters to generate cryptographically secure random string
func GenerateRandomStringFromSet(n int, charSet []byte) (string, error) {
	// assertAvailablePRNG is what differs this generator from one in gen.
	// We check that it's ok to use rand.Reader right now
	err := assertAvailablePRNG()
	if err != nil {
		return "", fmt.Errorf("[GenerateRandomString]%w", err)
	}

	return string(gen.GenerateRandomBytesFromSet(n, charSet)), nil
}

func assertAvailablePRNG() error {
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return fmt.Errorf("[assertAvailablePRNG] crypto/rand is unavailable: %w", err)
	}

	return nil
}
