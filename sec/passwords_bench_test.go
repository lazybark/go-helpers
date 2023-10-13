package sec

import (
	"testing"
)

var comparePasswordsBenchmarkResult bool

func BenchmarkHashPasswordAndComparePasswordWithHash(b *testing.B) {
	pwd := "some_password"
	hash, _ := HashAndSaltPasswordString(pwd, 10)

	for i := 0; i < b.N; i++ {
		comparePasswordsBenchmarkResult, _ = ComparePasswords(hash, pwd)
	}

}
