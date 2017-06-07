package main

import "testing"

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestA()
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestB()
	}
}
