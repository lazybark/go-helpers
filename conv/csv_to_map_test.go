package conv

import (
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCSVString struct {
	keyColValue  string
	stringValues []string
}

const (
	csvDivider = ";"
)

var (
	csvCols  = []string{"id", "name", "email"}
	csvLines = []testCSVString{
		{
			keyColValue:  "1",
			stringValues: []string{"1", "user 1", "user1@gmail.com"},
		},
		{
			keyColValue:  "2",
			stringValues: []string{"2", "user 2", "user2@gmail.com"},
		},
		{
			keyColValue:  "3",
			stringValues: []string{"3", "user 3", "user3@gmail.com"},
		},
		{
			keyColValue:  "yy",
			stringValues: []string{"yy", "dfsdfsdf", "uyjhgrtyhj"},
		},
	}
)

func TestCSVToMap(t *testing.T) {
	file := &mock.MockWriteReader{}

	//Appending makes content more readable
	//Head / cols
	for _, v := range csvCols {
		file.Bytes = append(file.Bytes, []byte(v+csvDivider)...)
	}
	file.Bytes = append(file.Bytes, '\n')

	for _, v := range csvLines {
		for _, svv := range v.stringValues {
			file.Bytes = append(file.Bytes, []byte(svv+csvDivider)...)
		}
		file.Bytes = append(file.Bytes, '\n')
	}

	mapOne, _, err := ConvertCSVFiletoMap(file, csvDivider, csvCols...)
	require.NoError(t, err)

	assert.Equal(t, len(csvLines), len(mapOne))
	for k, v := range csvLines {
		for svk, svv := range v.stringValues {
			assert.Equal(t, svv, mapOne[k][csvCols[svk]])
		}
	}
}
