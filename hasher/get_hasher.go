package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func GetHasher(ht HashType) hash.Hash {
	if ht == MD5 {
		return md5.New()
	} else if ht == SHA1 {
		return sha1.New()
	} else if ht == SHA256 {
		return sha256.New()
	} else if ht == SHA512 {
		return sha512.New()
	}

	return nil
}
