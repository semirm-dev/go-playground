# 🚀 The Master DSA & Systems Architecture Roadmap

Passing an algorithmic interview requires a different mental model than building a production system. This roadmap starts with **foundations built from scratch**, then works through the **NeetCode roadmap patterns** with real LeetCode problems, and connects them to **low-level mechanics** where it's genuinely relevant.

> **Data structures and LeetCode are not separate stages.** The patterns *are* the algorithms. You don't master a structure by implementing it once — you master it by using it on problems you haven't seen. Build the foundations in Phase 1, then solve problems continuously from week 4 onward.

---

## 📅 The 4-Phase Master Timeline

```
[ Phase 1: Weeks 1-3 ] ──> [ Phase 2: Weeks 4-7 ] ──> [ Phase 3: Weeks 8-12 ] ──> [ Phase 4: Weeks 13-16 ]
  Foundations &              Linear Patterns            Trees, Heaps,              DP, Advanced Graphs
  Structures From Scratch                               Backtracking, Graphs       & Specialty Topics
```

| Phase | Duration | Primary Focus | Pace | Goal |
| :--- | :--- | :--- | :--- | :--- |
| **Phase 1** | Weeks 1-3 | Big-O, recursion, sorting, 10 structures | ~1 structure / day | Understand memory layout and cost of every operation. |
| **Phase 2** | Weeks 4-7 | Arrays, two pointers, windows, stacks, binary search, lists | 1-2 problems / day | Recognize linear patterns on sight. |
| **Phase 3** | Weeks 8-12 | Trees, tries, heaps, backtracking, graphs | 1-2 problems / day | Get comfortable with recursion and traversal. |
| **Phase 4** | Weeks 13-16 | DP, advanced graphs, greedy, intervals, bits, math | 1-2 problems / day | Handle multi-idea problems. |

**Realistic expectation:** 16 weeks gets you *competent* across all patterns, not "mastered." Mastery comes from the review loop (see Execution Rules) over the following months.

**Code layout in this repo:** structures go in `dsa/`, problem solutions go in `leetcode/`.

**Day-by-day schedule:** see [plan.md](plan.md).

---

## 🧠 Phase 0 → 1: Foundations (Week 1, before the structures)

These are prerequisites for everything else. Don't skip them.

1. **Big-O Analysis**
   * Time and space complexity for loops, nested loops, recursion, and amortized operations (e.g. why slice `append` is amortized O(1)).
   * Know the common classes cold: O(1), O(log N), O(N), O(N log N), O(N²), O(2ᴺ), O(N!).
2. **Recursion**
   * Base case, recursive case, call stack. Implement by hand: factorial, Fibonacci (naive → memoized), power of a number, reversing a string.
   * Draw the recursion tree for Fibonacci — this is the seed of both backtracking and DP.
3. **Sorting Algorithms**
   * Implement from scratch: **Insertion Sort**, **Merge Sort**, **Quick Sort** (with Lomuto or Hoare partition).
   * *Systems Insight:* Merge sort needs O(N) scratch space; quicksort is in-place but has an O(N²) worst case. Go's `slices.Sort` uses pdqsort, a quicksort hybrid that avoids that worst case.

---

## 🛠️ Phase 1: Data Structures From Scratch (Weeks 1-3)

Implement these **10 structures** using only raw Go primitives and pointers — no external libraries or `container/*` packages. Every structure gets a `_test.go` file (see Testing Rule below).

### 🧱 Contiguous vs. Pointer Memory Layouts
1. **Dynamic Array / Slice Buffer**
   * *Implementation:* A struct holding a fixed backing array, length, and capacity. Write `.Append()` that, when full, allocates a new array at double capacity and copies values over.
   * *Systems Insight:* This is what the Go runtime does on slice growth. The old array is reclaimed by the garbage collector.
2. **Singly Linked List**
   * *Implementation:* A node struct with data and a `*Node`. Implement `.InsertAtHead()`, `.DeleteValue()`, and `.Reverse()`.
   * *Systems Insight:* Nodes are scattered across the heap, so traversal causes cache misses — this is why arrays usually beat lists in practice even when Big-O says otherwise.
3. **Doubly Linked List**
   * *Implementation:* Add `prev` pointers. Track both `head` and `tail`. Implement insert/delete at both ends and delete-by-node.
   * *Systems Insight:* Extra pointer per node is metadata overhead. Every relink must update 4 pointers consistently or you get broken or dangling links. (Foundation for LRU Cache.)
4. **Custom Hash Map**
   * *Implementation:* An array of buckets plus a hash function (e.g. FNV-1a over the key's bytes). Handle collisions with **Chaining** or **Open Addressing**, and resize when the load factor passes a threshold.
   * *Systems Insight:* Lookups are O(1) on average because memory is traded for speed — the table is kept sparse on purpose.

### 🥞 Abstract Rule Layers & Trees
5. **Stack (LIFO)**
   * *Implementation:* Built on your dynamic array. Implement `.Push()`, `.Pop()`, and `.Peek()`.
   * *Systems Insight:* Models the thread call stack. Any recursive algorithm can be rewritten with an explicit stack.
6. **Queue (FIFO)**
   * *Implementation:* A **Ring Buffer** over a fixed array (head/tail indices wrapping with modulo). Optionally a second version on your doubly linked list.
   * *Systems Insight:* A fixed-size ring buffer allocates nothing after creation and stays cache-friendly. Used in event loops, channels, and network buffers.
7. **Binary Search Tree (BST)**
   * *Implementation:* Nodes with `left` and `right` pointers. Implement `.Insert()`, `.Search()`, `.Delete()` (all three cases: leaf, one child, two children), and in-order traversal.
   * *Systems Insight:* An unbalanced BST degrades to a linked list (O(N)). This is why real systems use balanced trees (red-black, B-trees).
8. **Binary Min/Max Heap**
   * *Implementation:* A flat Go slice with index algebra: `Left = 2i + 1`, `Right = 2i + 2`, `Parent = (i-1)/2`. Implement `.Push()` (sift up) and `.Pop()` (sift down), plus heapify from an existing slice in O(N).
   * *Systems Insight:* Stores a tree in contiguous memory with no pointers at all.
9. **Trie (Prefix Tree)**
   * *Implementation:* Nodes with `children [26]*Node` and an `end bool`. Implement `.Insert()`, `.Search()`, `.StartsWith()`.
   * *Systems Insight:* Fixed-size child arrays are fast but waste memory; a map per node saves memory but is slower. Classic trade-off.
10. **Union-Find (Disjoint Set Union)**
    * *Implementation:* A `parent []int` and `rank []int`. Implement `.Find()` with path compression and `.Union()` by rank.
    * *Systems Insight:* Near-O(1) operations using only two flat arrays. Essential for connectivity problems and Kruskal's MST.

---

## 🎯 Phases 2-4: The Pattern Problems (Weeks 4-16)

Problems are grouped following the NeetCode roadmap. Difficulty is marked (E)asy / (M)edium / (H)ard.
**First pass:** do all E and M. **Second pass:** come back for the H problems once the whole roadmap is done.

After finishing this list, work through the rest of the **NeetCode 150** for the extra reps.

---

### 🔄 Phase 2: Linear Patterns (Weeks 4-7)

#### 1. Arrays & Hashing (incl. Prefix Sums)
* **The Core Idea:** Trade memory for time — use a hash map/set or a precomputed array to avoid nested loops.
* **LeetCode targets:**
  - [ ] 217: Contains Duplicate (E)
  - [ ] 242: Valid Anagram (E)
  - [ ] 1: Two Sum (E)
  - [ ] 49: Group Anagrams (M)
  - [ ] 238: Product of Array Except Self (M)
  - [ ] 128: Longest Consecutive Sequence (M)
  - [ ] 303: Range Sum Query - Immutable (E) — prefix sums
  - [ ] 560: Subarray Sum Equals K (M) — prefix sums + hash map

#### 2. Two Pointers
* **The Core Idea:** Walk two indices inward or side-by-side over a (usually sorted) array to get linear runtime.
* **LeetCode targets:**
  - [ ] 125: Valid Palindrome (E)
  - [ ] 167: Two Sum II - Input Array Is Sorted (M)
  - [ ] 15: 3Sum (M)
  - [ ] 11: Container With Most Water (M)
  - [ ] 42: Trapping Rain Water (H)
* **Systems Insight:** Pure index arithmetic over contiguous memory — no extra allocation needed.

#### 3. Sliding Window
* **The Core Idea:** Maintain a running window over a sequence, growing the right edge and shrinking the left, instead of recomputing every subarray.
* **LeetCode targets:**
  - [ ] 121: Best Time to Buy and Sell Stock (E)
  - [ ] 3: Longest Substring Without Repeating Characters (M)
  - [ ] 424: Longest Repeating Character Replacement (M)
  - [ ] 567: Permutation in String (M)
  - [ ] 76: Minimum Window Substring (H)
  - [ ] 239: Sliding Window Maximum (H) — monotonic deque
* **Systems Insight:** Same idea as rate limiters and moving averages over streams: O(1) update per new element.

#### 4. Stack & Monotonic Stack
* **The Core Idea:** Use LIFO order for matching/nesting; keep the stack sorted (monotonic) to find the "next greater/smaller" element in O(N).
* **LeetCode targets:**
  - [ ] 20: Valid Parentheses (E)
  - [ ] 155: Min Stack (M)
  - [ ] 150: Evaluate Reverse Polish Notation (M)
  - [ ] 739: Daily Temperatures (M) — monotonic stack
  - [ ] 853: Car Fleet (M)
  - [ ] 84: Largest Rectangle in Histogram (H)
* **Systems Insight:** Each element is pushed and popped at most once, turning an O(N²) scan into O(N).

#### 5. Binary Search
* **The Core Idea:** Halve the search space each step — on sorted arrays, and also on a *range of possible answers*.
* **LeetCode targets:**
  - [ ] 704: Binary Search (E)
  - [ ] 74: Search a 2D Matrix (M)
  - [ ] 875: Koko Eating Bananas (M) — binary search on the answer
  - [ ] 153: Find Minimum in Rotated Sorted Array (M)
  - [ ] 33: Search in Rotated Sorted Array (M)
  - [ ] 981: Time Based Key-Value Store (M)
* **Systems Insight:** Get the loop invariant right (`lo <= hi` vs `lo < hi`, `mid = lo + (hi-lo)/2` to avoid overflow). Off-by-one bugs here are the most common.

#### 6. Linked List (incl. Fast & Slow Pointers, In-place Reversal)
* **The Core Idea:** Rewire pointers in place; use two pointers at different speeds to find cycles or midpoints.
* **LeetCode targets:**
  - [ ] 206: Reverse Linked List (E)
  - [ ] 21: Merge Two Sorted Lists (E)
  - [ ] 141: Linked List Cycle (E) — fast & slow
  - [ ] 19: Remove Nth Node From End of List (M)
  - [ ] 143: Reorder List (M) — find middle + reverse + merge
  - [ ] 92: Reverse Linked List II (M)
  - [ ] 2: Add Two Numbers (M)
  - [ ] 287: Find the Duplicate Number (M) — fast & slow on an array
  - [ ] 146: LRU Cache (M) — your doubly linked list + hash map
  - [ ] 23: Merge k Sorted Lists (H)
* **Systems Insight:** In Go, a wrong pointer gives a nil-pointer panic or a silently lost node. Use a dummy head node to remove edge cases. Floyd's cycle detection finds cycles in O(1) extra space.

---

### 🗺️ Phase 3: Trees, Heaps, Backtracking & Graphs (Weeks 8-12)

#### 7. Trees (DFS & BFS)
* **The Core Idea:** DFS via recursion (or an explicit stack) for depth/path questions; BFS via a queue for level-by-level questions.
* **LeetCode targets:**
  - [ ] 226: Invert Binary Tree (E)
  - [ ] 104: Maximum Depth of Binary Tree (E)
  - [ ] 100: Same Tree (E)
  - [ ] 543: Diameter of Binary Tree (E)
  - [ ] 110: Balanced Binary Tree (E)
  - [ ] 102: Binary Tree Level Order Traversal (M) — BFS
  - [ ] 199: Binary Tree Right Side View (M) — BFS
  - [ ] 235: Lowest Common Ancestor of a BST (M)
  - [ ] 98: Validate Binary Search Tree (M)
  - [ ] 230: Kth Smallest Element in a BST (M)
  - [ ] 105: Construct Binary Tree from Preorder and Inorder Traversal (M)
  - [ ] 124: Binary Tree Maximum Path Sum (H)
  - [ ] 297: Serialize and Deserialize Binary Tree (H)
* **Systems Insight:** Recursion depth equals tree height. A degenerate tree with 10⁵ nodes means 10⁵ stack frames, so practice the iterative (explicit stack/queue) versions too.

#### 8. Tries
* **The Core Idea:** Share common prefixes across many strings.
* **LeetCode targets:**
  - [ ] 208: Implement Trie (M)
  - [ ] 211: Design Add and Search Words Data Structure (M)
  - [ ] 212: Word Search II (H)

#### 9. Heap / Priority Queue (incl. Top K & Two Heaps)
* **The Core Idea:** Keep the min/max available in O(log N). A size-K heap finds the top K in O(N log K). Two heaps track a running median.
* **LeetCode targets:**
  - [ ] 703: Kth Largest Element in a Stream (E)
  - [ ] 1046: Last Stone Weight (E)
  - [ ] 347: Top K Frequent Elements (M) — also try bucket sort for O(N)
  - [ ] 973: K Closest Points to Origin (M)
  - [ ] 215: Kth Largest Element in an Array (M) — also try quickselect
  - [ ] 621: Task Scheduler (M)
  - [ ] 295: Find Median from Data Stream (H) — two heaps
* **Systems Insight:** A bounded heap uses O(K) memory no matter how large the stream is — the basis of top-K in log/metrics pipelines. In Go, learn `container/heap` too, since you'll use it in real code.

#### 10. Backtracking
* **The Core Idea:** Build candidates step by step, and undo the last choice when a branch is dead or finished.
* **LeetCode targets:**
  - [ ] 78: Subsets (M)
  - [ ] 90: Subsets II (M) — handling duplicates
  - [ ] 39: Combination Sum (M)
  - [ ] 46: Permutations (M)
  - [ ] 79: Word Search (M)
  - [ ] 131: Palindrome Partitioning (M)
  - [ ] 51: N-Queens (H)
* **The Go Rule:** Use one shared `path` slice: `append` before recursing, truncate (`path = path[:len(path)-1]`) after. **When you record a result, copy it** (`slices.Clone(path)`) — otherwise later appends overwrite the same backing array and corrupt your saved results. This is the #1 Go bug on these problems.

#### 11. Graphs (DFS, BFS, Topological Sort, Union-Find)
* **The Core Idea:** Model the problem as nodes + edges (adjacency list or grid), then traverse with a `visited` set.
* **LeetCode targets:**
  - [ ] 200: Number of Islands (M)
  - [ ] 695: Max Area of Island (M)
  - [ ] 133: Clone Graph (M)
  - [ ] 994: Rotting Oranges (M) — multi-source BFS
  - [ ] 417: Pacific Atlantic Water Flow (M)
  - [ ] 130: Surrounded Regions (M)
  - [ ] 207: Course Schedule (M) — cycle detection / topological sort
  - [ ] 210: Course Schedule II (M) — topological sort (Kahn's algorithm)
  - [ ] 684: Redundant Connection (M) — your Union-Find
  - [ ] 127: Word Ladder (H)
* **Systems Insight:** Topological sort is how build systems (Go modules, Make) order compilation and detect circular dependencies.

---

### 🧩 Phase 4: Dynamic Programming, Advanced Graphs & Specialty Topics (Weeks 13-16)

#### 12. Advanced Graphs
* **The Core Idea:** Weighted shortest paths and minimum spanning trees.
* **LeetCode targets:**
  - [ ] 743: Network Delay Time (M) — Dijkstra
  - [ ] 1584: Min Cost to Connect All Points (M) — Prim's / Kruskal's MST
  - [ ] 787: Cheapest Flights Within K Stops (M) — Bellman-Ford
  - [ ] 778: Swim in Rising Water (H)
  - [ ] 332: Reconstruct Itinerary (H)
  - (Optional, LeetCode Premium) 269: Alien Dictionary (H) — topological sort

#### 13. 1-D Dynamic Programming
* **The Core Idea:** Define `dp[i]` as the answer for a prefix/state, write the recurrence, then fill it bottom-up. Start from the memoized recursion you wrote in Phase 0.
* **LeetCode targets:**
  - [ ] 70: Climbing Stairs (E)
  - [ ] 746: Min Cost Climbing Stairs (E)
  - [ ] 198: House Robber (M)
  - [ ] 213: House Robber II (M)
  - [ ] 5: Longest Palindromic Substring (M)
  - [ ] 91: Decode Ways (M)
  - [ ] 322: Coin Change (M) — *unbounded* knapsack
  - [ ] 139: Word Break (M)
  - [ ] 300: Longest Increasing Subsequence (M)
  - [ ] 416: Partition Equal Subset Sum (M) — classic *0/1* knapsack
* **Systems Insight:** Most 1-D DP only needs the last one or two values, so O(N) memory can shrink to O(1).

#### 14. 2-D Dynamic Programming
* **The Core Idea:** State depends on two indices (two strings, a grid, or item × capacity).
* **LeetCode targets:**
  - [ ] 62: Unique Paths (M)
  - [ ] 1143: Longest Common Subsequence (M) — two-string DP
  - [ ] 518: Coin Change II (M) — unbounded knapsack, counting
  - [ ] 494: Target Sum (M) — 0/1 knapsack, counting
  - [ ] 309: Best Time to Buy and Sell Stock with Cooldown (M) — state machine DP
  - [ ] 72: Edit Distance (H)
* **Systems Insight:** When each row depends only on the previous row, the 2-D table shrinks to a single 1-D row — much less memory and far better cache behaviour.

#### 15. Greedy
* **The Core Idea:** Make the locally best choice at each step — and be able to argue *why* that's globally optimal.
* **LeetCode targets:**
  - [ ] 53: Maximum Subarray (M) — Kadane's algorithm
  - [ ] 55: Jump Game (M)
  - [ ] 45: Jump Game II (M)
  - [ ] 134: Gas Station (M)
  - [ ] 846: Hand of Straights (M)
  - [ ] 763: Partition Labels (M)

#### 16. Intervals
* **The Core Idea:** Sort by start (or end), then merge or count overlaps in one pass.
* **LeetCode targets:**
  - [ ] 56: Merge Intervals (M)
  - [ ] 57: Insert Interval (M)
  - [ ] 435: Non-overlapping Intervals (M)
  - [ ] 1851: Minimum Interval to Include Each Query (H)
* **Systems Insight:** Same logic as a memory allocator merging adjacent free blocks, or a calendar merging bookings.

#### 17. Bit Manipulation
* **The Core Idea:** Use XOR, AND, shifts, and masks to operate on individual bits.
* **LeetCode targets:**
  - [ ] 136: Single Number (E)
  - [ ] 191: Number of 1 Bits (E)
  - [ ] 338: Counting Bits (E)
  - [ ] 190: Reverse Bits (E)
  - [ ] 268: Missing Number (E)
  - [ ] 371: Sum of Two Integers (M)
* **Systems Insight:** Used for flags, compact state, and hashing. Know Go's `math/bits` package too.

#### 18. Math & Geometry
* **The Core Idea:** Matrix index manipulation and number tricks.
* **LeetCode targets:**
  - [ ] 202: Happy Number (E)
  - [ ] 48: Rotate Image (M)
  - [ ] 54: Spiral Matrix (M)
  - [ ] 73: Set Matrix Zeroes (M)
  - [ ] 50: Pow(x, n) (M)
  - [ ] 43: Multiply Strings (M)

---

## ⚡ The Senior Engineer Execution Rules

1. **The 20-Minute Boundary Rule**
   * Attempt a problem blind for **20 minutes** (up to 30-40 for Mediums once you're comfortable). If you're stuck, *stop*.
   * Read the solution pattern or watch a breakdown (e.g. NeetCode).
   * Close the solution, wait 10 minutes, then write the Go code from memory.
2. **Review Loop (Spaced Repetition)**
   * Re-solve every problem **1 week later** and again **1 month later**, from a blank file.
   * Keep a log so you know what's due:

     | Date | Problem | Pattern | Solved alone? | Re-solve 1w | Re-solve 1m |
     | :--- | :--- | :--- | :--- | :--- | :--- |
     | 2026-10-01 | 1: Two Sum | Hashing | ✅ | [ ] | [ ] |
   * Before coding, say the pattern out loud: *"This is a sliding window because…"*. Recognizing the pattern is the actual skill being trained.
3. **Testing Rule (Phase 1 structures)**
   * Every structure gets table-driven tests with `go test`.
   * Use Go fuzzing to compare against the standard library — it catches bugs you'd never think to test:
     ```go
     func FuzzHashMap(f *testing.F) {
         f.Fuzz(func(t *testing.T, keys []byte) {
             mine, ref := NewHashMap(), map[string]int{}
             for i, k := range keys {
                 key := string(k)
                 mine.Put(key, i)
                 ref[key] = i
             }
             for k, v := range ref {
                 if got, ok := mine.Get(k); !ok || got != v {
                     t.Fatalf("Get(%q) = %d, %v; want %d", k, got, ok, v)
                 }
             }
         })
     }
     ```
     Run with `go test -fuzz=FuzzHashMap ./dsa`. Do the same for your heap vs. `slices.Sort`, and your sorting algorithms vs. `slices.Sort`.
4. **The Zero-Allocation Pass — second pass only**
   * **First, solve the problem cleanly** with whatever allocations feel natural. Adding performance constraints while still learning a pattern slows learning down.
   * **Then**, as a separate exercise, rewrite it: pre-size slices with `make([]T, 0, n)`, avoid allocating inside loops, reuse buffers. Check your work with `go test -bench . -benchmem` and aim for `0 allocs/op` where possible.
