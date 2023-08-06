package gen

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomBytesSet returns set of random byte slices of latin symbols and/or numbers.
//
// It is NOT crypto safe and is meant to be used in tests or non-sensitive operations only.
func GenerateRandomBytesSet(lens []int) [][]byte {
	var ret [][]byte
	var s []byte
	for _, l := range lens {
		s = GenerateRandomBytes(l)
		ret = append(ret, s)
	}

	return ret
}

// GenerateRandomBytes returns a random byte slice of latin symbols and/or numbers.
//
// It is NOT crypto safe and is meant to be used in tests or non-sensitive operations only.
func GenerateRandomBytes(l int) []byte {
	const symbols = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	ret := make([]byte, l)
	for i := 0; i < l; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return nil
		}
		ret[i] = symbols[num.Int64()]
	}

	return ret
}
