package csvw

import (
	"fmt"
	"os"
	"strings"

	"github.com/lazybark/go-helpers/fsw"
)

// CSVBuilder uses strings.Builder to create CSV lines/cells one by one or in a batch.
// Capable of writing right into file specified.
type CSVBuilder struct {
	Builder   *strings.Builder
	Separator string
	f         *os.File
}

// NewCSVBuilder returns new CSVBuilder with no file attached
func NewCSVBuilder(sep string) CSVBuilder {
	return CSVBuilder{Builder: &strings.Builder{}, Separator: sep}
}

// Close closes the .csv file
func (b *CSVBuilder) Close() error {
	return b.f.Close()
}

// OpenFile open a file to write csv data into
func (b *CSVBuilder) OpenFile(p string, truncate bool) (err error) {
	b.f, err = fsw.MakePathToFile(p, truncate)
	return
}

// AddCell adds new cell to current string (with separator at the end)
func (b *CSVBuilder) AddCell(str ...string) {
	for _, s := range str {
		b.Builder.WriteString(s + b.Separator)
	}
}

// NewLine adds line break to current string
func (b *CSVBuilder) NewLine() {
	b.Builder.WriteString("\n")
}

// Reset cleans current string
func (b *CSVBuilder) Reset() {
	b.Builder.Reset()
}

// String returns current string data
func (b *CSVBuilder) String() string {
	return b.Builder.String()
}

// WriteBuffer writes current byte buffer into opened file and cleans the buffer right after.
func (b *CSVBuilder) WriteBuffer() (int, error) {
	n, err := b.f.Write([]byte(b.Builder.String()))
	if err != nil {
		return n, fmt.Errorf("[CSVBuilder][WriteBuffer]: %w", err)
	}
	b.Reset() //Always reset buffer before next write
	return n, nil

}

// Write writes bytes directly into file (no line break at the end)
func (b *CSVBuilder) Write(bts []byte) (int, error) {
	return b.f.Write(bts)
}

// WriteString writes s directly into file (no line break at the end)
func (b *CSVBuilder) WriteString(s string) (int, error) {
	bts := []byte(s)
	return b.f.Write(bts)
}

// WriteLine writes bytes directly into file and adds line break after last byte
func (b *CSVBuilder) WriteLine(bts []byte) (int, error) {
	bts = append(bts, '\n')
	return b.f.Write(bts)
}

// WriteLineString writes s directly into file and adds line break after last byte
func (b *CSVBuilder) WriteLineString(s string) (int, error) {
	bts := []byte(s)
	bts = append(bts, '\n')
	return b.f.Write(bts)
}
