package csvw

import (
	"fmt"
	"testing"

	"github.com/lazybark/go-helpers/gen"
	"github.com/lazybark/go-helpers/mock"
)

var (
	benchmarkBufferAddCellResult         error
	benchmarkBufferAddLineResult         error
	benchmarkBufferWriteBufferResult     error
	benchmarkBufferWriteBytesResult      error
	benchmarkBufferWriteStringResult     error
	benchmarkBufferWriteLineStringResult error
	benchmarkBufferWriteIntoResult       error
)

func BenchmarkBufferAddCell(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	cell1 := gen.GenerateRandomString(50)
	cell2 := gen.GenerateRandomString(50)
	cell3 := gen.GenerateRandomString(50)

	for i := 0; i < b.N; i++ {
		benchmarkBufferAddCellResult = bldr.AddCell(cell1, cell2, cell3)
	}
}

func BenchmarkBufferAddLine(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	line := gen.GenerateRandomString(150)

	for i := 0; i < b.N; i++ {
		benchmarkBufferAddLineResult = bldr.AddLine(line)
	}
}

func BenchmarkBufferWriteBuffer(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	line := gen.GenerateRandomString(150)
	err := bldr.AddLine(line)
	if err != nil {
		b.Error(fmt.Errorf("[BenchmarkBufferWriteBuffer] %w", err))
	}

	for i := 0; i < b.N; i++ {
		_, benchmarkBufferWriteBufferResult = bldr.WriteBuffer()
	}
}

func BenchmarkBufferWriteBytes(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	line := gen.GenerateRandomBytes(150)

	for i := 0; i < b.N; i++ {
		_, benchmarkBufferWriteBytesResult = bldr.Write(line)
	}
}

func BenchmarkBufferWriteString(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	line := gen.GenerateRandomString(150)

	for i := 0; i < b.N; i++ {
		_, benchmarkBufferWriteStringResult = bldr.WriteString(line)
	}
}

func BenchmarkBufferWriteLineString(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	bldr.UseFile(f)

	line := gen.GenerateRandomString(150)

	for i := 0; i < b.N; i++ {
		_, benchmarkBufferWriteLineStringResult = bldr.WriteLineString(line)
	}
}

func BenchmarkBufferWriteInto(b *testing.B) {
	f := &mock.MockFile{}

	sep := ";"
	bldr := NewCSVBuilder(sep)

	line := gen.GenerateRandomString(150)
	err := bldr.AddCell(line)
	if err != nil {
		b.Error(fmt.Errorf("[BenchmarkBufferWriteInto] %w", err))
	}

	for i := 0; i < b.N; i++ {
		_, benchmarkBufferWriteIntoResult = bldr.WriteInto(f)
	}
}
