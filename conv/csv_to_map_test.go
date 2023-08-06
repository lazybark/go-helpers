package conv

import (
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
