package sys_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/semirm-dev/go-playground/sys"
)

// BenchmarkReadAll-32    1280    916788 ns/op     5718.75 MB/s   11042754 B/op   31 allocs/op
// BenchmarkReadChunk-32  4905    237514 ns/op    22074.01 MB/s        168 B/op    3 allocs/op
// BenchmarkReadBuff-32   1033   1118089 ns/op     4689.14 MB/s   16744641 B/op   12 allocs/op
// BenchmarkReadBufio-32  1125   1075612 ns/op     4874.32 MB/s       4264 B/op    4 allocs/op

// # Go I/O Benchmark Analysis

// ### 1. `ReadChunk` is the Clear Winner (4x faster, 99.998% less memory)
// * **22,074 MB/s (22 GB/s)** throughput vs **4.6–5.7 GB/s** for full-buffer reads.
// * **168 B/op vs 11–16 MB/op:** Only **3 allocations** total (the `os.File` descriptor metadata itself).
// * Zero heap allocations during the read loop means zero garbage collector pressure. The 32 KB buffer stays on the stack/registers, cycling in L1/L2 cache.

// ---

// ### 2. The Dynamic Buffer Doubling Penalty (`ReadAll` vs `ReadBuff`)
// * **`BenchmarkReadAll`** allocated **~11 MB** for a 5 MB payload (`11,042,754 B/op`, **31 allocations**).
//   * Go starts with a small slice (512 B) and repeatedly doubles it ($512 \rightarrow 1024 \rightarrow 2048 \dots \rightarrow 8\text{ MB} \rightarrow 16\text{ MB}$). Every doubling creates garbage for the GC to collect.
// * **`BenchmarkReadBuff`** allocated **~16.7 MB** (`16,744,641 B/op`).
//   * Because `ReadBuff` created a new `bytes.Buffer` inside the function without pre-allocation or reuse, `Buffer.ReadFrom` grew its internal backing slice dynamically, triggering even larger geometric slice growth steps.

// ---

// ### 3. `ReadBufio` (Zero-Alloc Line Parsing)
// * **4,264 B/op with only 4 allocs:**
//   * The entire 5 MB file was parsed line-by-line using only the initial internal buffer of `bufio.Reader` (4096 bytes + descriptor structs).
//   * Not a single byte slice was allocated on the heap for any of the individual lines.

const benchFileSize = 5 * 1024 * 1024 // 5 MB payload

// Global sinks to prevent compiler dead-code elimination
var (
	sinkBytes []byte
)

// createBenchFile prepares a temporary file fixture inside .tmp/ for repeatable benchmarks.
func createBenchFile(b *testing.B, size int) string {
	b.Helper()

	// Ensure the local .tmp directory exists
	tmpDir := ".tmp"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		b.Fatalf("failed to create .tmp directory: %v", err)
	}

	path := filepath.Join(tmpDir, "bench_data.txt")

	line := "ts=2026-09-23T12:00:00Z level=INFO msg=order_matched price=104.25 qty=50\n"
	repeats := size / len(line)
	data := bytes.Repeat([]byte(line), repeats)

	if err := os.WriteFile(path, data, 0644); err != nil {
		b.Fatalf("failed to create benchmark fixture: %v", err)
	}

	// Optional cleanup: removes the file when the benchmark finishes
	b.Cleanup(func() {
		_ = os.Remove(path)
	})

	return path
}

// 1. BenchmarkReadAll: Measures unbounded heap growth and repeated slice reallocations.
func BenchmarkReadAll(b *testing.B) {
	filePath := createBenchFile(b, benchFileSize)
	b.SetBytes(benchFileSize)
	b.ReportAllocs()

	for b.Loop() {
		data, err := sys.ReadAll(filePath)
		if err != nil {
			b.Fatal(err)
		}
		sinkBytes = data
	}
}

// 2. BenchmarkReadChunk: Direct fixed 32 KB chunk streaming (near-zero heap allocations).
func BenchmarkReadChunk(b *testing.B) {
	filePath := createBenchFile(b, benchFileSize)
	b.SetBytes(benchFileSize)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		if err := sys.ReadChunk(filePath); err != nil {
			b.Fatal(err)
		}
	}
}

// 3. BenchmarkReadBuff: Reusable bytes.Buffer slurp to EOF with buffer pre-allocation.
func BenchmarkReadBuff(b *testing.B) {
	filePath := createBenchFile(b, benchFileSize)
	b.SetBytes(benchFileSize)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		data, err := sys.ReadBuff(filePath)
		if err != nil {
			b.Fatal(err)
		}
		sinkBytes = data
	}
}

// 4. BenchmarkReadBufio: Zero-allocation line parsing via ReadSlice.
func BenchmarkReadBufio(b *testing.B) {
	filePath := createBenchFile(b, benchFileSize)
	b.SetBytes(benchFileSize)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		if err := sys.ReadBufio(filePath); err != nil {
			b.Fatal(err)
		}
	}
}

// 5. BenchmarkPipeRW: Synchronous cross-goroutine streaming over io.Pipe.
// func BenchmarkPipeRW(b *testing.B) {
// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for b.Loop() {
// 		if err := sys.PipeRW(); err != nil {
// 			b.Fatal(err)
// 		}
// 	}
// }
