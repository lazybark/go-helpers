package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t, "MD5", MD5.String())
	assert.Equal(t, "SHA1", SHA1.String())
	assert.Equal(t, "SHA256", SHA256.String())
	assert.Equal(t, "SHA512", SHA512.String())
	assert.Equal(t, "", HashType(9).String())
	assert.Equal(t, "", HashType(-39).String())

	assert.True(t, MD5.CheckType())
	assert.True(t, SHA1.CheckType())
	assert.True(t, SHA256.CheckType())
	assert.True(t, SHA512.CheckType())
	assert.False(t, HashType(19).CheckType())
	assert.False(t, HashType(-3).CheckType())
}
