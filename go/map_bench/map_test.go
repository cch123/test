package main
import "testing"

func BenchmarkNormalMapInsert(b * testing.B) {
	for i:=0;i<b.N;i++{
		normalMapInsert(i, i+10)
	}
}

func BenchmarkSyncMapInsert(b * testing.B) {
	for i:=0; i<b.N; i++{
		syncMapInsert(i, i+10)
	}
}

func BenchmarkNormalMapGet(b * testing.B) {
	for i:=0;i<b.N;i++{
		normalMapGet(i)
	}
}

func BenchmarkSyncMapGet(b * testing.B) {
	for i:=0;i<b.N;i++{
		syncMapGet(i)
	}
}