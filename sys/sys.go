package sys

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"os"
)

// each write/read is system call, make as little as possible such system calls - use buffers
// buffer size: 4kb - 64kb
// 4kb -> High concurrency, matches OS page size
// 32kb - 64kb -> Large files, reduce system calls

// ------------------------------------------------------------
// 1. io.TeeReader(r, w) - (Stream Duplication / Auditing)
// io.TeeReader(r, w)
// Best for: Calculating hashes (SHA256), logging, or metrics while streaming data to its destination.
// Why: Every byte read from r is automatically written to w with zero intermediate buffer.
// hasher := sha256.New()
// tee := io.TeeReader(fileReader, hasher)

// Upload reads from tee -> data simultaneously flows to S3 and into hasher
// _, err := s3Uploader.Upload(ctx, tee)
// checksum := hex.EncodeToString(hasher.Sum(nil))

// ------------------------------------------------------------

// 2. io.MultiReader(r1, r2, ...) - (Concatenating Streams Without Allocations)
// Best for: Prepending headers, framing bytes, or stitching multiple readers sequentially.
// Why: Avoids copying disparate byte slices into a single giant buffer before sending.
// header := bytes.NewReader([]byte("HEADER_V1\n"))
// body, _ := os.Open("payload.bin")

// combined := io.MultiReader(header, body)
// // Reads header first until EOF, then seamlessly transitions to body
// _, err := io.Copy(networkConn, combined)

// ------------------------------------------------------------

// 3. io.MultiWriter(w1, w2, ...) - (Fan-Out Logging & Multi-Destination Writes)
// Best for: Duplicating a write stream (e.g., writing logs simultaneously to stdout and a file).
// Why: Executes writes to each writer sequentially, returning an error if any fail.
// file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// mw := io.MultiWriter(os.Stdout, file)

// logger := logrus.New()
// logger.SetOutput(mw) // Broadcasts every log line to terminal and disk

// ------------------------------------------------------------

// 4. io.ReadFull(r, buf) - (Strict Fixed-Size Protocol Decoding)
// Best for: Binary protocols, packet headers, or fixed-width message framing.
// Why: Standard r.Read(buf) might read fewer bytes than len(buf). ReadFull guarantees
// the slice is completely filled, or it returns io.ErrUnexpectedEOF.
// header := make([]byte, 8) // Exactly 8-byte message header
// _, err := io.ReadFull(conn, header)
// if err != nil {
//     // If stream closed before reading all 8 bytes -> returns ErrUnexpectedEOF
// }

// ------------------------------------------------------------

// 5. io.LimitReader(r, maxBytes) - (DDoS & Memory Exhaustion Protection)
// Best for: Bounding untrusted HTTP request bodies or file uploads to prevent out-of-memory crashes.
// Why: Wraps r and emits io.EOF once maxBytes is read, preventing malicious clients from streaming gigabytes.
// const maxUploadSize = 10 * 1024 * 1024 // 10 MB limit
// boundedReader := io.LimitReader(r.Body, maxUploadSize)

// data, err := io.ReadAll(boundedReader)

// ------------------------------------------------------------

// ## Tool Selection Matrix ##

// | Requirement | Primary Tool |
// | :--- | :--- |
// | **Stream file in 32–64 KB blocks** | `f.Read(b)` |
// | **Duplicate file efficiently** | `io.Copy(dst, src)` |
// | **Slurp dynamic unknown payload to EOF** | `buff.ReadFrom(r)` |
// | **Line-by-line parsing (idiomatic)** | `bufio.Scanner` |
// | **Line parsing (zero heap allocations)** | `bufio.Reader.ReadSlice('\n')` |
// | **Bridge `io.Writer` to `io.Reader`** | `io.Pipe()` |
// | **Read exact number of bytes** | `io.ReadFull(r, b)` |
// | **Stream hashing or tap inspections** | `io.TeeReader(r, w)` |
// | **Concatenate multiple streams** | `io.MultiReader(r1, r2)` |
// | **Fan out single write to multiple targets** | `io.MultiWriter(w1, w2)` |
// | **Cap payload byte size** | `io.LimitReader(r, limit)` |

// ------------------------------------------------------------

// 1. Raw bytes buffer - zero allocation, read into fixed buffer.
// Best for: High-throughput raw binary streaming, file copies, large/unknown streams, and chunked processing.
// Why: Zero allocation in the loop; performs direct syscalls straight into your user-space
// slice without intermediate buffer copies. Sizing to 32 KB-64 KB hits optimal OS/disk throughput.
func ReadChunk(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)

	for {
		n, err := f.Read(buf) // returns number of bytes read, error
		if n > 0 {
			chunk := buf[:n] // read only the bytes that were read
			_ = chunk        // process chunk
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
	}

	return nil
}

// 2. bytes.Buffer.ReadFrom - dynamic payload, bridging 'io' interfaces.
// Best for: Reading an entire stream of unknown size into memory in a single call.
// Why: Avoids io.ReadAll's small 512B start by allowing upfront capacity via buf.Grow().
// Warning: Allocates per call without a sync.Pool/external buffer; incurs geometric slice doubling if the stream exceeds initial capacity.
func ReadBuff(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var buf bytes.Buffer
	buf.Grow(32 * 1024) // Pre-allocate the 32 KB capacity upfront!

	_, err = buf.ReadFrom(f) // No for-loop needed! ReadFrom automatically reads to EOF
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// 3. bufio.Reader.ReadSlice - reduce system calls, tokenized parsing, lookup.
// Best for: Ultra-high-performance line or token parsing where you cannot afford heap allocations (e.g., hot market feeds, log ingestion).
// Why: Returns a slice referencing bufio's internal buffer directly (zero heap alloc).
// Warning: The slice is invalidated and overwritten on the next call; returns ErrBufferFull if the delimiter isn't found within buffer capacity.
func ReadBufio(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		// 3.1 bufio.Reader.ReadSlice
		// Best for: Ultra-high-performance line or token parsing where you cannot afford heap allocations (e.g., hot market feeds, log ingestion).
		// Why: Returns a slice referencing bufio's internal buffer directly (zero heap alloc).
		// Warning: The slice is invalidated and overwritten on the next call; returns ErrBufferFull if the delimiter isn't found within buffer capacity.
		// line, err := r.ReadSlice('\n')

		// 3.2 bufio.Reader.ReadBytes / ReadString
		// Best for: General line-by-line or delimiter-delimited parsing when lines can exceed buffer size.
		// Why: Handles arbitrarily long lines automatically without failing with ErrBufferFull.
		// Tradeoff: Allocates a new byte slice or string on every single call to safely own the memory.
		// line, err := r.ReadBytes('\n')

		// 3.3 bufio.Reader.Peek / ReadByte
		// Best for: Lexers, protocol decoders, and file signature sniffers.
		// Why: Peek lets you inspect the next N bytes without consuming them or moving the read pointer.
		// ReadByte steps byte-by-byte through a stream without incurring an OS syscall per byte.
		// magicBytes, err := r.Peek(4) // Inspect header without advancing
		// b, err := r.ReadByte()       // Step 1 byte forward from buffer cache

		// 3.4 bufio.Scanner
		// Best for: Idiomatic, clean line-by-line or token-by-token text reading (e.g., CLI tools, scripts, CSV/word splitting).
		// Why: Cleaner API than bufio.Reader; automatically strips delimiters (\n, \r\n).
		// Tradeoff: Defaults to a 64 KB token limit (bufio.MaxScanTokenSize); lines exceeding this require manual buffer configuration.
		// scanner := bufio.NewScanner(f)
		// for scanner.Scan() {
		//     line := scanner.Bytes() // scanner.Text() if you need an allocated string
		//     _ = line
		// }

		line, err := r.ReadSlice('\n')
		if len(line) > 0 {
			// Immediate processing only (zero allocation)
			// if processing in another goroutine, make a copy because the slice/line will be overwritten on next read
			// safeCopy := make([]byte, len(line))
			// copy(safeCopy, line)
			_ = line
		}

		if err != nil {
			if errors.Is(err, bufio.ErrBufferFull) {
				// Line exceeded 4 KB buffer; handle chunk continuation
				continue
			}
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
	}

	return nil
}

// 4. io.ReadAll
func ReadAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

// 5. io.Pipe (pr, pw)
// Rule: Use exclusively when bridging an io.Writer interface (e.g., json.Encoder, gzip.Writer)
// to an io.Reader interface (e.g., http.Request.Body, s3.PutObject) across goroutines.
// Note: If you only need to pass data or byte chunks between goroutines inside your own
// application code, use standard Go channels (chan []byte) instead.
func PipeRW() error {
	pr, pw := io.Pipe()

	// Background producer: writes to pw
	go func() {
		gw := gzip.NewWriter(pw)

		// Stream data directly into gzip -> pipe
		for i := 0; i < 1_000_000; i++ {
			gw.Write([]byte("large continuous streaming data payload\n"))
		}

		// Always close in reverse: child writer first, then pipe writer
		gw.Close()
		pw.Close() // Sends io.EOF to the reader (use pw.CloseWithError(err) on failure)
	}()

	// Consumer: reads directly from pr in streaming chunks
	// Example: s3Client.PutObject(ctx, &s3.PutObjectInput{Body: pr, ...})
	_, err := io.Copy(os.Stdout, pr)
	return err
}

// 6. CopyFile reference comment:
// Best for: Standard file-to-file duplication.
// Why: Lets Go select the fastest path available: kernel zero-copy (copy_file_range)
// on supported platforms, or an internal 32 KB chunk loop fallback.
func CopyFile(from, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	// os.Create internally specifies O_RDWR|O_CREATE|O_TRUNC with permissions 0666
	dst, err := os.Create(to)
	if err != nil {
		return 0, err
	}
	defer func() {
		if closeErr := dst.Close(); err == nil {
			err = closeErr
		}
	}()

	return io.Copy(dst, src)
}
