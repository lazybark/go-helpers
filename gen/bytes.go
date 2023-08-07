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

func GenerateRandomBytesSetFromSet(lens []int, charSet []byte) [][]byte {
	var ret [][]byte
	var s []byte
	for _, l := range lens {
		s = GenerateRandomBytesFromSet(l, charSet)
		ret = append(ret, s)
	}

	return ret
}

// GenerateRandomBytes returns a random byte slice of latin symbols and/or numbers.
//
// It is NOT crypto safe and is meant to be used in tests or non-sensitive operations only.
func GenerateRandomBytes(n int) []byte {
	return GenerateRandomBytesFromSet(n, []byte(DigitsAndEnglish))
}

func GenerateRandomBytesFromSet(n int, charSet []byte) []byte {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return nil
		}
		ret[i] = charSet[num.Int64()]
	}

	return ret
}
