package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

// As of Go 1.8, garbage collection pauses are generally between 10 and 100 microseconds!

type A struct {
	Val string
}

type B struct {
	AA *A
	AB A
}

func modPtr(a *A, wg *sync.WaitGroup) {
	defer wg.Done()
	a.Val = "changed"
}

func mod(a A, wg *sync.WaitGroup) {
	defer wg.Done()
	a.Val = "changed"
}

func modPtr2(a *A, wg *sync.WaitGroup) {
	defer wg.Done()
	aa := a
	aa.Val = "oops"
}

func mod2(a A, wg *sync.WaitGroup) {
	defer wg.Done()
	aa := a
	aa.Val = "oops"
}

func main() {
	logrus.Info("playground")

	dataRace()
}

func dataRace() {
	a := &A{
		Val: "abc",
	}

	b := &B{
		AA: a,
		AB: A{
			Val: "bcd",
		},
	}

	wg := sync.WaitGroup{}

	wg.Add(1000)
	// pass by value will never cause data-race, unless inner struct is pointer and changing it!!
	// pass by pointer will always cause data-race, even for inner structs!
	for i := 0; i < 1000; i++ {
		// data-race because passed a is pointer
		//go func(a *A, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//	a.Val = "a"
		//}(a, &wg)

		// no data-race because passed a is value (not pointer)
		//go func(a A, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//	a.Val = "a"
		//}(*a, &wg)

		// ---

		// data-race because passed a is pointer
		//go modPtr(a, &wg)

		// no data-race because passed a is value (not pointer)
		//go mod(*a, &wg)

		// ---

		// data-race because aa holds pointer of a
		//go func(a *A, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//
		//	aa := a
		//	aa.Val = "oops"
		//}(a, &wg)

		// no data-race because aa holds value of a (not pointer)
		//go func(a A, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//
		//	aa := a
		//	aa.Val = "oops"
		//}(*a, &wg)

		// ---

		// data-race because aa holds pointer of a
		//go modPtr2(a, &wg)

		// no data-race because aa holds value of a (not pointer
		//go mod2(*a, &wg)

		//go func(b *B, wg *sync.WaitGroup) {
		//	defer wg.Done()
		//
		//	// data-race, b is pointer, we still change b!
		//	//b.AA.Val = "uh"
		//
		//	// data-race, b is pointer, we still change b!
		//	//b.AA = &A{
		//	//	Val: "trs",
		//	//}
		//
		//	// data-race, aa is pointer of b.AA -> b is pointer!! we still change b!
		//	//aa := b.AA
		//	//aa.Val = "pfw"
		//
		//	// data-race, b is pointer!!! doesnt matter if AB is not pointer, we still change b!
		//	//b.AB.Val = "f"
		//
		//	// no data-race, aa holds value of b.AA, AA is not pointer - dereference!!
		//	//aa := *b.AA
		//	//aa.Val = "s"
		//
		//	// no data-race, ab holds value of b.AB, AB is not pointer!
		//	//ab := b.AB
		//	//ab.Val = "tr"
		//}(b, &wg)

		go func(b B, wg *sync.WaitGroup) {
			defer wg.Done()

			// data-race, b.AA is pointer
			//b.AA.Val = "wv"

			// no data-race, b.AB is not pointer
			b.AB.Val = "fw"
		}(*b, &wg)
	}

	wg.Wait()
}

func arrr() {
	arr := make([]int, 2, 6) // [0, 0], len 2, cap 6
	arr2 := make([]int, 0)   // empty array
	arr3 := make([]int, 2)   // [0, 0], len 2, cap 2

	logrus.Info(arr)
	logrus.Info(arr2)
	logrus.Info(arr3)

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	logrus.Info(arr)

	logrus.Info(arr[1:5])

	// copy rest of the array
	copy(arr[2:], arr[2+1:]) // dst: 3, 4..., start: 4, 5...
	//arr[len(arr)-1] = nil
	arr = arr[:len(arr)-1]

	logrus.Info(arr)

	// replace with last element
	arr[3] = arr[len(arr)-1]
	//arr[len(arr)-1] = nil
	arr = arr[:len(arr)-1]

	logrus.Info(arr)
}

// we can use context to close goroutine - instead of done chan
func contextToClose(ctx context.Context) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for i := 0; ; i++ {
			select {
			case result <- i:
				// do nothing
			case <-ctx.Done():
				return
			}
		}
	}()

	return result
}

func mem() {
	var s1 = &s{
		val:  "s1",
		val2: 1,
	}

	var s2 = s{
		val:  "s2",
		val2: 2,
	}

	sn := modify(*s1)
	logrus.Info(s1)
	logrus.Info(sn)
	sn.val = "changed again"
	sn.val2 = 2222222

	logrus.Info(s1)
	logrus.Info(sn)

	sn = modify(s2)
	logrus.Info(s2)
	logrus.Info(sn)
	sn.val = "changed again 2222"
	sn.val2 = 4444444

	logrus.Info(s2)
	logrus.Info(sn)

	si := &impl{}
	run(si)
}

type runner interface {
	run()
}

type impl struct{}

func (i *impl) run() {}

func run(r runner) {
	var i interface{} = r

	switch i.(type) {
	case int:
		logrus.Info("its int")
	case *impl:
		logrus.Info("its *impl")
	}
}

func modify(s s) *s {
	s.val = "s changed"
	s.val2 = 11111

	return &s
}

type s struct {
	val  string
	val2 int
}

func memSyncEx() {
	var counter int
	var cLock sync.Mutex

	increment := func(i int) {
		cLock.Lock()
		defer cLock.Unlock()

		fmt.Println("Incrementing from:" + strconv.Itoa(counter))

		counter += i

		fmt.Println("Incremented to: " + strconv.Itoa(counter))
	}

	decrement := func(i int) {
		cLock.Lock()
		defer cLock.Unlock()

		fmt.Println("Decrementing from: " + strconv.Itoa(counter))

		time.Sleep(3 * time.Second)
		counter -= i

		fmt.Println("Decremented to: " + strconv.Itoa(counter))
	}

	var wgCounter sync.WaitGroup
	wgCounter.Add(1)
	go func() {
		defer wgCounter.Done()

		increment(3)
	}()

	wgCounter.Add(1)
	go func() {
		defer wgCounter.Done()

		decrement(2)
	}()
	wgCounter.Wait()

	fmt.Println()
	fmt.Println("Finished all")
	fmt.Println()
}

func condEx() {
	// when two or more goroutines wait for signals to occur
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(msg string) {
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from " + msg + " queue")
		fmt.Println("Total items in " + msg + " queue: " + strconv.Itoa(len(queue)))
		fmt.Println()
		c.L.Unlock()

		// 2.) send signal
		c.Signal()
	}

	addToQueue := func(nItems int, qMax int, msg string) {
		for i := 0; i < nItems; i++ {
			c.L.Lock()

			// blocking part until condition becomes false
			for len(queue) == qMax {
				fmt.Println("Waiting for more space -> Total items in " + msg + " queue: " + strconv.Itoa(len(queue)))
				fmt.Println()
				// 1.) wait until we are signaled
				c.Wait()
			}
			// end

			fmt.Println("Adding to " + msg + " queue")
			queue = append(queue, struct{}{})
			fmt.Println("Item " + strconv.Itoa(i) + " added in " + msg + " queue")
			fmt.Println("Total items in " + msg + " queue: " + strconv.Itoa(len(queue)))
			fmt.Println()

			// remove from queue after 2 seconds or some other cicrumstances/conditions
			time.Sleep(2 * time.Second)

			// 3.) repeat the process
			go removeFromQueue(msg)

			c.L.Unlock()
		}
	}

	addToQueue(5, 2, "First")
}

func broadcastCondEx() {
	// sync.Once
	// Will make sure function passed to once.Do(fn) will run only ONCE, even in different goroutines
	// calls to once.Do(fn) and once.Do(fn2) will happen only ONCE (in first call to once.Do(fn))
	// so, once.Do() will run only once, no matter how many different functions we pass,
	// or how many calls we make to once.Do()

	// sync.Pool
	// when frequently allocating many objects of the same type

	// example 1

	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, run func()) {
		var running sync.WaitGroup

		running.Add(1)

		go func() {
			running.Done()

			c.L.Lock()

			c.Wait()

			go run()

			c.L.Unlock()
		}()

		running.Wait()
	}

	subscribe(button.Clicked, func() {
		fmt.Println("Action 1 started")
		time.Sleep(3 * time.Second)
		fmt.Println("Action 1 finished")
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Action 2 started")
		time.Sleep(1 * time.Second)
		fmt.Println("Action 2 finished")
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Action 3 started")
		time.Sleep(2 * time.Second)
		fmt.Println("Action 3 finished")
	})

	button.Clicked.Broadcast()

	// example 2

	fmt.Println("example 2")
	type Downloader struct {
		Started *sync.Cond
	}
	downloader := &Downloader{Started: sync.NewCond(&sync.Mutex{})}

	download := func(started *sync.WaitGroup, downloader *sync.Cond, path string) {
		logrus.Info("download about to started: ", path)

		time.Sleep(2 * time.Second)

		go func() {
			logrus.Info("started: ", path)
			started.Done()

			downloader.L.Lock()
			downloader.Wait()

			go func() {
				logrus.Info("downloading from: ", path)
				time.Sleep(2 * time.Second)
				logrus.Info("downloaded from: ", path)
			}()

			downloader.L.Unlock()
		}()

	}

	var started sync.WaitGroup
	started.Add(3)

	go download(&started, downloader.Started, "url.com/1")
	go download(&started, downloader.Started, "url.com/2")
	go download(&started, downloader.Started, "url.com/3")

	started.Wait()

	logrus.Warn("broadcasting")
	downloader.Started.Broadcast()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("All finished")
}

func channelsEx() {
	dataStream := make(chan string)

	exec := func(msg string) {
		fmt.Println("Sending big data...")
		time.Sleep(2 * time.Second)
		dataStream <- msg
	}

	go exec("Data 1")

	fmt.Println("Data received: " + <-dataStream)
	fmt.Println()

	// ###################################

	runner := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// blocking part, until the channel gets closed
			<-runner
			fmt.Println(i, " running")
		}(i)
	}

	// useful to run/free/unblock many goroutines at once
	// just like sync.Cond broadcast()
	fmt.Println("Unblocking goroutines...")
	close(runner)

	wg.Wait()
}

func bufferedChanEx() {
	// performance boost if we know number of elements to send
	// make buffer close (big) enough, so it "looks like" all data has been sent/written "at once"

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)

		for i := 0; i < 5; i++ {
			// time.Sleep(1 * time.Second)
			fmt.Println("Sent: " + strconv.Itoa(i))
			intStream <- i
		}

		fmt.Println()
		fmt.Println("Sent all")
		fmt.Println()
	}()

	for i := range intStream {
		fmt.Println("Received: " + strconv.Itoa(i))
	}
}

func selectChanEx() {
	c1 := make(chan string)
	c2 := make(chan string)

	start := time.Now()

	go func() {
		time.Sleep(300 * time.Millisecond)
		c1 <- "Data 1"
		close(c1)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "Data 2"
		close(c2)
	}()

	running := true
	for running {
		select {
		case <-c1:
			c := <-c1
			fmt.Println("From c1: " + c)
		case <-c2:
			c2Data := <-c2
			fmt.Println("From c2: " + c2Data)
		default:
			// fmt.Println("Waiting for data...")
		}

		go func() {
			// fmt.Println("Do something else in meantime")
		}()

		if time.Since(start) >= 500*time.Millisecond {
			running = false
		}
	}

	fmt.Println("End")
}

func leakingGoroutineEx() {
	// let caller decide when to stop our clojure by doing close(done)
	// quote:
	// If a goroutine is responsible for creating a goroutine,
	// it is also responsible for ensuring it can stop the goroutine

	// Leaking example
	doSmtn := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		defer fmt.Println("doSmtn exited")

		go func() {
			fmt.Println("doSmtn clojure")

			defer fmt.Println("doSmtn exited")
			// never called! goroutine remains in background, hence leaking
			defer close(completed)

			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()

		fmt.Println("Last line in doSmtn")
		return completed
	}
	doSmtn(nil)
	// Perhaps more work is done here
	fmt.Println("After doSmtn")
	fmt.Println()
	// End

	// pass done <-chan so we can get notified/signaled to stop
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		// without terminated chan, we could not "block" main/parent goroutine,
		// it would exit before this clojure completes its job
		// this way main/parent goroutine knows there is a pending/running gorouting, so wait for it
		terminated := make(chan interface{})
		defer fmt.Println("doWork exited")

		go func() {
			fmt.Println("Started clojure")
			// if nil passed for strings, this part will never be reached
			defer fmt.Println("doWork exited")
			// let caller (parent goroutine) know we finished, unblock blocking part
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// Do something with strings
					fmt.Println(s)
				case <-done: // will be called only on close(done) - its caller's job to do so
					// the only way to make defer close(terminated) gets called
					return
				}
			}
		}()

		// Goroutine continues doing its job (leaking clojure in case of nil strings <-chan, because <-done never happened)
		fmt.Println("Last line in doWork")
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(2 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	fmt.Println("Blocking part until close(terminated) gets called")
	<-terminated
	fmt.Println("Done")

	// ################################

	fmt.Println()

	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done: // when notified, stop this goroutine and prevent infinite leaking
					// the only way to make defer close(randStream) gets called
					return
				}
			}
		}()

		return randStream
	}

	done = make(chan interface{})
	randStream := newRandStream(done)

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	// notify newRandStream we finished with it
	close(done)

	fmt.Println("Do some other work")
	time.Sleep(3 * time.Second)
	fmt.Println("Finished")
}

// concurrent function body
func chanEx1() {
	download := func(done <-chan interface{}) <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)

			select {
			case <-done:
				return
			default:
				for i := 0; i < 2; i++ {
					fmt.Printf("\nSending: %v", i)
					time.Sleep(time.Second * 2)
					fmt.Printf("\nSent: %v", i)
					result <- i
				}
			}
		}()

		return result
	}

	// Usage of download():
	// Since we call download() we make sure we close it too
	// We do so by passing done chan to download(), which is then handled in download()
	done := make(chan interface{})

	received := download(done)

	for r := range received {
		fmt.Printf("\nReceived: %v", r)
	}

	close(done)
}

// concurrently called, communication via channels
func chanEx2() {
	download := func(path string, rec *chan string) {
		fmt.Printf("\nDownloading: %v", path)
		time.Sleep(3 * time.Second)
		*rec <- path
	}

	handleDownload := func(rec *chan string) string {
		downloaded := <-*rec
		fmt.Printf("\nHandling downloaded data: %v", downloaded)
		return downloaded
	}

	// Usage of download():
	rec := make(chan string)
	go download("path1", &rec)

	fmt.Printf("\nWaiting for download...\n")
	downloaded := handleDownload(&rec)
	fmt.Printf("\nDownloaded data: %v", downloaded)
}

// how to properly handle errors in channels
func chanEx3() {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get(url)

				result := Result{Error: err, Response: resp}

				select {
				case <-done:
					return
				case results <- result:
					// do nothing but write to result to results chan
				}
			}
		}()

		return results
	}

	done := make(chan interface{})
	urls := []string{"https://www.google.com", "https://badhost"}

	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("\nError: %v", result.Error)
			continue
		}
		fmt.Printf("\nResponse: %v\n", result.Response.Status)
	}
	close(done)
}

// generator for a pipeline is any function that converts a set of discrete values into a stream of values on a channel
func chanEx4() {
	// this function will repeat func call you pass in infinitely until you tell it to stop
	repeat := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		outStream := make(chan interface{})

		go func() {
			defer close(outStream)

			for {
				select {
				case <-done:
					return
				case outStream <- fn():
				}
			}
		}()

		return outStream
	}

	// take first n values from readStream
	take := func(done <-chan interface{}, readStream <-chan interface{}, n int) <-chan interface{} {
		outStream := make(chan interface{})

		go func() {
			defer close(outStream)

			for i := 0; i < n; i++ {
				select {
				case <-done:
					return
				case outStream <- <-readStream:
					// read from readStream (<-readStream will return first value) and store in outStream
				}
			}
		}()

		return outStream
	}

	done := make(chan interface{})
	repeated := repeat(done, func() interface{} {
		return rand.Int()
	})

	for num := range take(done, repeated, 5) {
		fmt.Printf("\nNum: %v", num)
	}

	close(done)
}

func chanReadStream() {
	done := make(chan bool)

	download := func(done <-chan bool, path string, stream <-chan string) <-chan []byte {
		completed := make(chan []byte)

		go func() {
			defer func() {
				logrus.Info("closed")
				close(completed)
			}()

			// infinitely read from stream, until close(done)
			for {
				select {
				case <-done:
					return
				case s := <-stream:
					completed <- []byte(s)
				case <-time.After(300 * time.Millisecond):
					completed <- []byte(path)
				}
			}
		}()

		return completed
	}

	stream := make(chan string)

	// mock streaming
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				stream <- "data " + fmt.Sprint(rand.Intn(9999))
			}
		}
	}()

	result := download(done, "some/path/1", stream)

	// mock closing done channel, stop reading from stream
	go func() {
		select {
		case <-time.After(3 * time.Second):
			close(done)
		}
	}()

	logrus.Info("reading")

	// handle results
	for r := range result {
		logrus.Info("result_1: ", string(r))
	}
}

func ctxs() {
	//list := []string{
	//	"localhost:9090",
	//	"localhost:9091",
	//	"localhost:9092",
	//}
	//go func() {
	//	listener, err := net.Listen("tcp", list[0])
	//	if err != nil {
	//		log.Fatalln("Listener:", list[0], err)
	//	}
	//	time.Sleep(time.Second * 5)
	//	c, err := listener.Accept()
	//	if err != nil {
	//		log.Fatalln("Listener:", list[0], err)
	//	}
	//	defer c.Close()
	//}()
	//ctx, canc := context.WithTimeout(context.Background(), time.Second*10)
	//defer canc()
	//wg := sync.WaitGroup{}
	//wg.Add(len(list))
	//for _, addr := range list {
	//	go func(addr string) {
	//		defer wg.Done()
	//		conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", addr)
	//		if err != nil {
	//			log.Println("-> Connection:", err)
	//			return
	//		}
	//		log.Println("-> Connection to", addr, "cancelling context")
	//		canc()
	//		conn.Close()
	//	}(addr)
	//}
	//wg.Wait()

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	//defer cancel()
	//
	//cc(ctx)
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		break
	//	default:
	//		continue
	//	}
	//
	//	break
	//}
	//
	//logrus.Info("main completed")

	//ctx, cancel := context.WithCancel(context.Background())
	//time.AfterFunc(time.Second*10, cancel)
	//done := ctx.Done()
	//for i := 0; ; i++ {
	//	select {
	//	case <-done:
	//		fmt.Println("exit", ctx.Err())
	//		return
	//	case <-time.After(time.Second):
	//		fmt.Println("tick", i)
	//	}
	//}

	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	//defer cancel()

	//time.AfterFunc(time.Second*10, cancel)
	//done := ctx.Done()
	//for i := 0; ; i++ {
	//	select {
	//	case <-done:
	//		fmt.Println("exit", ctx.Err())
	//		return
	//	case <-time.After(time.Second):
	//		fmt.Println("tick", i)
	//	}
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//time.AfterFunc(time.Second*10, cancel)
	//done := ctx.Done()
	//for i := 0; ; i++ {
	//	select {
	//	case <-done:
	//		fmt.Println("exit", ctx.Err())
	//		return
	//	case <-time.After(time.Second):
	//		fmt.Println("tick", i)
	//	}
	//}

	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	//gen := func(ctx context.Context) <-chan int {
	//	dst := make(chan int)
	//	n := 1
	//	go func() {
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				return // returning not to leak the goroutine
	//			case dst <- n:
	//				n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers
	//
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}

	//const addr = "localhost:8080"
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(time.Second * 5)
	//})
	//go func() {
	//	if err := http.ListenAndServe(addr, nil); err != nil {
	//		log.Fatalln(err)
	//	}
	//}()
	//req, _ := http.NewRequest(http.MethodGet, "http://"+addr, nil)
	//ctx, canc := context.WithTimeout(context.Background(), time.Second*2)
	//defer canc()
	//time.Sleep(time.Second)
	//if _, err := http.DefaultClient.Do(req.WithContext(ctx)); err != nil {
	//	log.Fatalln(err)
	//}
}

func cc(ctx context.Context) {
forloop:
	for {
		select {
		case <-ctx.Done():
			logrus.Warn("request canceled")
			break forloop
		case <-time.After(time.Second):
			logrus.Info("working...")
		}
	}

	logrus.Info("returned")
}

func rd() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("starting dir:", wd)

	if err = os.Chdir("/"); err != nil {
		logrus.Fatal(err)
	}

	if wd, err = os.Getwd(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("final dir:", wd)
}

func buf() {
	str := strings.NewReader("my very long string")
	buff := bytes.NewBuffer(make([]byte, 16))

	for {
		p := make([]byte, 16)

		_, err := str.Read(p)

		if err == io.EOF {
			break
		}

		buff.Write(p)
	}

	logrus.Info(buff.String())
}

func bufbuf() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout

	for {
		s.Scan()

		msg := string(s.Bytes())
		if msg == "exit" {
			return
		}

		if msg != "" {
			_, err := fmt.Fprint(w, "You wrote: "+msg)
			if err != nil {
				logrus.Error(err)
				return
			}
		}
	}
}

func ctxx() {
	ctx := context.Background()
	src := sourceLine(ctx, ioutil.NopCloser(strings.NewReader("some random text")))
	filter := textFilter(ctx, src, "some")
	printer(ctx, filter, 31, "some", os.Stdout)
}

func sourceLineWords(ctx context.Context, r io.ReadCloser) <-chan []string {
	ch := make(chan []string)

	go func() {
		defer func() {
			r.Close()
			close(ch)
		}()

		b := bytes.Buffer{}
		sc := bufio.NewScanner(r)

		for sc.Scan() {
			b.Reset()
			b.Write(sc.Bytes())

			var words []string
			w := bufio.NewScanner(&b)
			w.Split(bufio.ScanWords)

			for w.Scan() {
				words = append(words, w.Text())
			}

			select {
			case <-ctx.Done():
				return
			case ch <- words:
			}
		}
	}()

	return ch
}

// each func result is input for next function in the line, func = stage
// 1.) read from reader, pass to next in line
func sourceLine(ctx context.Context, r io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer func() {
			r.Close()
			close(ch)
		}()

		s := bufio.NewScanner(r)

		for s.Scan() {
			select {
			case <-ctx.Done():
				return
			case ch <- s.Text():
			}
		}
	}()

	return ch
}

// 2.) next in line, filtering from previous reader, send to next in line, printer
func textFilter(ctx context.Context, src <-chan string, filter string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for v := range src {
			if !strings.Contains(v, filter) {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}

// 3.) print filtered data, last in line
func printer(ctx context.Context, src <-chan string, color int, highlight string, w io.Writer) {
	const close = "\x1b[39m"
	open := fmt.Sprintf("\x1b[%dm", color)

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-src:
			if !ok {
				return
			}
			i := strings.Index(v, highlight)
			if i == -1 {
				panic(v)
			}
			fmt.Fprint(w, v[:i], open, highlight, close, v[i+len(highlight):], "\n")
		}
	}
}
