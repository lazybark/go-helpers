package hasher

import (
	"fmt"
)

// HashString returns hash of string s
func HashString(s string, ht HashType) string {
	return fmt.Sprintf("%x", GetHasher(ht).Sum([]byte(s)))
}
