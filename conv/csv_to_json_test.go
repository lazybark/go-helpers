package conv

import (
	"encoding/json"
	"testing"

	"github.com/lazybark/go-helpers/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSVToJSON(t *testing.T) {
	file := &mock.MockFile{}

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

	map1, n1, err := ConvertCSVFiletoMap(file, csvDivider, csvCols...)
	require.NoError(t, err)

	bytes, err := json.Marshal(map1)
	require.NoError(t, err)

	file.SetLastRead(0) // Reset reading marker
	jsonned, n2, err := ConvertCSVFiletoJSON(file, csvDivider, csvCols...)
	require.NoError(t, err)

	assert.Equal(t, n1, n2)

	assert.Equal(t, string(bytes), string(jsonned))
}
