package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomBytes(t *testing.T) {
	b1 := GenerateRandomBytes(50)
	b2 := GenerateRandomBytes(50)
	b3 := GenerateRandomBytes(23)

	assert.NotEqual(t, b1, b2)
	assert.Equal(t, 50, len(b1))
	assert.Equal(t, 50, len(b2))
	assert.Equal(t, 23, len(b3))
}
