package main

import "testing"

func BenchmarkTestWithLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useAppendWithLength()
	}
}

func BenchmarkTestWithoutLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useAppendWithoutLength()
	}
}

func BenchmarkTestNotUseAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		notUseAppend()
	}
}
