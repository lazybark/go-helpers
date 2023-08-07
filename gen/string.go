package gen

// GenerateRandomStringSet returns set of random strings of latin symbols and/or numbers.
//
// It is NOT cryptographically secure and is meant to be used in tests or non-sensitive operations only
func GenerateRandomStringSet(lens []int) []string {
	var ret []string
	var s string
	for _, l := range lens {
		s = GenerateRandomString(l)
		ret = append(ret, s)
	}

	return ret
}

// GenerateRandomStringSetFromSet returns set of random strings made from charSet you provide.
//
// It is NOT cryptographically secure and is meant to be used in tests or non-sensitive operations only
func GenerateRandomStringSetFromSet(lens []int, charSet []byte) []string {
	var ret []string
	var s string
	for _, l := range lens {
		s = GenerateRandomStringFromSet(l, charSet)
		ret = append(ret, s)
	}

	return ret
}

// GenerateRandomString returns a random string of latin symbols and/or numbers.
//
// It is NOT cryptographically secure and is meant to be used in tests or non-sensitive operations only
func GenerateRandomString(n int) string {
	return GenerateRandomStringFromSet(n, []byte(DigitsAndEnglish))
}

// GenerateRandomStringFromSet returns a random string of symbols you provide in charSet.
//
// It is NOT cryptographically secure and is meant to be used in tests or non-sensitive operations only
func GenerateRandomStringFromSet(n int, charSet []byte) string {
	return string(GenerateRandomBytesFromSet(n, charSet))
}
