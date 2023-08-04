package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lazybark/go-helpers/cli"
	"github.com/lazybark/go-helpers/cli/clf"
	"github.com/lazybark/go-helpers/csvw"
	"github.com/lazybark/go-helpers/fsw"
)

func main() {
	cfg, err := ParseEnv()
	if err != nil {
		fmt.Println(clf.Red("Error: ", err))
		return
	}

	fmt.Println(clf.Green("Welcome to CSV comparing tool\n"))
	if cfg.First == "" {
		fmt.Printf("Provide path to first %s file. It will be used as initial set. 1.csv if empty\n", clf.Red("(base)"))
		cfg.First = cli.AwaitCLIcommand()
		if cfg.First == "" {
			cfg.First = "1.csv"
		}
	}
	if cfg.SepOne == "" {
		fmt.Printf("Separator symbol for the 1st file %s. ';' if empty\n", clf.Red("(divides cells, e.g. ',' or ';')"))
		cfg.SepOne = cli.AwaitCLIcommand()
		if cfg.SepOne == "" {
			cfg.SepOne = ";"
		}
	}
	if cfg.Second == "" {
		fmt.Println("Provide path to second file. It will be used as new set that needs to be compared. 2.csv if empty")
		cfg.Second = cli.AwaitCLIcommand()
		if cfg.Second == "" {
			cfg.Second = "2.csv"
		}
	}
	if cfg.SepTwo == "" {
		fmt.Printf("Separator symbol for the 2nd file %s ';' if empty\n", clf.Red("(divides cells, e.g. ',' or ';')"))
		cfg.SepTwo = cli.AwaitCLIcommand()
		if cfg.SepTwo == "" {
			cfg.SepTwo = ";"
		}
	}
	if cfg.KeyCol == "" {
		fmt.Println("Enter name of the key column that will be used to identify rows between files.")
		fmt.Println("Wrong choice of not existing column will result in wrong output")
		cfg.KeyCol = cli.AwaitCLIcommand()
	}
	if cfg.ColsString == "" {
		fmt.Println("Enter list of columns that need to be compared, divided by comma without space (", clf.Red("Including key column)"))
		fmt.Println("All columns should exist. Otherwise you can get unexpected results")
		cfg.ColsString = cli.AwaitCLIcommand()
	}

	fOne, err := os.OpenFile(cfg.First, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(clf.Red("Error: can not open file: %w", err))
		return
	}
	fTwo, err := os.OpenFile(cfg.Second, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(clf.Red("Error: can not open file: %w", err))
		return
	}

	c, err := csvw.CompareCSVs(fOne, fTwo, cfg.First, cfg.Second, cfg.SepOne, cfg.SepTwo, cfg.KeyCol, strings.Split(cfg.ColsString, ",")...)
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

	if cfg.WriteDiffs {
		if cfg.DiffPath == "" {
			cfg.DiffPath = fmt.Sprintf("%s_different_rows_%d.csv", cfg.First, time.Now().Unix())
		}
		df, err := fsw.MakePathToFile(cfg.DiffPath, true)
		if err != nil {
			fmt.Println(clf.Red("Error: ", err))
			return
		}

		err = c.WriteDifferent(df)
		if err != nil {
			fmt.Println(clf.Red("Error: ", err))
			return
		}
	} else {
		fmt.Println("Write differences to new CSV file? y/n")
		wd := cli.AwaitCLIcommand()
		if wd == "y" || wd == "Y" {
			fmt.Println("Path to new file. Leave empty to create new in the same directory. If file exists, it will be truncated")
			p := cli.AwaitCLIcommand()

			if p == "" {
				p = fmt.Sprintf("%s_different_rows_%d.csv", cfg.First, time.Now().Unix())
			}
			df, err := fsw.MakePathToFile(p, true)
			if err != nil {
				fmt.Println(clf.Red("Error: ", err))
				return
			}

			err = c.WriteDifferent(df)
			if err != nil {
				fmt.Println(clf.Red("Error: ", err))
				return
			}
			fmt.Println(clf.Green("File created"))
		}
	}

	if cfg.WriteDeleted {
		if cfg.DiffPath == "" {
			cfg.DiffPath = fmt.Sprintf("%s_deleted_rows_%d.csv", cfg.First, time.Now().Unix())
		}
		df, err := fsw.MakePathToFile(cfg.DiffPath, true)
		if err != nil {
			fmt.Println(clf.Red("Error: ", err))
			return
		}

		err = c.WriteDeleted(df)
		if err != nil {
			fmt.Println(clf.Red("Error: ", err))
			return
		}
	} else {
		fmt.Println("Write deleted rows to new CSV file? y/n")
		dd := cli.AwaitCLIcommand()
		if dd == "y" || dd == "Y" {
			fmt.Println("Path to new file. Leave empty to create new in the same directory. If file exists, it will be truncated")
			p := cli.AwaitCLIcommand()

			if p == "" {
				p = fmt.Sprintf("%s_different_rows_%d.csv", cfg.First, time.Now().Unix())
			}
			df, err := fsw.MakePathToFile(p, true)
			if err != nil {
				fmt.Println(clf.Red("Error: ", err))
				return
			}

			err = c.WriteDeleted(df)
			if err != nil {
				fmt.Println(clf.Red("Error: ", err))
				return
			}
			fmt.Println(clf.Green("File created"))
		}
	}

	fmt.Println(clf.Green("===DONE==="))
}
