package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

// HashString returns hash of string s
func HashString(s string, ht HashType) string {
	var h hash.Hash

	if ht == MD5 {
		h = md5.New()
	} else if ht == SHA1 {
		h = sha1.New()
	} else if ht == SHA256 {
		h = sha256.New()
	} else if ht == SHA512 {
		h = sha512.New()
	}

	return fmt.Sprintf("%x", h.Sum([]byte(s)))
}
