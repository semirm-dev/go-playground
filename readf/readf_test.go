package readf_test

import (
	"github.com/semirm-dev/go-playground/readf"
	"testing"
)

func BenchmarkReadAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readf.ReadAll()
	}
}

func BenchmarkReadChunk(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readf.ReadChunk()
	}
}

func BenchmarkReadBuff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readf.ReadBuff()
	}
}

func BenchmarkReadBuffR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readf.ReadBuffR()
	}
}

func BenchmarkPipeRW(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readf.PipeRW()
	}
}
