// fsw is a package with simple methods to work with filesystems.
package fsw

import (
	"fmt"
	"os"
	"path/filepath"
)

// MakePathToFile creates full path to file in filesystem, creates the file and truncates
// in case truncate = true.
//
// Useful in cases where there is no need to check content (logs, for example).
func MakePathToFile(path string, truncate bool) (*os.File, error) {
	dir, _ := filepath.Split(path)
	if dir == "" {
		dir = "/"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("[MakePathToFile] can not make dir: %w", err)
	}

	flags := os.O_CREATE | os.O_APPEND
	if truncate {
		flags += os.O_TRUNC
	}

	f, err := os.OpenFile(path, flags, 0666)
	if err != nil {
		return nil, fmt.Errorf("[MakePathToFile] can not open file: %w", err)
	}

	return f, nil
}

// CutBOMFromString returns string with byte order mark removed in case it exists at the start
// of the string.
//
// Useful for reading UTF-8 files
func CutBOMFromString(str string) string {
	bom := []byte(str)
	if bom[0] == 239 && bom[1] == 187 && bom[2] == 191 {
		return string(bom[3:])
	}

	return str
}
