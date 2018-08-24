package main

import "testing"

func BenchCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc()
	}
}
