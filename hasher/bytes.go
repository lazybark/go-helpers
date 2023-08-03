package hasher

import (
	"fmt"
)

// HashBytes returns hash of []byte b
func HashBytes(b []byte, ht HashType) string {
	return fmt.Sprintf("%x", GetHasher(ht).Sum(b))
}
