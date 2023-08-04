package csvw

import (
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSVBuilderWriteBuffer(t *testing.T) {
	f := &mock.MockFile{}

	sep := ";"
	b := NewCSVBuilder(sep)
	assert.Equal(t, sep, b.Separator)

	b.UseFile(f)

	//Add cell to buffer
	cell1 := "one"
	cell2 := "two"
	err := b.AddCell(cell1, cell2)
	require.NoError(t, err)
	assert.Equal(t, cell1+sep+cell2+sep, b.String())
	//NewLine
	err = b.NewLine()
	require.NoError(t, err)
	assert.Equal(t, cell1+sep+cell2+sep+"\n", b.String())
	//Add line
	line := "some new;line;here"
	err = b.AddLine(line)
	require.NoError(t, err)
	err = b.NewLine()
	require.NoError(t, err)
	assert.Equal(t, cell1+sep+cell2+sep+"\n"+line+sep+"\n", b.String())
	//Clean buffer
	b.Reset()
	assert.Empty(t, b.String())
}

func TestCSVBuilderWriteFile(t *testing.T) {
	f := &mock.MockWriteReader{}

	sep := ";"
	b := NewCSVBuilder(sep)
	assert.Equal(t, sep, b.Separator)

	b.UseFile(f)

	//Put buffer data into file
	cell1 := "one"
	cell2 := "two"
	err := b.AddCell(cell1, cell2)
	require.NoError(t, err)
	assert.Equal(t, cell1+sep+cell2+sep, b.String())
	_, err = b.WriteBuffer()
	require.NoError(t, err)
	assert.Empty(t, b.String())                           //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep, string(f.Bytes)) // And file is full
	//Write bytes to file
	cell3 := "three"
	_, err = b.Write([]byte(cell3))
	require.NoError(t, err)
	assert.Empty(t, b.String())                                 //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep+cell3, string(f.Bytes)) // And file is full
	//Write string to file
	_, err = b.WriteString(cell3)
	require.NoError(t, err)
	assert.Empty(t, b.String())                                       //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep+cell3+cell3, string(f.Bytes)) // And file is full
	//Write bytes to file as new line
	cell4 := "four"
	_, err = b.WriteLine([]byte(cell4))
	require.NoError(t, err)
	assert.Empty(t, b.String())                                                  //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep+cell3+cell3+cell4+"\n", string(f.Bytes)) // And file is full
	//Write string to file as new line
	_, err = b.WriteLineString(cell4)
	require.NoError(t, err)
	assert.Empty(t, b.String())                                                             //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep+cell3+cell3+cell4+"\n"+cell4+"\n", string(f.Bytes)) // And file is full

	cell5 := "four"
	err = b.AddCell(cell5)
	require.NoError(t, err)
	_, err = b.WriteInto(f)
	require.NoError(t, err)
	assert.Empty(t, b.String())                                                                       //Now buffer is empty
	assert.Equal(t, cell1+sep+cell2+sep+cell3+cell3+cell4+"\n"+cell4+"\n"+cell5+sep, string(f.Bytes)) // And file is full
	/*

		WriteInto*/
}
