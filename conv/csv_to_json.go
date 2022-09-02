package conv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lazybark/go-helpers/fsw"
)

// ConvertCSVFiletoJSON takes a file with csv formatting and marshals into json byte array
//
// To convert only specific columns or in specific order, use a slice of column names as "cols"
func ConvertCSVFiletoJSON(from string, divider string, cols ...string) ([]byte, int, error) {
	flags := os.O_RDONLY

	var array []map[string]string

	f, err := os.OpenFile(from, flags, 0666)
	if err != nil {
		return nil, 0, fmt.Errorf("[ConvertCSVFiletoJSON] can not open file: %w", err)
	}

	fileScanner := bufio.NewScanner(f)

	//We will split the file by newline
	fileScanner.Split(bufio.ScanLines)

	//First we have do determine properties of json objects in array.
	//head connects col number and col name
	head := make(map[int]string)
	headReal := make(map[int]string)
	headReverted := make(map[string]int)
	//If we have demanded head, we will just use that one
	if len(cols) > 0 {
		for n, val := range cols {
			head[n] = val
		}
	}

	var mp map[string]string
	var strArr []string
	converted := 0

	//We scan each line of file and split it by provided divider
	for fileScanner.Scan() {
		strArr = strings.Split(fileScanner.Text(), divider)
		//Now, using the first string, we need to get head of the csv table
		if converted == 0 {
			//We should cut out first magical numbers (BOM) from the text in case persists (for UTF-8)
			strArr[0] = fsw.CutBOMFromString(strArr[0])

			for n, val := range strArr {
				if len(cols) > 0 {
					//If there are specific rows only
					for nh, valh := range cols {
						if val == valh {
							headReverted[val] = nh
						}
					}
				} else {
					//If no specific rows = use all
					head[n] = val
					headReverted[val] = n
				}
				headReal[n] = val
			}
		} else {
			mp = make(map[string]string)
			for n, val := range strArr {
				//Now we need to understand if the column should be converted (is in head)
				//And then push the value into the right json object attribute
				if _, ok := headReverted[headReal[n]]; ok {
					mp[headReal[n]] = val
				}
			}
			array = append(array, mp)
		}
		converted++
	}
	f.Close()

	bytes, err := json.Marshal(array)
	if err != nil {
		return nil, 0, fmt.Errorf("[ConvertCSVFiletoJSON] can not marshal map: %w", err)
	}

	return bytes, converted, nil
}

// ConvertCSVFiletoMap takes a file with csv formatting and returns slice of map[string]string
//
// To convert only specific columns or in specific order, use a slice of column names as "cols"
func ConvertCSVFiletoMap(from string, divider string, cols ...string) ([]map[string]string, int, error) {
	flags := os.O_RDONLY

	var array []map[string]string

	f, err := os.OpenFile(from, flags, 0666)
	if err != nil {
		return nil, 0, fmt.Errorf("[ConvertCSVFiletoJSON] can not open file: %w", err)
	}

	fileScanner := bufio.NewScanner(f)

	//We will split the file by newline
	fileScanner.Split(bufio.ScanLines)

	//First we have do determine properties of json objects in array.
	//head connects col number and col name
	head := make(map[int]string)
	headReal := make(map[int]string)
	headReverted := make(map[string]int)
	//If we have demanded head, we will just use that one
	if len(cols) > 0 {
		for n, val := range cols {
			head[n] = val
		}
	}

	var mp map[string]string
	var strArr []string
	converted := 0

	//We scan each line of file and split it by provided divider
	for fileScanner.Scan() {
		strArr = strings.Split(fileScanner.Text(), divider)
		//Now, using the first string, we need to get head of the csv table
		if converted == 0 {
			//We should cut out first magical numbers (BOM) from the text in case persists (for UTF-8)
			strArr[0] = fsw.CutBOMFromString(strArr[0])

			for n, val := range strArr {
				if len(cols) > 0 {
					//If there are specific rows only
					for nh, valh := range cols {
						if val == valh {
							headReverted[val] = nh
						}
					}
				} else {
					//If no specific rows = use all
					head[n] = val
					headReverted[val] = n
				}
				headReal[n] = val
			}
		} else {
			mp = make(map[string]string)
			for n, val := range strArr {
				//Now we need to understand if the column should be converted (is in head)
				//And then push the value into the right json object attribute
				if _, ok := headReverted[headReal[n]]; ok {
					mp[headReal[n]] = val
				}
			}
			array = append(array, mp)
		}
		converted++
	}
	f.Close()

	return array, converted, nil
}
