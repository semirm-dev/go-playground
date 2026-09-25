package dsa

import "log"

func Run() {
	Arr()
	Str()
}

// Big-O measures how execution time or memory footprint scales relative to the input size (n).
// O(1) < O(log n) < O(n) < O(n log n) < O(n^2) < O(2^n) < O(n!)
// 3 rules to simplify Big-O:
// 1. Drop Constant Factors: O(2n) -> O(n), O(500) -> O(1).
// 2. Drop Lower-Order Terms: O(n^2 + 5n + 100) -> O(n^2).
// 3. Multi-Variable Inputs: Separate distinct input dimensions. A function iterating through slice A (length n) and slice B (length m) runs in O(n + m) or O(n * m), not O(n).
func BigO() {

}

func Arr() {
	// from = inclusive, to = exclusive
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println("-- Remove First -----------")

	// fifo, queue pop / dequeue
	item := arr[0]
	log.Println("popped item:", item)

	// 1. set nil,  needed for pointers only, clear the element so GC can reclaim the memory
	arr[0] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [0 2 3 4 5 6 7 8 9 10]

	// 2. reslice
	arr = arr[1:]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 3 4 5 6 7 8 9 10]

	log.Println("-- Remove Last -----------")

	// lifo, stack pop
	item = arr[len(arr)-1]
	log.Println("popped item:", item)

	// 1. set nil, needed for pointers only, clear the element so GC can reclaim the memory
	arr[len(arr)-1] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 3 4 5 6 7 8 9 0]

	// 2. reslice
	arr = arr[:len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 3 4 5 6 7 8 9]

	log.Println("-- Remove Middle, Unordered -----------")

	// 1. replace index with last element
	arr[1] = arr[len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 9 4 5 6 7 8 9]

	// 2. set nil, needed for pointers only, clear the element so GC can reclaim the memory
	arr[len(arr)-1] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 9 4 5 6 7 8 0]

	// 3. reslice
	arr = arr[:len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 9 4 5 6 7 8]

	log.Println("-- Remove Middle, Ordered -----------")

	// 1. shift elements left by 1 to override the element to remove
	copy(arr[1:], arr[1+1:])
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8 8]

	// 2. set nil, needed for pointers only, clear the element so GC can reclaim the memory
	arr[len(arr)-1] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8 0]

	// 3. reslice
	arr = arr[:len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8]

	log.Println("--------------------------------")
}

func Str() {
	str := "hello123" // 'hello' = letters, '123' = digits
	i := 0
	s := 6
	log.Println("char 'h' = ", str[i], "(ascii), char 'a' = ", 'a', "(ascii), digit '2' = ", str[s], "(ascii)")

	// letter => ascii
	// 'a' = 97, 'b' = 98, 'c' = 99, ... 'h' = 104, ... 'z' = 122
	// digit => ascii
	// '0' = 48, '1' = 49, '2' = 50, ... '9' = 57

	// index in alphabet (0-25) => letter
	// 0: 'a', 1: 'b', 2: 'c', ... 7: 'h', ... 25: 'z'

	// letters
	// ascii - ascii = index (7)
	// ascii + index = ascii (104) -> then to string(104)

	// convert letter -> index
	idx := str[i] - 'a' // 104 - 97 = 7 (index)
	log.Println("index in alphabet:", idx)

	// convert index -> ascii
	var ch byte = 'a' + idx // 97 + 7 = 104 (ascii)
	log.Println("ascii:", ch, "-> string:", string(ch))

	// digits
	// ascii - ascii = digit (2)
	// ascii + digit = ascii (50)

	// convert ascii -> digit
	var d int = int(str[s] - '0')
	log.Println("'50' (ascii) -> digit:", d) // 50 - 48 = 2 (digit)
}
