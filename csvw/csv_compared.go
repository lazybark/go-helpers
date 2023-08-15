package csvw

import (
	"fmt"
	"path/filepath"

	"github.com/lazybark/go-helpers/fsw"
)

// Compared holds data of two compared csv files, including statistic. Methods WriteDeleted &
// WriteDifferent will create csv files with resulting difference dataset.
type Compared struct {
	one         string              //Path to first file
	two         string              //Path to second file
	Divider     string              //Divider can be manually changed if you need to see results with different one
	keyCol      string              //Column to use as key
	compareCols []string            //Columns to compare (should include keyCol!)
	different   []Different         //All different lines
	deleted     []map[string]string //All deleted lines
	totalOne    int
	totalTwo    int
	diff        int
	same        int
	del         int
	diffFields  map[string]int
}

// Different represents two different rows and list of fields that differ
type Different struct {
	Cols   map[string]string
	RowOne map[string]string
	RowTwo map[string]string
}

// WriteDeleted writes into file full list of rows that differ from first to second file
func (c *Compared) WriteDifferent(file fsw.IFileWriter) error {
	diffB := NewCSVBuilder(c.Divider)
	diffB.UseFile(file)
	defer diffB.Close()

	err := diffB.AddCell("doc")
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
	}

	for _, cc := range c.compareCols {
		err = diffB.AddCell(cc)
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}

		if _, exist := c.diffFields[cc]; exist {
			err = diffB.AddCell(cc + "_d")
			if err != nil {
				return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
			}
		}
	}
	err = diffB.NewLine()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
	}

	_, err = diffB.WriteBuffer()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
	}

	for _, d := range c.different {
		err = diffB.AddCell("//")
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}
		err = diffB.NewLine()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}

		err = diffB.AddCell(filepath.Base(c.one))
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}
		for _, cc := range c.compareCols {
			err = diffB.AddCell(d.RowOne[cc])
			if err != nil {
				return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
			}
			if _, exist := c.diffFields[cc]; exist {
				if dc, ok := d.Cols[cc]; ok {
					err = diffB.AddCell(dc)
					if err != nil {
						return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
					}
				} else {
					err = diffB.AddCell("")
					if err != nil {
						return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
					}
				}
			}
		}
		err = diffB.NewLine()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}

		err = diffB.AddCell(filepath.Base(c.two))
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}
		for _, cc := range c.compareCols {
			err = diffB.AddCell(d.RowTwo[cc])
			if err != nil {
				return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
			}
			if _, exist := c.diffFields[cc]; exist {
				if dc, ok := d.Cols[cc]; ok {
					err = diffB.AddCell(dc)
					if err != nil {
						return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
					}
				} else {
					err = diffB.AddCell("")
					if err != nil {
						return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
					}
				}
			}

		}
		err = diffB.NewLine()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}

		_, err = diffB.WriteBuffer()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}
	}

	return nil
}

// WriteDeleted writes into file full list of deleted rows (that exist in first file, but not in second)
func (c *Compared) WriteDeleted(file fsw.IFileWriter) error {
	delB := NewCSVBuilder(c.Divider)
	delB.UseFile(file)
	defer delB.Close()

	err := delB.AddCell(c.keyCol)
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
	}

	for _, c := range c.compareCols {
		err = delB.AddCell(c)
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
		}
	}
	err = delB.NewLine()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
	}

	_, err = delB.WriteBuffer()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
	}

	for _, l := range c.deleted {
		err = delB.AddCell(l[c.keyCol])
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
		}

		for _, c := range c.compareCols {
			err = delB.AddCell(l[c])
			if err != nil {
				return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
			}
		}

		err = delB.NewLine()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
		}

		_, err = delB.WriteBuffer()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
		}
	}

	return nil
}

// TotalRowsInSecondFile returns list of rows that differ (compared by ID)
func (c *Compared) DifferentRows() []Different {
	return c.different
}

// TotalRowsInSecondFile returns list of deleted rows (exist in first file, but not in second)
func (c *Compared) DeletedRows() []map[string]string {
	return c.deleted
}

// TotalRowsInSecondFile returns number of rows in first file
func (c *Compared) TotalRowsInFirstFile() int {
	return c.totalOne
}

// TotalRowsInSecondFile returns number of rows in second file
func (c *Compared) TotalRowsInSecondFile() int {
	return c.totalTwo
}

// DifferentRowsCount returns number of rows that differ from document to document, but have same ID
func (c *Compared) DifferentRowsCount() int {
	return c.diff
}

// SameRowsCount returns number of rows that are same in both documents
func (c *Compared) SameRowsCount() int {
	return c.same
}

// DeletedRowsCount returns number of rows that exist in first document, but not in second
func (c *Compared) DeletedRowsCount() int {
	return c.del
}

// DifferentFieldsStat returns list of column names with number of how many rows have different value in each column
func (c *Compared) DifferentFieldsStat() map[string]int {
	return c.diffFields
}
