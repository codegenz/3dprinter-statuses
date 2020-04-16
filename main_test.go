package main

import (
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifDouble(10000, 1, 10000, 1)
	}
}
