package main

import (
	"testing"
)

func TestAwesomeToJSON(t *testing.T) {
	EncodeJson()
	t.Skip()

}

func BenchmarkAwesomeToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeJson()
	}
}

func BenchmarkAwesomeToJSONPretty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeJson()
	}
}