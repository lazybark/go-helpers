package mock

import (
	"errors"
	"io"
	"testing"

	"github.com/lazybark/go-helpers/sec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Testing testing package here

func TestFileRead(t *testing.T) {
	rnd, err := sec.GenerateRandomString(50)
	require.NoError(t, err)

	mf := &MockWriteReader{
		Bytes: []byte(rnd),
	}

	//Empty slice = empty string read
	var b []byte
	n, err := mf.Read(b)
	require.NoError(t, err)
	assert.Equal(t, 0, n)
	assert.Empty(t, string(b))

	//Now we need 2 reads before EOF
	b = make([]byte, len(mf.Bytes)/2)
	n, err = mf.Read(b)
	require.NoError(t, err)
	assert.Equal(t, len(mf.Bytes)/2, n)
	assert.Equal(t, string(mf.Bytes[:len(mf.Bytes)/2]), string(b))

	n, err = mf.Read(b)
	assert.True(t, errors.Is(io.EOF, err))
	assert.Equal(t, len(mf.Bytes)/2, n)
	assert.Equal(t, string(mf.Bytes[len(mf.Bytes)/2:]), string(b))

	n, err = mf.Read(b)
	assert.True(t, errors.Is(io.EOF, err))
	assert.Empty(t, n)

	mf.SetLastRead(1) //Set buffer to some value

	b = make([]byte, len(mf.Bytes)-1)
	n, err = mf.Read(b)
	assert.True(t, errors.Is(io.EOF, err))
	assert.Equal(t, len(mf.Bytes)-1, n)
	assert.Equal(t, string(mf.Bytes[1:]), string(b))

	n, err = mf.Read(b)
	assert.True(t, errors.Is(io.EOF, err))
	assert.Empty(t, n)

	//Should return EOF if asked to
	mf.SetLastRead(0)
	mf.ReturnEOF = true
	n, err = mf.Read(b)
	assert.True(t, errors.Is(io.EOF, err))
	assert.Empty(t, n)
}
