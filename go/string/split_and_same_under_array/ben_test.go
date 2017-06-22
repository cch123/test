package main

import "testing"

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		A()
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B()
	}
}
