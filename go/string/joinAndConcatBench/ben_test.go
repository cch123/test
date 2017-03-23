package main

import "testing"

func BenchmarkSprintfAndStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfAndStringConcat()
	}
}

func BenchmarkSprintfAndSliceAppendAndStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfAndSliceAppendAndStringJoin()
	}
}

func BenchmarkSprintfAndPreallocatedArrayAndJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfAndPreallocatedArrayAndJoin()
	}

}

func BenchmarkSprintfAndPreallocatedSliceAndJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfAndPreallocatedSliceAndJoin()
	}
}

func BenchmarkNoSprintfAndSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noSprintfAndSliceAppend()
	}
}

func BenchmarkNoSprintfAndPreallocatedSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noSprintfAndPreallocatedSlice()
	}
}

func BenchmarkBytesBufferAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesBufferAppend()
	}
}
