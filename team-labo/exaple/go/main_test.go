package main

import (
	"testing"
)

func BenchmarkMult7Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultipleOf7Add()
	}
}

func BenchmarkMult7OneByOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultipleOf7OneByOne()
	}
}
