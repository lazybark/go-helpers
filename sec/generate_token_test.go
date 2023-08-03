package sec

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	s, err := GenerateRandomString(1)
	require.NoError(t, err)
	assert.Equal(t, 1, len(s))

	s, err = GenerateRandomString(2)
	require.NoError(t, err)
	assert.Equal(t, 2, len(s))

	s, err = GenerateRandomString(17)
	require.NoError(t, err)
	assert.Equal(t, 17, len(s))

	s, err = GenerateRandomString(150)
	require.NoError(t, err)
	assert.Equal(t, 150, len(s))

	assert.Equal(t, false, strings.ContainsAny(s, "%&*/\\|	=+ "))
}
