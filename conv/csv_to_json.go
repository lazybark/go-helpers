package conv

import (
	"encoding/json"
	"fmt"

	"github.com/lazybark/go-helpers/fsw"
)

// ConvertCSVFiletoJSON takes a file with csv formatting and marshals into json byte array.
//
// To convert only specific columns or in specific order, use a slice of column names as "cols"
func ConvertCSVFiletoJSON(f fsw.IFileReader, divider string, cols ...string) ([]byte, int, error) {
	array, converted, err := ConvertCSVFiletoMap(f, divider, cols...)
	if err != nil {
		return nil, 0, fmt.Errorf("[ConvertCSVFiletoJSON]%w", err)
	}

	fmt.Println("ss")

	bytes, err := json.Marshal(array)
	if err != nil {
		return nil, 0, fmt.Errorf("[ConvertCSVFiletoJSON] can not marshal map: %w", err)
	}

	return bytes, converted, nil
}
