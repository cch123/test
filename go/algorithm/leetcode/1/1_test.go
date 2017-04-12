package main

import "testing"

func BenchmarkSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}

func BenchmarkSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
