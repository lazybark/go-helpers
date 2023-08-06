package conv

import (
	"testing"

	"github.com/lazybark/go-helpers/gen"
	"github.com/lazybark/go-helpers/mock"
)

var csvColsBench = []string{"id", "name", "email", "some1", "some2", "some3", "some4", "some5", "some6", "some7"}

func BenchmarkCSVToMap(b *testing.B) {
	file := &mock.MockWriteReader{}

	//Appending makes content more readable
	//Head / cols
	for _, v := range csvColsBench {
		file.Bytes = append(file.Bytes, []byte(v+csvDivider)...)
	}
	file.Bytes = append(file.Bytes, '\n')

	for i := 0; i < 1000; i++ {
		for range csvColsBench {
			file.Bytes = append(file.Bytes, []byte(gen.GenerateRandomString(50)+csvDivider)...)
		}
		file.Bytes = append(file.Bytes, '\n')
	}

	for i := 0; i < b.N; i++ {
		ConvertCSVFiletoMap(file, csvDivider, csvColsBench...)
	}

}
