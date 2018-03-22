package main

import (
	"fmt"
	"go-playground/memsync"
	"net/http"
	"playground/cart"
	"playground/token"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

const port = "9000"

func main() {

	// server()

	// redisCall()

	// slices()

	memsyncCall()
}

func server() {
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

func redisCall() {
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

func slices() {
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

func memsyncCall() {
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
	fmt.Println("Finished all")
	fmt.Println()
}
