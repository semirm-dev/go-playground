package memsync

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Data to perform tests on
type Data struct {
	m   sync.Mutex
	Val int
}

// New will create new data model
func New() *Data {
	d := new(Data)

	return d
}

// AddVal will add/increment new value to val
func (d *Data) AddVal(n string, v int, wg *sync.WaitGroup) {
	fmt.Println("Start AddVal: " + n)

	d.m.Lock()
	defer d.m.Unlock()

	// mark fork/child as finished, create join point
	defer wg.Done()

	defer func() {
		fmt.Println("Finished AddVal: " + n)
		fmt.Println("After " + n + " AddVal: " + strconv.Itoa(d.Val))
	}()

	d.Val += v

	time.Sleep(1500 * time.Millisecond)
}
