package hasher

type (
	HashType int
)

// Legal hash types
const (
	hash_types_start HashType = iota

	MD5
	SHA1
	SHA256
	SHA512

	hash_types_end
)

// String returns name for hash type or "illegal" if it is wrong
func (t HashType) String() string {
	if !t.CheckType() {
		return "illegal"
	}
	return [...]string{"illegal", "MD5", "SHA1", "SHA256", "SHA512", "illegal"}[t]
}

// CheckType returns false if t has illegal hash type
func (t *HashType) CheckType() bool {
	return hash_types_start < *t && *t < hash_types_end
}
