package sec

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	var s string
	var last string

	var err error

	for i := 1; i < 500; i += 10 {
		s, err = GenerateRandomString(i)
		require.NoError(t, err)
		assert.Equal(t, i, len(s))
		assert.False(t, strings.ContainsAny(s, "%&*/\\|	=+ "))
		assert.False(t, s == last)

		last = s
	}

}
