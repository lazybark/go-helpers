package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lazybark/go-helpers/cli/clf"
	"github.com/lazybark/go-helpers/csvw"
	"github.com/lazybark/go-helpers/fsw"
)

const ()

var (
	csvCols1       = []string{"user_id", "name", "email", "city", "some", "another"}
	csvColsCompare = []string{"user_id", "name", "email", "city"}
)

func main() {
	csvDivider := ";"

	fOneStr := "example-1.csv"
	fOne, err := os.OpenFile(fOneStr, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(clf.Red("Error: can not open file: %w", err))
		return
	}
	fTwoStr := "example-2.csv"
	fTwo, err := os.OpenFile(fTwoStr, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(clf.Red("Error: can not open file: %w", err))
		return
	}

	c, err := csvw.CompareCSVs(fOne, fTwo, fOneStr, fTwoStr, csvDivider, csvDivider, csvCols1[0], csvColsCompare...)
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}
	fmt.Printf("Found %d & %d strings in files 1 and 2\n", c.TotalRowsInFirstFile(), c.TotalRowsInSecondFile())
	fmt.Printf("Different lines: %d\n", c.DifferentRowsCount())
	fmt.Printf("Same lines: %d\n", c.SameRowsCount())
	fmt.Printf("Lines that not exist in second file: %d\n", c.DeletedRowsCount())

	df := c.DifferentFieldsStat()
	if len(df) > 0 {
		fmt.Println("\nBy fields:")
		for n, i := range df {
			fmt.Printf("%s: %d\n", n, i)
		}
	}

	diff, err := fsw.MakePathToFile(fmt.Sprintf("%s_different_rows_%d.csv", fOneStr, time.Now().Unix()), true)
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}

	err = c.WriteDifferent(diff)
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}

	del, err := fsw.MakePathToFile(fmt.Sprintf("%s_deleted_rows_%d.csv", fTwoStr, time.Now().Unix()), true)
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}

	err = c.WriteDeleted(del)
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}

	/*file1 := &mock.MockWriteReader{}
	file2 := &mock.MockWriteReader{}

	//Head / cols
	for _, v := range csvCols1 {
		file1.Bytes = append(file1.Bytes, []byte(v+csvDivider)...)
	}
	file1.Bytes = append(file1.Bytes, '\n')

	for _, v := range csvCols2 {
		file2.Bytes = append(file2.Bytes, []byte(v+csvDivider)...)
	}
	file2.Bytes = append(file2.Bytes, '\n')

	for _, v := range csvLines1 {
		for _, svv := range v.stringValues {
			file1.Bytes = append(file1.Bytes, []byte(svv+csvDivider)...)
		}
		file1.Bytes = append(file1.Bytes, '\n')
	}
	for _, v := range csvLines2 {
		for _, svv := range v.stringValues {
			file2.Bytes = append(file2.Bytes, []byte(svv+csvDivider)...)
		}
		file2.Bytes = append(file2.Bytes, '\n')
	}

	name1 := "file1"
	name2 := "file2"
	c, err := csvw.CompareCSVs(file1, file2, name1, name2, csvDivider, csvDivider, csvCols1[0], csvColsCompare...)
	if err != nil {
		log.Fatal(err)
	}*/
}
