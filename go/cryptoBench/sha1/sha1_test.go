package main

import "testing"

func BenchmarkSHA1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		test()
	}
}
