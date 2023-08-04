package hasher

import (
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileHashSameFile(t *testing.T) {
	sameText := "This is some file here"
	file := mock.MockWriteReader{
		Bytes: []byte(sameText),
	}
	file2 := mock.MockWriteReader{
		Bytes: []byte(sameText),
	}

	hash, err := HashFile(&file, MD5, 50)
	require.NoError(t, err)
	hash2, err := HashFile(&file2, MD5, 15)
	require.NoError(t, err)
	assert.Equal(t, hash, hash2)

	hash, err = HashFile(&file, SHA1, 50)
	require.NoError(t, err)
	hash2, err = HashFile(&file2, SHA1, 15)
	require.NoError(t, err)
	assert.Equal(t, hash, hash2)

	hash, err = HashFile(&file, SHA256, 50)
	require.NoError(t, err)
	hash2, err = HashFile(&file2, SHA256, 15)
	require.NoError(t, err)
	assert.Equal(t, hash, hash2)

	hash, err = HashFile(&file, SHA512, 50)
	require.NoError(t, err)
	hash2, err = HashFile(&file2, SHA512, 15)
	require.NoError(t, err)
	assert.Equal(t, hash, hash2)
}
