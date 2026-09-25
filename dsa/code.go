package dsa

import "log"

func Run() {
	Arr()
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

	// 1. set nil,  needed for pointers only, clear the element so GC can reclaim the memory
	arr[0] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [0 2 3 4 5 6 7 8 9 10]

	// 2. reslice
	arr = arr[1:]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 3 4 5 6 7 8 9 10]

	log.Println("-- Remove Last -----------")

	// 1. set nil,  needed for pointers only, clear the element so GC can reclaim the memory
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

	// 2. set nil,  needed for pointers only, clear the element so GC can reclaim the memory
	arr[len(arr)-1] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 9 4 5 6 7 8 0]

	// 3. reslice
	arr = arr[:len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 9 4 5 6 7 8]

	log.Println("-- Remove Middle, Ordered -----------")

	// 1. Shift elements left by 1
	copy(arr[1:], arr[1+1:])
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8 8]

	// 2. Clear the abandoned tail slot
	arr[len(arr)-1] = 0
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8 0]

	// 3. Truncate
	arr = arr[:len(arr)-1]
	log.Println(arr)
	// 2026/09/25 12:19:11 [2 4 5 6 7 8]
}
