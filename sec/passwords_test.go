package sec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPasswordAndComparePasswordWithHash(t *testing.T) {
	pwd := "some_password"
	pwd2 := "some_password"
	hash, err := HashAndSaltPasswordString(pwd, 10)
	require.NoError(t, err)

	assert.Greater(t, len(hash), 0)
	assert.Greater(t, len(hash), len(pwd))
	assert.NotEqual(t, pwd, hash)

	pwd3 := "some_password2"
	hash3, err := HashAndSaltPasswordString(pwd3, 10)
	require.NoError(t, err)
	assert.Greater(t, len(hash3), 0)
	assert.Greater(t, len(hash3), len(pwd2))
	assert.NotEqual(t, pwd2, hash3)

	//Now we compare

	yes, err := ComparePasswords(hash, pwd)
	require.NoError(t, err)
	assert.Equal(t, true, yes)

	yes, err = ComparePasswords(hash, pwd2)
	require.NoError(t, err)
	assert.Equal(t, true, yes)

	yes, err = ComparePasswords(hash, pwd3)
	require.NoError(t, err)
	assert.NotEqual(t, true, yes)

	yes, err = ComparePasswords(hash3, pwd3)
	require.NoError(t, err)
	assert.Equal(t, true, yes)

	yes, err = ComparePasswords(hash3, pwd2)
	require.NoError(t, err)
	assert.NotEqual(t, true, yes)

	yes, err = ComparePasswords(hash3, pwd)
	require.NoError(t, err)
	assert.NotEqual(t, true, yes)
}
