package csvw

import (
	"fmt"

	"github.com/lazybark/go-helpers/conv"
)

// CompareCSVs takes one as base csv dataset and two as changed dataset. Then compares
// column by column (compareCols) using keyCol as line primary ID.
//
// Generates a Compared struct that can write results if needed
func CompareCSVs(one, two, dividerOne, dividerTwo, keyCol string, compareCols ...string) (c Compared, err error) {
	var totalOne, totalTwo, diff, same, del int
	var diffFields = make(map[string]int)
	mapOne, _, err := conv.ConvertCSVFiletoMap(one, dividerOne, compareCols...)
	if err != nil {
		err = fmt.Errorf("[CompareCSVs] %s: %w", two, err)
		return
	}
	totalOne = len(mapOne)

	mapTwo, _, err := conv.ConvertCSVFiletoMap(two, dividerTwo, compareCols...)
	if err != nil {
		err = fmt.Errorf("[CompareCSVs] %s: %w", two, err)
		return
	}
	totalTwo = len(mapTwo)

	comp := make(map[string]map[string]string)
	var deleted []map[string]string
	var different []Different

	for _, c := range mapTwo {
		comp[c[keyCol]] = make(map[string]string)
		comp[c[keyCol]] = c
	}

	var d bool
	var dCols map[string]string
	for _, c := range mapOne {
		if _, ok := comp[c[keyCol]]; !ok {
			del++
			deleted = append(deleted, c)
		} else {
			d = false
			dCols = make(map[string]string)
			for _, col := range compareCols {
				//There will be no fields that don't exist - ConvertCSVFiletoMap will make sure of that
				if c[col] != comp[c[keyCol]][col] {
					d = true
					diffFields[col]++
					dCols[col] = "TRUE"
				}
			}
			if d {
				diff++
				different = append(different, Different{Cols: dCols, RowOne: c, RowTwo: comp[c[keyCol]]})
			} else {
				same++
			}
		}
	}

	c = Compared{
		one:         one,
		two:         two,
		Divider:     dividerOne,
		keyCol:      keyCol,
		compareCols: compareCols,
		different:   different,
		deleted:     deleted,
		totalOne:    totalOne,
		totalTwo:    totalTwo,
		diff:        diff,
		same:        same,
		del:         del,
		diffFields:  diffFields,
	}

	return
}
