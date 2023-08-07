package csvw

import (
	"testing"

	"github.com/lazybark/go-helpers/gen"
	"github.com/lazybark/go-helpers/mock"
)

var benchmarkCompareCSVsResult Compared

func BenchmarkCompareCSVs(b *testing.B) {
	file1 := &mock.MockWriteReader{}
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

	for i := 0; i < 1000; i++ {
		for range csvCols1 {
			file1.Bytes = append(file1.Bytes, []byte(gen.GenerateRandomString(50)+csvDivider)...)
		}
		file1.Bytes = append(file1.Bytes, '\n')
	}
	for i := 0; i < 1000; i++ {
		for range csvCols1 {
			file2.Bytes = append(file2.Bytes, []byte(gen.GenerateRandomString(50)+csvDivider)...)
		}
		file2.Bytes = append(file2.Bytes, '\n')
	}

	name1 := "file1"
	name2 := "file2"

	for i := 0; i < b.N; i++ {
		benchmarkCompareCSVsResult, _ = CompareCSVs(file1, file2, name1, name2, csvDivider, csvDivider, csvCols1[0], csvColsCompare...)
	}

}
