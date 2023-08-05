package gen

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomStringSet returns set of random strings of latin symbols and numbers.
//
// It is NOT crypto safe and is meant to be used in tests or non-sensitive operations only
func GenerateRandomStringSet(lens []int) []string {
	var ret []string
	var s string
	for _, l := range lens {
		s = GenerateRandomString(l)
		ret = append(ret, s)
	}

	return ret
}

// GenerateRandomString returns a random string of latin symbols and numbers.
//
// It is NOT crypto safe and is meant to be used in tests or non-sensitive operations only
func GenerateRandomString(l int) string {
	const symbols = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	ret := make([]byte, l)
	for i := 0; i < l; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return ""
		}
		ret[i] = symbols[num.Int64()]
	}

	return string(ret)
}
