package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashBytes(t *testing.T) {
	s := []byte("some+string_here")

	hash := HashBytes(s, MD5)
	hash2 := HashBytes(s, MD5)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashBytes(s, SHA1)
	hash2 = HashBytes(s, SHA1)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashBytes(s, SHA256)
	hash2 = HashBytes(s, SHA256)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashBytes(s, SHA512)
	hash2 = HashBytes(s, SHA512)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)
}
