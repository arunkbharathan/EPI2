package main

import (
	"testing"
)

func BenchmarkPow1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow1(1.21, 1000)
	}
}
func BenchmarkPow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow2(1.21, 1000)
	}
}
func BenchmarkPow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow3(1.21, 1000)
	}
}

func BenchmarkPow4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow4(1.21, 1000)
	}
}
