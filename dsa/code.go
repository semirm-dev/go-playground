package dsa

// Big-O measures how execution time or memory footprint scales relative to the input size (n).
// O(1) < O(log n) < O(n) < O(n log n) < O(n^2) < O(2^n) < O(n!)
// 3 rules to simplify Big-O:
// 1. Drop Constant Factors: O(2n) -> O(n), O(500) -> O(1).
// 2. Drop Lower-Order Terms: O(n^2 + 5n + 100) -> O(n^2).
// 3. Multi-Variable Inputs: Separate distinct input dimensions. A function iterating through slice A (length n) and slice B (length m) runs in O(n + m) or O(n * m), not O(n).
func BigO() {

}
