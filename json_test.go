package main

import "testing"

func TestHalderEasyJson(t *testing.T) {
	HalderEasyJson(1000)
	t.Skip()
}

func TestHalderEncodeJson(t *testing.T) {
	HalderEncodeJson(1000)
	t.Skip()
}

func TestHalderJsonIter(t *testing.T) {
	HalderJsonIter(1000)
	t.Skip()
}

func TestHalderFfjson(t *testing.T) {
	HalderFfjson(1000)
	t.Skip()
}

func BenchmarkEasyJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HalderEasyJson(1000)
	}
}

func BenchmarkEncodeJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HalderEncodeJson(1000)
	}
}

func BenchmarkJsonIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HalderJsonIter(10)
	}
}

func BenchmarkFfjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HalderFfjson(1)
	}
}
