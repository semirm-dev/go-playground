# When to use pooled buffers (`sync.Pool`)

## Use a pool when all of these are true

1. **Hot path.** The code runs thousands of times per second or more: request handlers, encoders, loggers, message parsers.
2. **Medium or large buffers.** About 1 KB and up. Below that, the pool's overhead cancels out the gain.
3. **Short, bounded lifetime.** `Get` and `Put` happen in the same function (use `defer`), and the data doesn't outlive it.
4. **A profile shows it matters.** Allocations or GC show up in measurements, not just in intuition.

If any one of these is false, you probably don't need a pool.

## Good fits

| Situation | Pool type |
|---|---|
| Encoding medium or large responses (JSON, CSV, protobuf) | `*bytes.Buffer` pool with a `maxCap` |
| Streaming through a wrapper: hashing, gzip, upload processing, proxying | Fixed-size `*[]byte` scratch pool + `io.CopyBuffer` |
| Custom logger or formatter building lines with `append` | Growable `*[]byte` pool (len 0) with a `maxCap` |
| Parsing network messages or frames in a read loop | Fixed-size `*[]byte` scratch pool |

## Don't use a pool when

- **Payloads are small** (under ~1 KB). `json.Marshal` and `fmt.Sprintf` are fine.
- **The standard library already pools internally.** `json.Marshal` and `fmt` do.
- **The data must escape** the function (returned, sent on a channel, stored). Copying it out defeats the purpose.
- **You can stream instead.** `json.NewEncoder(w)` and `io.Copy` avoid holding the payload at all, which beats pooling it.
- **Payloads are regularly huge** (MBs). Stream them, and cap the input with `http.MaxBytesReader`.
- **The code is a cold path, CLI, or one-off script.** The extra complexity buys nothing.
- **You haven't measured.** A pool adds rules and subtle bugs if they're broken. Pay that cost only for a measured gain.

## How to find out if you need one

```bash
go test -bench=. -benchmem                                   # B/op, allocs/op
go tool pprof -sample_index=alloc_space http://localhost:6060/debug/pprof/heap
GODEBUG=gctrace=1 ./yourservice                              # GC frequency and CPU share
```

Signs you need one:

- A buffer allocation is near the top of `alloc_space`.
- GC takes about 10% or more of CPU.
- p99 latency spikes line up with GC cycles.

Benchmark again after adding the pool. Keep it only if the numbers improve.

## Rules once you use one

1. **`defer Put` right after `Get`.**
2. **Never touch a buffer after `Put`.** That includes slices from `buf.Bytes()`. If the data must outlive the buffer, copy it with `bytes.Clone`.
3. **One owner at a time.** The pool is safe to call from many goroutines; the buffers are not. To hand a buffer to another goroutine, transfer ownership: the receiver calls `Put`, and the sender stops using it.
4. **Never `Put` twice.** Two future `Get` calls would return the same buffer to two goroutines at once.
5. **Cap what you keep.** Drop buffers that grew past `maxCap` (about 2× your p99 payload size) so rare huge payloads don't stay in memory.
6. **The pool is not a cache.** The GC empties it, so every `Get` may return a new buffer.
7. **Test with `go test -race`,** including a test that exercises the pooled path from many goroutines at once.

## Short version

Write plain code first. Profile under realistic load. Pool only buffers that show up in a hot path, and confirm with a benchmark.
