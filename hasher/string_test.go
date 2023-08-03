package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashString(t *testing.T) {
	s := "some+string_here"

	hash := HashString(s, MD5)
	hash2 := HashString(s, MD5)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashString(s, SHA1)
	hash2 = HashString(s, SHA1)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashString(s, SHA256)
	hash2 = HashString(s, SHA256)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)

	hash = HashString(s, SHA512)
	hash2 = HashString(s, SHA512)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, hash2)
	assert.Equal(t, hash, hash2)
}
