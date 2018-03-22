package main

import (
	"fmt"
	"net/http"
	"playground/cart"
	"playground/token"

	"github.com/go-redis/redis"
)

const port = "9000"

func main() {

	// server()

	// redisTest()

	slices()
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

func redisTest() {
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
