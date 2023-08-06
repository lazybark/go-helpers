package sec

import (
	"testing"
)

var generateRandomStringBenchmarkResult string

func BenchmarkGenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomStringBenchmarkResult, _ = GenerateRandomString(50)
	}
}
