package hasher

import (
	"testing"

	"github.com/lazybark/go-helpers/gen"
	"github.com/lazybark/go-helpers/mock"
)

func BenchmarkFileHashFile(b *testing.B) {
	file := mock.MockWriteReader{
		Bytes: []byte(gen.GenerateRandomBytes(50)),
	}

	for i := 0; i < b.N; i++ {
		HashFile(&file, MD5, 50)
	}

}
