package main

import (
	"fmt"
	"go-playground/cart"
	"go-playground/memsync"
	"go-playground/token"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

const port = "9000"

func main() {

	// serverEx()

	// redisEx()

	// slicesEx()

	// forkJoinEx()

	// memSyncEx()

	// condEx()

	// broadcastCondEx()

	// channelsEx()

	// bufferedChanEx()

	// selectChanEx()

	leakingGoroutineEx()
}

func serverEx() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		t := token.New()

		err := t.Generate(map[string]string{
			"username": "userx",
			"id":       "123",
		})

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(t.Token))
	})

	r.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// tmp test token (from local storage, cookies...)
		t := token.New()

		err := t.Generate(map[string]string{
			"username": "userx",
			"id":       "123",
		})

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		// end

		if t.Valid(t.Token) {
			w.Write([]byte("Protected route: " + t.Token))
			return
		}

		w.Write([]byte("Wrong token"))
	})

	fmt.Println("Listening on port: " + port)
	http.ListenAndServe(":"+port, r)
}

func redisEx() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	err = client.Set("user", "userx", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func slicesEx() {
	c := cart.New("Cart 1")

	i := cart.NewItem("Item 1", 123)
	i3 := cart.NewItem("Item 3", 222)
	i4 := cart.NewItem("Item 4", 444)

	c.Add(i)
	c.Add(i4)
	c.Add(&cart.Item{
		Name:  "Item 2",
		Price: 111,
	})
	c.Add(i3)

	fmt.Println("Before delete:")
	fmt.Println()
	for _, el := range c.Items {
		fmt.Println(el)
	}

	c.Remove(i3)
	c.Remove(i)

	fmt.Println()
	fmt.Println("-----------------------")
	fmt.Println()

	fmt.Println("After delete:")
	fmt.Println()

	for _, el := range c.Items {
		fmt.Println(el)
	}
}

func forkJoinEx() {
	var wg sync.WaitGroup
	ms := memsync.New()

	fmt.Println("Initial: " + strconv.Itoa(ms.Val))
	fmt.Println()

	wg.Add(3)

	// fork-join model
	// go will ensure "fork" part, means the code will "create a child" and run on its own
	go ms.AddVal("First", 2, &wg)
	go ms.AddVal("Second", 2, &wg)

	msg := "Hello"
	go func() {
		defer wg.Done()

		msg = "Changed"
	}()

	// this will ensure "join" part, means the main goroutine will not exit until all forks/childs are done
	// make sure all forks/childs joined back
	wg.Wait()

	fmt.Println()
	// value changed because goroutines execute within the same address space they were created in
	fmt.Println("Msg: " + msg)

	defer func() {
		fmt.Println("Total AddVal: " + strconv.Itoa(ms.Val))
	}()

	fmt.Println()
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
	// when we want to notify all registered handlers - which are probably blocked until signaled
	// unblock all blocking goroutines
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	// sync.Once
	// Will make sure function passed to once.Do(fn) will run only ONCE, even in different goroutines
	// calls to once.Do(fn) and once.Do(fn2) will happen only ONCE (in first call to once.Do(fn))
	// so, once.Do() will run only once, no matter how many different functions we pass,
	// or how many calls we make to once.Do()

	// sync.Pool
	// when frequently allocating many objects of the same type

	// condition to register, and when c *sync.Cond occurs do the following...
	subscribe := func(c *sync.Cond, run func()) {
		var running sync.WaitGroup

		running.Add(1)

		go func() {
			// confirmation that goroutine has started
			running.Done()

			// wait until broadcast gets triggered and then run the run() function
			c.L.Lock() // obtain a lock
			defer c.L.Unlock()
			fmt.Println("Waiting for signal...")
			c.Wait()

			fmt.Println("Before run")
			run()
			fmt.Println("After run")
		}()

		running.Wait()
	}

	var wg sync.WaitGroup

	wg.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Action 1")
		time.Sleep(3 * time.Second)
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Action 2")
		time.Sleep(3 * time.Second)
		wg.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Action 3")
		time.Sleep(3 * time.Second)
		wg.Done()
	})

	// make the button.Clicked (c *sync.Cond in subscribe) signal occur
	button.Clicked.Broadcast()

	wg.Wait()

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
