package main

import "testing"

func BenchmarkAtomic(b * testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				atomicStore(0)
			}
		})
}
