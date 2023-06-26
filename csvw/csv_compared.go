package csvw

import (
	"fmt"
	"time"
)

// Compared holds data of two compared csv files, including statistic. Methods WriteDeleted &
// WriteDifferent will create csv files with resulting difference dataset.
type Compared struct {
	one         string              //Path to first file
	two         string              //Path to second file
	Divider     string              //Divider can be manually changed if you need to see results with different one
	keyCol      string              //Column to use as key
	compareCols []string            //Columns to compare (should include keyCol!)
	different   []map[string]string //All different lines
	deleted     []map[string]string //All deleted lines
	totalOne    int
	totalTwo    int
	diff        int
	same        int
	del         int
}

// WriteDeleted creates new .csv file with full list of rows that differ from first to second file
func (c *Compared) WriteDifferent(path string) error {
	if path == "" {
		path = fmt.Sprintf("%s_different_rows_%d.csv", c.one, time.Now().Unix())
	}

	diffB := NewCSVBuilder(c.Divider)
	err := diffB.OpenFile(path, true)
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
	}
	defer diffB.Close()

	diffB.AddCell(c.keyCol)
	for _, c := range c.compareCols {
		diffB.AddCell(c)
	}
	diffB.NewLine()
	_, err = diffB.WriteBuffer()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
	}

	for _, l := range c.different {
		for _, c := range c.compareCols {
			diffB.AddCell(l[c])
		}
		diffB.NewLine()
		_, err = diffB.WriteBuffer()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDifferences]: %w", err)
		}
	}

	return nil
}

// WriteDeleted creates new .csv file with full list of deleted rows (exist in first file, but not in second)
func (c *Compared) WriteDeleted(path string) error {
	if path == "" {
		path = fmt.Sprintf("%s_deleted_rows_%d.csv", c.one, time.Now().Unix())
	}

	delB := NewCSVBuilder(c.Divider)
	err := delB.OpenFile(path, true)
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
	}
	defer delB.Close()

	delB.AddCell(c.keyCol)
	for _, c := range c.compareCols {
		delB.AddCell(c)
	}
	delB.NewLine()
	_, err = delB.WriteBuffer()
	if err != nil {
		return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
	}

	for _, l := range c.deleted {
		for _, c := range c.compareCols {
			delB.AddCell(l[c])
		}
		delB.NewLine()
		_, err = delB.WriteBuffer()
		if err != nil {
			return fmt.Errorf("[CompareCSVs][WriteDeleted]: %w", err)
		}
	}

	return nil
}

// TotalRowsInSecondFile returns list of rows that differ (compared by ID)
func (c *Compared) DifferentRows() []map[string]string {
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

// SameRowsCount returns number of rows that exist in both documents
func (c *Compared) SameRowsCount() int {
	return c.same
}

// DeletedRowsCount returns number of rows that exist in first document, but to in second
func (c *Compared) DeletedRowsCount() int {
	return c.del
}
