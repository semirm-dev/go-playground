# 🚀 The Master DSA Roadmap

This roadmap starts with **foundations built from scratch**, then works through the **NeetCode roadmap patterns** with real LeetCode problems, week by week, day by day.

> **Data structures and LeetCode are not separate stages.** The patterns *are* the algorithms. You don't master a structure by implementing it once — you master it by using it on problems you haven't seen. Build the foundations in weeks 1-2, then solve problems every day from week 3 onward.

**Code layout in this repo:** structures go in `dsa/`, problem solutions go in `leetcode/`.

---

## 📅 Timeline (4 hours/day, 6 days/week)

| Stage | Weeks | Focus | Pace |
| :--- | :--- | :--- | :--- |
| **Phase 1** | 1-2 | Big-O, recursion, sorting, 12 structures from scratch | 1-2 structures / day |
| **Phase 2** | 3-4 | Arrays, two pointers, windows, stacks, binary search, lists | 3 problems / day |
| **Phase 3** | 5-6 | Trees, tries, heaps, backtracking, graphs | 2-3 problems / day |
| **Phase 4** | 7-8 | Advanced graphs, DP, greedy, intervals, bits, math | 3-4 problems / day |
| **Hard Pass** | 9-10 | All deferred Hard problems | 1-2 problems / day |
| **Interview Prep** | 11-14 | Mixed-topic problems, mock interviews, contests | 2-3 problems / day |

**Weeks 1-8 are where you learn; weeks 9-14 are where you learn to perform in an interview.**

**Realistic expectation:** interview-ready in about **2.5-3.5 months**. Start applying around weeks 10-12. If a day runs over, let the schedule slip — never drop the review block. If Phase 1 feels rushed, give it a third week, but no more.

---

## ⏱️ How to Use Each Day

Every day below has three parts: **📖 Learn**, **🛠️ Do**, and **✅ Done when**. Work through them in this order:

| Block | Time | What |
| :--- | :--- | :--- |
| **1 — Learn** | 45 min | The day's 📖 Learn items. |
| **2 — Do** | 1.5 h | The day's 🛠️ Do items. Tick each one off. |
| **3 — Review** | 45 min | Re-solve from a blank file everything solved **7 days ago** and **30 days ago** (see the Review Log). In weeks 1-2, rebuild a structure from earlier in the week instead. |
| **4 — Perform** | 1 h | Check the ✅ Done when items out loud, as if explaining to an interviewer. From week 3: also re-solve one old problem out loud against a 25-min timer. |

A day is finished when every ✅ item is true. If it isn't, carry it into tomorrow. **Sunday: fully off** — burning out in week 5 costs more than a slower pace.

### 📚 Resources
- **[NeetCode.io](https://neetcode.io/roadmap)** — a video explanation for every problem here.
- **[VisuAlgo](https://visualgo.net)** — animations of every data structure and sorting algorithm.
- **Book: *Grokking Algorithms* (Aditya Bhargava, 1st ed.)** — the most beginner-friendly intro to Big-O, recursion, sorting, hashing, graphs, and DP.
- **Go:** [Go Slices: usage and internals](https://go.dev/blog/slices-intro), and the [fuzzing tutorial](https://go.dev/doc/tutorial/fuzz).

---

## ⚡ Rules

1. **The 20-Minute Rule**
   * Attempt a problem blind for **20 minutes** (up to 30-40 for Mediums once you're comfortable). If you're stuck, *stop*.
   * Read the solution pattern or watch the NeetCode breakdown.
   * Close the solution, wait 10 minutes, then write the Go code from memory.
   * Before coding, say the pattern out loud: *"This is a sliding window because…"*. Recognizing the pattern is the actual skill being trained.
2. **Review Log** — tells block 3 what's due:

   | Date | Problem | Pattern | Solved alone? | Re-solve 7d | Re-solve 30d |
   | :--- | :--- | :--- | :--- | :--- | :--- |
   | 2026-10-01 | 1: Two Sum | Hashing | ✅ | [ ] | [ ] |
3. **Test what you build** — every Phase 1 structure gets a `_test.go` file. When a day says *fuzz*, compare your version against Go's built-in one on random input (example on week 1 Fri).
4. **Hard problems (H) wait** for the Hard Pass (weeks 9-10). First pass is Easy and Medium only.
5. **Zero-Allocation Pass — later, not now.** Solve cleanly first. After week 14, rewrite your favourite solutions to pre-size slices (`make([]T, 0, n)`), avoid allocating inside loops, and check with `go test -bench . -benchmem`.

---

## 🧱 Phase 1: Foundations & Structures From Scratch (Weeks 1-2)

No LeetCode yet. Use only raw Go primitives and pointers — no external libraries or `container/*` packages. Every structure gets a `_test.go` file.

### Week 1 — Big-O, Recursion, Sorting, Linear Structures

#### Mon — Big-O & Recursion
📖 **Learn**
- *Grokking* ch. 1 (Big-O) and ch. 3 (recursion).
- Big-O measures how work grows with input size: O(1), O(log N), O(N), O(N log N), O(N²), O(2ᴺ), O(N!).
- Recursion = a base case + a recursive case; each call is a frame on the call stack.

🛠️ **Do**
- [ ] Write the Big-O (time and space) of 10 small Go snippets: single loop, nested loops, halving loop, recursion.
- [ ] Recursive factorial, power, reverse a string — with tests.

✅ **Done when**
- You can give one code example for each of O(1), O(log N), O(N), O(N log N), O(N²), O(2ᴺ).
- You can explain what happens on the call stack when `factorial(3)` runs.

#### Tue — Memoization & Simple Sorts
📖 **Learn**
- *Grokking* ch. 2 (selection sort, arrays vs linked lists).
- VisuAlgo "Sorting": watch insertion sort and merge sort.
- Memoization: cache results of recursive calls so each is computed once. The Fibonacci recursion tree is the seed of both backtracking and DP.

🛠️ **Do**
- [ ] Fibonacci naive → memoized. Draw the recursion tree for `fib(5)` on paper.
- [ ] Insertion Sort.
- [ ] Merge Sort.

✅ **Done when**
- You can explain why naive `fib` is O(2ᴺ) and memoized `fib` is O(N).
- You can explain why merge sort needs O(N) extra space.

#### Wed — Quicksort & Dynamic Array
📖 **Learn**
- *Grokking* ch. 4 (quicksort, divide and conquer).
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro).
- [Go fuzzing tutorial](https://go.dev/doc/tutorial/fuzz): Go feeds your test random inputs to find bugs you didn't think of.
- Quicksort is in-place but has an O(N²) worst case. Go's `slices.Sort` uses pdqsort, a quicksort hybrid that avoids it.
- A dynamic array is a backing array + length + capacity. When full, `Append` allocates a new array at double capacity and copies values over — exactly what Go does on slice growth.

🛠️ **Do**
- [ ] Quick Sort (Lomuto or Hoare partition).
- [ ] Fuzz all three sorts: sort a random slice with yours and with `slices.Sort`, and fail if they differ.
- [ ] Dynamic Array: `Append`, `Get`, `Set`, `Pop`, doubling growth.

✅ **Done when**
- Fuzzing runs for 1 minute with no failures.
- You can explain why `append` is amortized O(1) even though growth copies everything.

#### Thu — Linked Lists
📖 **Learn**
- VisuAlgo "Linked List" (singly and doubly).
- Linked list nodes are scattered across the heap, so traversal causes cache misses — why arrays usually beat lists in practice.
- In a doubly linked list, every relink must update 4 pointers consistently. (Foundation for LRU Cache.)

🛠️ **Do**
- [ ] Singly Linked List: `InsertAtHead`, `DeleteValue`, `Reverse`.
- [ ] Doubly Linked List: `head`/`tail`, insert and delete at both ends, delete a given node.

✅ **Done when**
- Tests cover empty, one-element, and many-element lists.
- You can reverse a singly linked list on paper, drawing each pointer change.

#### Fri — Hash Map
📖 **Learn**
- *Grokking* ch. 5 (hash tables); VisuAlgo "Hash Table".
- Buckets + a hash function (e.g. FNV-1a). Collisions via chaining or open addressing. Resize when the load factor (items / buckets) passes ~0.75.

🛠️ **Do**
- [ ] Hash Map: `Put`, `Get`, `Delete`, chaining, resizing.
- [ ] Fuzz against Go's `map`: put the same keys in both, then check they return the same values.
  Run with `go test -fuzz=FuzzHashMap ./dsa`.

✅ **Done when**
- Fuzzing passes.
- You can explain why lookups are O(1) on average but O(N) in the worst case.

#### Sat — Stack, Queue, Deque
📖 **Learn**
- VisuAlgo "Linked List" → Stack, Queue, and Deque tabs.
- A ring buffer is a fixed array with head/tail indices that wrap around with modulo. It allocates nothing after creation — used in event loops, channels, network buffers.
- Any recursive algorithm can be rewritten with an explicit stack.

🛠️ **Do**
- [ ] Stack on your dynamic array: `Push`, `Pop`, `Peek`.
- [ ] Queue as a ring buffer.
- [ ] Deque: extend the ring buffer to push/pop at both ends (needed later for 239 Sliding Window Maximum).

✅ **Done when**
- Tests include the head/tail wrapping past the end of the array.
- You can explain LIFO vs FIFO and give one real use of each.

### Week 2 — Trees, Heap, Trie, Union-Find, Graphs

#### Mon — Binary Search Tree
📖 **Learn**
- VisuAlgo "BST".
- Left subtree < node < right subtree, so in-order traversal gives sorted output. An unbalanced BST degrades to a linked list — why real systems use balanced trees (red-black, B-trees).

🛠️ **Do**
- [ ] BST: `Insert`, `Search`, in-order traversal.
- [ ] `Delete`: leaf, one child, two children.

✅ **Done when**
- Tests cover all three delete cases.
- You can explain why a BST built from sorted input is O(N) per operation.

#### Tue — Min Heap
📖 **Learn**
- VisuAlgo "Binary Heap".
- A tree stored in a flat slice, no pointers: `Left = 2i+1`, `Right = 2i+2`, `Parent = (i-1)/2`.

🛠️ **Do**
- [ ] Min Heap: `Push` (sift up), `Pop` (sift down), heapify in O(N).
- [ ] Fuzz: push N values, pop them all, the result must be sorted.

✅ **Done when**
- Fuzzing passes.
- You can draw the slice and the tree side by side for 7 values.

#### Wed — Trie & Union-Find
📖 **Learn**
- NeetCode's Trie video and Union-Find video.
- Trie: nodes with `children [26]*Node` and `end bool`. Fixed arrays are fast but waste memory; a map per node is the opposite trade-off.
- Union-Find: `parent []int` + `rank []int`. Near-O(1) operations; essential for connectivity problems and Kruskal's MST.

🛠️ **Do**
- [ ] Trie: `Insert`, `Search`, `StartsWith`.
- [ ] Union-Find: `Find` with path compression, `Union` by rank.

✅ **Done when**
- You can explain what path compression does to the tree.
- You can explain the difference between `Search` and `StartsWith` in one sentence.

#### Thu — Traversals & Graphs
📖 **Learn**
- *Grokking* ch. 6 (BFS); VisuAlgo "Graph Traversal".
- DFS goes deep first (stack or recursion); BFS goes level by level (queue) and finds shortest paths in unweighted graphs.
- Graphs are stored as an adjacency list (`[][]int` or `map[int][]int`). Grids are graphs too: neighbours via direction vectors `{{1,0},{-1,0},{0,1},{0,-1}}` plus bounds checks.

🛠️ **Do**
- [ ] Tree traversals, recursive and iterative: pre-, in-, post-order (your stack) and level order (your queue).
- [ ] Graph: build an adjacency list from an edge list, then BFS and DFS with a `visited` set.
- [ ] Same BFS and DFS on a 2-D grid.

✅ **Done when**
- Recursive and iterative versions produce the same output in tests.
- You can explain when you'd pick BFS over DFS.

#### Fri — Go Toolkit for LeetCode
📖 **Learn**
- pkg.go.dev docs for `slices`, `sort`, `container/heap`, `strings`, `math`.

🛠️ **Do**
- [ ] A tiny example of each: `slices.Sort`/`SortFunc`/`Reverse`, `sort.Search`, a `container/heap` min-heap of ints, `strings.Builder`, `math.MaxInt`/`math.MinInt`.

✅ **Done when**
- You can write the `container/heap` boilerplate (`Len`, `Less`, `Swap`, `Push`, `Pop`) from memory.

#### Sat — 🏁 Phase 1 Checkpoint
📖 **Learn**
- Re-read your own Phase 1 code and note the parts you found hardest.

🛠️ **Do**
- [ ] From blank files, in 90 min: Dynamic Array, Hash Map, Min Heap — all tests pass.

✅ **Done when**
- All three pass in time → start Phase 2 on Monday.
- Otherwise → spend up to one more week rebuilding the structures you struggled with, then retry.

---

## 🔄 Phase 2: Linear Patterns (Weeks 3-4)

LeetCode starts. Solutions go in `leetcode/`. On every problem day, **Done when** also includes: each problem passes on LeetCode and is in your Review Log.

### Week 3 — Arrays & Hashing, Two Pointers, Sliding Window, Stack

#### Mon — Hashing
📖 **Learn**
- NeetCode "Arrays & Hashing" intro.
- Trade memory for time: a hash map or set replaces a nested loop.

🛠️ **Do**
- [ ] 217 Contains Duplicate (E)
- [ ] 242 Valid Anagram (E)
- [ ] 1 Two Sum (E)

✅ **Done when**
- You can explain how a map turns Two Sum from O(N²) into O(N).

#### Tue — Hashing (Mediums)
📖 **Learn**
- Prefix/suffix passes: precompute left-to-right and right-to-left arrays, then combine.
- Compare yesterday's solutions with NeetCode's.

🛠️ **Do**
- [ ] 49 Group Anagrams (M)
- [ ] 238 Product of Array Except Self (M)
- [ ] 128 Longest Consecutive Sequence (M)

✅ **Done when**
- You can explain why 128 is O(N) despite the inner loop.

#### Wed — Prefix Sums & Two Pointers
📖 **Learn**
- Prefix sums: `sum(i..j) = pre[j+1] - pre[i]`, any range sum in O(1).
- NeetCode "Two Pointers" intro.

🛠️ **Do**
- [ ] 303 Range Sum Query - Immutable (E)
- [ ] 560 Subarray Sum Equals K (M)
- [ ] 125 Valid Palindrome (E)

✅ **Done when**
- You can explain why 560 needs a map of prefix-sum counts, not a sliding window (hint: negative numbers).

#### Thu — Two Pointers
📖 **Learn**
- Walk two indices inward over a sorted array — linear time, no extra memory.

🛠️ **Do**
- [ ] 167 Two Sum II (M)
- [ ] 15 3Sum (M)
- [ ] 11 Container With Most Water (M)

✅ **Done when**
- You can explain which pointer moves in 11 and why that never skips the answer.

#### Fri — Sliding Window
📖 **Learn**
- NeetCode "Sliding Window" intro.
- Grow the right edge, shrink the left, update in O(1) per element instead of recomputing every subarray. Same idea as rate limiters and moving averages.

🛠️ **Do**
- [ ] 121 Best Time to Buy and Sell Stock (E)
- [ ] 3 Longest Substring Without Repeating Characters (M)
- [ ] 424 Longest Repeating Character Replacement (M)

✅ **Done when**
- You can state the "shrink when…" condition for 3 and 424.

#### Sat — Window & Stack
📖 **Learn**
- NeetCode "Stack" intro: LIFO for matching and nesting.

🛠️ **Do**
- [ ] 567 Permutation in String (M)
- [ ] 20 Valid Parentheses (E)
- [ ] 155 Min Stack (M)

✅ **Done when**
- You can explain how 155 returns the minimum in O(1).

### Week 4 — Monotonic Stack, Binary Search, Linked List

#### Mon — Monotonic Stack
📖 **Learn**
- Keep the stack sorted to find the next greater/smaller element. Each element is pushed and popped once, so O(N²) becomes O(N).

🛠️ **Do**
- [ ] 150 Evaluate Reverse Polish Notation (M)
- [ ] 739 Daily Temperatures (M)
- [ ] 853 Car Fleet (M)

✅ **Done when**
- You can explain why 739 is O(N) and not O(N²).

#### Tue — Binary Search
📖 **Learn**
- NeetCode "Binary Search" intro.
- Loop invariants: `lo <= hi` vs `lo < hi`, and `mid = lo + (hi-lo)/2`. Off-by-one bugs are the most common.
- "Binary search on the answer": search a *range of possible answers*, not an array.

🛠️ **Do**
- [ ] 704 Binary Search (E)
- [ ] 74 Search a 2D Matrix (M)
- [ ] 875 Koko Eating Bananas (M)

✅ **Done when**
- You can write binary search from memory with no off-by-one bugs.
- You can explain what range is being searched in 875.

#### Wed — Binary Search (Rotated & Custom)
📖 **Learn**
- In a rotated sorted array, one half around `mid` is always sorted — check which one.
- Compare yesterday's solutions with NeetCode's.

🛠️ **Do**
- [ ] 153 Find Minimum in Rotated Sorted Array (M)
- [ ] 33 Search in Rotated Sorted Array (M)
- [ ] 981 Time Based Key-Value Store (M)

✅ **Done when**
- You can explain how you decide which half to discard in 33.

#### Thu — Linked Lists
📖 **Learn**
- NeetCode "Linked List" intro.
- The dummy-head-node trick removes most edge cases.
- Floyd's fast & slow pointers find cycles and midpoints in O(1) space.

🛠️ **Do**
- [ ] 206 Reverse Linked List (E)
- [ ] 21 Merge Two Sorted Lists (E)
- [ ] 141 Linked List Cycle (E)

✅ **Done when**
- You can reverse a linked list from memory in under 3 minutes.

#### Fri — Linked Lists (Mediums)
📖 **Learn**
- A wrong pointer in Go gives a nil-pointer panic or a silently lost node — draw the pointers before coding.

🛠️ **Do**
- [ ] 19 Remove Nth Node From End of List (M)
- [ ] 143 Reorder List (M)
- [ ] 92 Reverse Linked List II (M)

✅ **Done when**
- You can explain how 143 combines three things you already know (middle, reverse, merge).

#### Sat — Linked List Design + 🏁 Phase 2 Checkpoint
📖 **Learn**
- Re-read your Phase 1 doubly linked list — 146 uses it with a map.

🛠️ **Do**
- [ ] 2 Add Two Numbers (M)
- [ ] 287 Find the Duplicate Number (M)
- [ ] 146 LRU Cache (M)
- [ ] 🏁 Checkpoint (block 4): 2 unseen NeetCode 150 problems from weeks 3-4 topics, 45 min.

✅ **Done when**
- You solved at least one checkpoint problem. If not, spend 2 extra days on Phase 2 review before moving on.

---

## 🌳 Phase 3: Trees, Tries, Heaps, Backtracking, Graphs (Weeks 5-6)

### Week 5 — Trees, Tries, Heaps

#### Mon — Tree DFS
📖 **Learn**
- NeetCode "Trees" intro: DFS (pre/in/post-order) for depth/path questions, BFS (level order) for level questions.
- Recursion depth equals tree height — a degenerate tree with 10⁵ nodes means 10⁵ stack frames, so know the iterative versions too.

🛠️ **Do**
- [ ] 226 Invert Binary Tree (E)
- [ ] 104 Maximum Depth of Binary Tree (E)
- [ ] 100 Same Tree (E)

✅ **Done when**
- You solved 104 both recursively and with BFS.

#### Tue — Returning Values Up the Tree
📖 **Learn**
- Many tree problems compute something (like height) on the way up and update a global answer as a side effect.

🛠️ **Do**
- [ ] 543 Diameter of Binary Tree (E)
- [ ] 110 Balanced Binary Tree (E)
- [ ] 102 Binary Tree Level Order Traversal (M)

✅ **Done when**
- You can explain how 543 gets the height and the diameter in one pass.

#### Wed — BST Properties
📖 **Learn**
- In-order traversal of a BST is sorted. Validating a BST means passing `(min, max)` bounds down, not just comparing with children.

🛠️ **Do**
- [ ] 199 Binary Tree Right Side View (M)
- [ ] 235 Lowest Common Ancestor of a BST (M)
- [ ] 98 Validate Binary Search Tree (M)

✅ **Done when**
- You can draw a tree where "only compare with children" wrongly passes 98.

#### Thu — BST & Trie
📖 **Learn**
- Re-read your Phase 1 Trie.
- Preorder's first value is the root; its position in inorder splits the left and right subtrees.

🛠️ **Do**
- [ ] 230 Kth Smallest Element in a BST (M)
- [ ] 105 Construct Binary Tree from Preorder and Inorder (M)
- [ ] 208 Implement Trie (M)

✅ **Done when**
- You can explain how 105 splits the arrays at each step.

#### Fri — Trie & Heaps
📖 **Learn**
- Go's `container/heap` interface (from week 2 Fri).
- A size-K heap finds the top K in O(N log K) using O(K) memory no matter how big the input is.

🛠️ **Do**
- [ ] 211 Design Add and Search Words (M)
- [ ] 703 Kth Largest Element in a Stream (E)
- [ ] 1046 Last Stone Weight (E)

✅ **Done when**
- You can explain why a size-K *min*-heap gives the Kth *largest*.

#### Sat — Top K
📖 **Learn**
- Bucket sort: index buckets by frequency.
- Quickselect: quicksort's partition, but recurse into one side only — O(N) average.

🛠️ **Do**
- [ ] 347 Top K Frequent Elements (M) — heap, then bucket sort
- [ ] 973 K Closest Points to Origin (M)
- [ ] 215 Kth Largest Element in an Array (M) — heap, then quickselect

✅ **Done when**
- You can compare heap O(N log K) vs quickselect O(N) average and say when you'd use each.

### Week 6 — Backtracking, Graphs

Backtracking gets only 2 problems a day — it's where most people get stuck.

#### Mon — Backtracking Intro
📖 **Learn**
- NeetCode "Backtracking" intro. Draw the decision tree for Subsets on paper before coding.
- Build candidates step by step; undo the last choice when a branch is dead or done.
- **The Go Rule:** use one shared `path` slice — `append` before recursing, `path = path[:len(path)-1]` after. **When you record a result, copy it** (`slices.Clone(path)`), or later appends overwrite the same backing array and corrupt your saved results. This is the #1 Go bug on these problems.

🛠️ **Do**
- [ ] 621 Task Scheduler (M)
- [ ] 78 Subsets (M)

✅ **Done when**
- You can explain the Go Rule and what breaks without `slices.Clone`.

#### Tue — Duplicates & Reuse
📖 **Learn**
- Skip duplicates by sorting and `if i > start && nums[i] == nums[i-1] { continue }`.
- Reusing an element: recurse with `i`; using it once: recurse with `i+1`.

🛠️ **Do**
- [ ] 90 Subsets II (M)
- [ ] 39 Combination Sum (M)

✅ **Done when**
- You can explain why 39 recurses with `i` and 78 with `i+1`.

#### Wed — Permutations & Grid Search
📖 **Learn**
- Permutations track a `used []bool`. Grid search marks a cell visited, recurses, then unmarks it.

🛠️ **Do**
- [ ] 46 Permutations (M)
- [ ] 79 Word Search (M)

✅ **Done when**
- You can explain why 79 must unmark the cell after recursing.

#### Thu — Partitioning & Grid Graphs
📖 **Learn**
- NeetCode "Graphs" intro: adjacency list vs grid, `visited` set.
- Re-read your week 2 grid BFS/DFS.

🛠️ **Do**
- [ ] 131 Palindrome Partitioning (M)
- [ ] 200 Number of Islands (M)
- [ ] 695 Max Area of Island (M)

✅ **Done when**
- You can write a grid flood fill from memory.

#### Fri — Graph BFS
📖 **Learn**
- Multi-source BFS: put every starting cell in the queue at once.
- Cloning a graph: a map from old node → new node doubles as the `visited` set.

🛠️ **Do**
- [ ] 133 Clone Graph (M)
- [ ] 994 Rotting Oranges (M)
- [ ] 417 Pacific Atlantic Water Flow (M)

✅ **Done when**
- You can explain why 994 starts BFS from all rotten oranges at once.

#### Sat — Topological Sort + 🏁 Phase 3 Checkpoint
📖 **Learn**
- Kahn's algorithm: count in-degrees, queue every node with in-degree 0, remove edges as you go. If not every node is processed, there's a cycle. This is how build systems (Go modules, Make) order compilation.

🛠️ **Do**
- [ ] 130 Surrounded Regions (M)
- [ ] 207 Course Schedule (M)
- [ ] 210 Course Schedule II (M)
- [ ] 🏁 Checkpoint (block 4): 2 unseen NeetCode 150 problems from weeks 5-6 topics, 45 min.

✅ **Done when**
- You can explain how Kahn's algorithm detects a cycle.
- You solved at least one checkpoint problem. If not, spend 2 extra days on Phase 3 review.

---

## 🧩 Phase 4: Advanced Graphs, DP, Greedy & Specialty Topics (Weeks 7-8)

Bit manipulation problems are short, so they're mixed in alongside the heavier days.

### Week 7 — Advanced Graphs, 1-D DP, Bits

#### Mon — Union-Find & Dijkstra
📖 **Learn**
- *Grokking* ch. 7 (Dijkstra); VisuAlgo "SSSP".
- Dijkstra: always expand the closest unvisited node, using a min-heap.
- XOR: `x ^ x = 0` and `x ^ 0 = x`.

🛠️ **Do**
- [ ] 684 Redundant Connection (M) — your Union-Find
- [ ] 743 Network Delay Time (M) — Dijkstra
- [ ] 136 Single Number (E)

✅ **Done when**
- You can explain why Dijkstra fails with negative edge weights.

#### Tue — MST & Bellman-Ford
📖 **Learn**
- VisuAlgo "MST": Prim (grow from a node with a heap) and Kruskal (sort edges + Union-Find).
- Bellman-Ford: relax every edge N-1 times; works with negative weights.
- `n & (n-1)` clears the lowest set bit.

🛠️ **Do**
- [ ] 1584 Min Cost to Connect All Points (M)
- [ ] 787 Cheapest Flights Within K Stops (M)
- [ ] 191 Number of 1 Bits (E)

✅ **Done when**
- You can explain why 787 runs Bellman-Ford for exactly K+1 rounds.

#### Wed — 1-D DP
📖 **Learn**
- *Grokking* ch. 9 (DP).
- Method for every DP problem: **recursion → memoization → bottom-up table → reduce memory** (most 1-D DP only needs the last 1-2 values, so O(N) → O(1)).

🛠️ **Do**
- [ ] 70 Climbing Stairs (E)
- [ ] 746 Min Cost Climbing Stairs (E)
- [ ] 198 House Robber (M)
- [ ] 338 Counting Bits (E)

✅ **Done when**
- You took 198 through all four steps of the method.

#### Thu — DP on Strings
📖 **Learn**
- `dp[i]` = the answer for the first `i` characters. Palindromes: expand around each center.

🛠️ **Do**
- [ ] 213 House Robber II (M)
- [ ] 5 Longest Palindromic Substring (M)
- [ ] 91 Decode Ways (M)

✅ **Done when**
- You can explain how 213 reduces to running 198 twice.

#### Fri — Unbounded Knapsack
📖 **Learn**
- Knapsack types: *unbounded* (reuse items, e.g. Coin Change) vs *0/1* (each item once, e.g. Partition Equal Subset Sum). The difference is one loop's direction.

🛠️ **Do**
- [ ] 322 Coin Change (M)
- [ ] 139 Word Break (M)
- [ ] 268 Missing Number (E)

✅ **Done when**
- You wrote the recurrence for 322 in a comment before coding it.

#### Sat — LIS & 0/1 Knapsack
📖 **Learn**
- Go's `math/bits` package.
- Re-read Friday's knapsack notes.

🛠️ **Do**
- [ ] 300 Longest Increasing Subsequence (M)
- [ ] 416 Partition Equal Subset Sum (M)
- [ ] 190 Reverse Bits (E)

✅ **Done when**
- You can explain why 416 loops over capacity backwards.

### Week 8 — 2-D DP, Greedy, Intervals, Math

#### Mon — 2-D DP
📖 **Learn**
- NeetCode "2-D DP" intro.
- State depends on two indices (two strings, a grid, or item × capacity). When each row only needs the previous row, the table shrinks to one row.

🛠️ **Do**
- [ ] 62 Unique Paths (M)
- [ ] 1143 Longest Common Subsequence (M)
- [ ] 518 Coin Change II (M)

✅ **Done when**
- You can fill the 1143 table by hand for `"abc"` and `"ac"`.

#### Tue — DP with States
📖 **Learn**
- State-machine DP: track one value per state (e.g. holding / sold / resting) and define the transitions.

🛠️ **Do**
- [ ] 494 Target Sum (M)
- [ ] 309 Best Time to Buy and Sell Stock with Cooldown (M)
- [ ] 202 Happy Number (E)

✅ **Done when**
- You can draw the three states and their transitions for 309.

#### Wed — Greedy
📖 **Learn**
- NeetCode "Greedy" intro: take the locally best choice, and be able to argue *why* it's globally optimal.
- Kadane's algorithm: extend the current subarray or start fresh, whichever is bigger.

🛠️ **Do**
- [ ] 53 Maximum Subarray (M)
- [ ] 55 Jump Game (M)
- [ ] 45 Jump Game II (M)
- [ ] 371 Sum of Two Integers (M)

✅ **Done when**
- You can argue why the greedy choice in 55 is always safe.

#### Thu — Greedy (Mediums)
📖 **Learn**
- Compare yesterday's solutions with NeetCode's.

🛠️ **Do**
- [ ] 134 Gas Station (M)
- [ ] 846 Hand of Straights (M)
- [ ] 763 Partition Labels (M)

✅ **Done when**
- You can explain the key insight of 134 in one sentence.

#### Fri — Intervals
📖 **Learn**
- NeetCode "Intervals" intro: sort by start (or end), then merge or count overlaps in one pass.

🛠️ **Do**
- [ ] 56 Merge Intervals (M)
- [ ] 57 Insert Interval (M)
- [ ] 435 Non-overlapping Intervals (M)
- [ ] 43 Multiply Strings (M)

✅ **Done when**
- You can explain why 435 sorts by end, not start.

#### Sat — Matrix & Math + 🏁 Phase 4 Checkpoint
📖 **Learn**
- Matrix tricks: rotate = transpose + reverse each row; spiral = walk layer by layer with four boundaries.

🛠️ **Do**
- [ ] 48 Rotate Image (M)
- [ ] 54 Spiral Matrix (M)
- [ ] 73 Set Matrix Zeroes (M)
- [ ] 50 Pow(x, n) (M)
- [ ] 🏁 Checkpoint (block 4): 2 unseen NeetCode 150 problems from weeks 7-8 topics, 45 min.

✅ **Done when**
- You solved at least one checkpoint problem. If not, spend 2 extra days on Phase 4 review.

---

## 🔥 Hard Pass (Weeks 9-10)

Every Hard builds on a pattern you already know. **Learn** = re-solve that earlier problem quickly to warm up. **Do** = give the Hard **40 min** before looking at the solution. Leftover time: unseen NeetCode 150 problems, topics mixed. **Done when** = it passes, and you can explain the solution in 2 minutes.

### Week 9
| Day | 📖 Learn (warm-up) | 🛠️ Do |
| :--- | :--- | :--- |
| **Mon** | 11 Container With Most Water | [ ] 42 Trapping Rain Water |
| **Tue** | 424, 567 (sliding window) | [ ] 76 Minimum Window Substring |
| **Wed** | Your Phase 1 Deque + 739 | [ ] 239 Sliding Window Maximum |
| **Thu** | 739, 853 (monotonic stack) | [ ] 84 Largest Rectangle in Histogram |
| **Fri** | 21 Merge Two Sorted Lists + a heap | [ ] 23 Merge k Sorted Lists |
| **Sat** | 543 Diameter, 102 Level Order | [ ] 124 Binary Tree Maximum Path Sum · [ ] 297 Serialize and Deserialize Binary Tree |

### Week 10
| Day | 📖 Learn (warm-up) | 🛠️ Do |
| :--- | :--- | :--- |
| **Mon** | 208 Trie, 79 Word Search | [ ] 212 Word Search II |
| **Tue** | 703 Kth Largest in a Stream | [ ] 295 Find Median from Data Stream (two heaps) |
| **Wed** | 46 Permutations | [ ] 51 N-Queens |
| **Thu** | 994 Rotting Oranges (BFS) | [ ] 127 Word Ladder |
| **Fri** | 743 Network Delay Time (Dijkstra) | [ ] 778 Swim in Rising Water · [ ] 332 Reconstruct Itinerary |
| **Sat** | 1143 LCS, 56 Merge Intervals | [ ] 72 Edit Distance · [ ] 1851 Minimum Interval to Include Each Query |

(Optional, LeetCode Premium: 269 Alien Dictionary.)

**Start applying around weeks 10-12.** Early interviews are practice too.

---

## 🎤 Interview Prep (Weeks 11-14)

Every day looks the same, so there's one plan for all of them:

📖 **Learn (45 min)**
- **System design** (for mid/senior roles): *System Design Interview* vol. 1 (Alex Xu), one chapter every 2 days.
- Once, early on: write 5-6 **behavioural stories** in STAR format (Situation, Task, Action, Result).

🛠️ **Do (1.5 h)**
- 2-3 unseen NeetCode 150 problems, **topics mixed** — pick at random so you don't know the pattern in advance. When the 150 runs out, move to the NeetCode 250.

🔁 **Review (45 min)** — same as always: 7-day and 30-day re-solves.

🎤 **Perform (1 h)**
- **Mock interviews** 2× a week (a friend, or free peer platforms like Pramp).
- Other days: one unseen Medium out loud, 25-min timer.

Plus the **LeetCode Weekly Contest** every week (it runs on Sunday — swap that week's rest day).

✅ **Ready when:** you solve ~70% of unseen Mediums in ≤ 30 min, name the pattern in the first 5 min, and consistently solve 2 of 4 contest problems.

### After Week 14
- Keep Do and Review going at 1-2 h/day while interviewing.
- Do the Zero-Allocation Pass (Rule 5) on your 20 favourite solutions.
- Before each company's interview: their company-tagged LeetCode problems (Premium) or public interview reports.
