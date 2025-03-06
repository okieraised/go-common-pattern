package main

import "testing"

func BenchmarkMultiply(b *testing.B) {
	x := 42
	for i := 0; i < b.N; i++ {
		_ = x * 16
	}
}

func BenchmarkShift(b *testing.B) {
	x := 42
	for i := 0; i < b.N; i++ {
		_ = x << 4
	}
}
