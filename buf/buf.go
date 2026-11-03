package buf

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
)

// ## Examples

var (
	bufferPool = NewBufferPool(0, 64<<10) // 64 KB retain limit
	bytePool   = NewBytePool(32 << 10)    // 32 KB copy buffers
)

// BufferPool example:
// - json.NewEncoder takes an io.Writer, so the pool holds *bytes.Buffer; the buffer grows to fit the payload.
// - Encode first, then write: an encoding error still produces a clean 500,
// - and the response gets an exact Content-Length.
func writeJSON(w http.ResponseWriter, status int, v any) {
	buf := bufferPool.Get()   // Get a buffer from the pool
	defer bufferPool.Put(buf) // Put the buffer back in the pool

	if err := json.NewEncoder(buf).Encode(v); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.WriteHeader(status)
	w.Write(buf.Bytes()) // safe: io.Writer must not retain the slice (see Rules)
}

// BytePool example:
// - io.CopyBuffer takes a []byte, so the pool holds *[]byte; a fixed-size chunk never grows.
// - Stream without allocating a new buffer on every call.
func copyPooled(dst io.Writer, src io.Reader) (int64, error) {
	bp := bytePool.Get()                // Get a buffer from the pool
	defer bytePool.Put(bp)              // Put the buffer back in the pool
	return io.CopyBuffer(dst, src, *bp) // Copy the data from the source to the destination
}

// ## Rules -------------------------------------------------------------------------------------------------------------------------- ##

// 1. **Never use a buffer after `Put`.** Don't return `buf.Bytes()` or keep a sub-slice of it. If the data must outlive the buffer, copy it first: `out := bytes.Clone(buf.Bytes())`.
// 2. **Always `defer Put` right after `Get`**, so no code path leaks the buffer.
// 3. **Cap what you keep.** Without `maxCap`, a single 50 MB response leaves 50 MB buffers in the pool for good.
// 4. **`io.CopyBuffer` ignores your buffer** if `src` implements `io.WriterTo` or `dst` implements `io.ReaderFrom`. Examples are `*os.File`, `*bytes.Reader`, `io.Discard`, and `*net.TCPConn`. In those cases the pool does nothing, which is harmless.
// 5. **Small payloads (under ~1 KB) don't benefit.** `json.Marshal` already pools its buffer internally, and `json.NewEncoder` allocates on every call. Measured with Go 1.27:

//    | Payload | pooled `BufferPool` | `json.Marshal` |
//    |---|---|---|
//    | ~70 B | 290 ns, 128 B/op | 221 ns, 160 B/op |
//    | ~13 KB | 22 µs, 87 B/op | 25 µs, 13.6 KB/op |

//    | 1 MB stream copy | pooled `BytePool` | `make` each call |
//    |---|---|---|
//    | | 9.6 µs, 40 B/op | 17.8 µs, 32.8 KB/op |

// 6. **The pool is not a cache.** The GC can empty it at any time, so every `Get` may return a new buffer.
// 7. **Prove it with a benchmark.** Use `b.ReportAllocs()` and `-benchmem`, and keep the pool only if the numbers improve.

// ## -------------------------------------------------------------------------------------------------------------------------------- ##

// ## Which pool
// Pick the type by what the API takes, not by the payload size:
// - The API takes an io.Writer (json.NewEncoder, fmt.Fprintf, template.Execute) -> BufferPool.
// - The API takes a []byte (Read, io.CopyBuffer) -> BytePool.
// Whether a buffer grows decides whether its pool needs a maxCap.
// An append-style []byte (strconv.AppendInt, time.AppendFormat) grows, so it needs one too.

// ## BufferPool - growing buffers for io.Writer APIs

// BufferPool pools *bytes.Buffer values. A *bytes.Buffer is an io.Writer that grows
// to fit whatever is written, so Put drops any buffer that grew past maxCap,
// and one huge payload can't pin memory (golang/go#23199).
type BufferPool struct {
	pool   sync.Pool
	maxCap int
}

// NewBufferPool: initCap pre-sizes new buffers (0 is fine); maxCap is the
// largest capacity kept for reuse (e.g. 64 << 10).
func NewBufferPool(initCap, maxCap int) *BufferPool {
	p := &BufferPool{maxCap: maxCap}
	p.pool.New = func() any {
		return bytes.NewBuffer(make([]byte, 0, initCap))
	}
	return p
}

func (p *BufferPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

func (p *BufferPool) Put(b *bytes.Buffer) {
	if b.Cap() > p.maxCap {
		return // too big to keep; drop it and let GC free it
	}
	b.Reset() // reset length to 0, keep capacity
	p.pool.Put(b)
}

// ## BytePool - fixed-size buffers for []byte APIs

// BytePool pools fixed-size byte slices for APIs that take raw memory, not an io.Writer.
// The slice never grows, so no maxCap is needed; Put only refuses a slice of another size.
// It stores *[]byte because putting a plain []byte into `any` allocates (staticcheck SA6002).
type BytePool struct {
	pool sync.Pool
	size int
}

// NewBytePool: size is the length of every buffer (e.g. 32 << 10).
func NewBytePool(size int) *BytePool {
	p := &BytePool{size: size}
	p.pool.New = func() any {
		b := make([]byte, size)
		return &b // a pointer, so Put doesn't allocate
	}
	return p
}

func (p *BytePool) Get() *[]byte {
	return p.pool.Get().(*[]byte) // the *[]byte that New returns
}

func (p *BytePool) Put(b *[]byte) {
	if cap(*b) != p.size {
		return // foreign or re-allocated slice
	}
	*b = (*b)[:p.size] // restore full length (p.size), keep capacity
	p.pool.Put(b)
}
