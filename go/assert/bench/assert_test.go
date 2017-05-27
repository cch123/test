package main

import "testing"

func BenchmarkAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asx()
	}
}
