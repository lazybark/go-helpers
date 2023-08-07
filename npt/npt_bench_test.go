package npt

import (
	"testing"
	"time"
)

var (
	nptNowBenchmarkResult NPT
)

func BenchmarkNPTNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nptNowBenchmarkResult = Now()
	}
}

func BenchmarkNPTFromTime(b *testing.B) {
	tm := time.Now()
	nt := Now()

	for i := 0; i < b.N; i++ {
		nt.FromTime(tm)
	}
}

func BenchmarkNPTAdd(b *testing.B) {
	tm := time.Now()

	nt := Now()
	nt.FromTime(tm)

	for i := 0; i < b.N; i++ {
		nt.Add(time.Minute)
	}
}
