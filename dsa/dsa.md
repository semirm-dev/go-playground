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

**Weeks 1-8 are where you learn; weeks 9-14 are where you learn to perform in an interview.** Put the most effort into 1-8, but don't skip the rest.

**Realistic expectation:** interview-ready in about **2.5-3.5 months**. Start applying around weeks 10-12. If a day runs over, let the schedule slip — never drop the review block.

---

## ⏱️ Daily Rhythm

**Mon–Sat:** four blocks, with real breaks between them. **Sunday: fully off** — burning out in week 5 costs more than a slower pace.

| Block | Time | What |
| :--- | :--- | :--- |
| **1 — Solve** | 1.5 h | The day's new problems (structures in weeks 1-2). Freshest focus goes here. |
| **2 — Review** | 45 min | Re-solve from a blank file everything solved **7 days ago** and **30 days ago**. If you remember the code and finish in 3 minutes, do a fresh problem of the same pattern instead. |
| **3 — Learn** | 45 min | Study **tomorrow's** topic (the "Learn" line of each week), then read other people's solutions to today's problems and compare. |
| **4 — Perform** | 1 h | Solve one already-done problem **out loud** against a 25-min timer. On the last Saturday of each phase, this block is the **🏁 Checkpoint**. |

### 📚 Resources
- **[NeetCode.io](https://neetcode.io/roadmap)** — a video explanation for every problem here.
- **[VisuAlgo](https://visualgo.net)** — animations of every data structure and sorting algorithm.
- **Book: *Grokking Algorithms* (Aditya Bhargava)** — the most beginner-friendly intro to Big-O, recursion, sorting, hashing, graphs, and DP.
- **Go:** [Go Slices: usage and internals](https://go.dev/blog/slices-intro), and the [fuzzing tutorial](https://go.dev/doc/tutorial/fuzz).

---

## ⚡ Rules

1. **The 20-Minute Rule**
   * Attempt a problem blind for **20 minutes** (up to 30-40 for Mediums once you're comfortable). If you're stuck, *stop*.
   * Read the solution pattern or watch the NeetCode breakdown.
   * Close the solution, wait 10 minutes, then write the Go code from memory.
   * Before coding, say the pattern out loud: *"This is a sliding window because…"*. Recognizing the pattern is the actual skill being trained.
2. **Review Log** — tells block 2 what's due:

   | Date | Problem | Pattern | Solved alone? | Re-solve 7d | Re-solve 30d |
   | :--- | :--- | :--- | :--- | :--- | :--- |
   | 2026-10-01 | 1: Two Sum | Hashing | ✅ | [ ] | [ ] |
3. **Testing Rule (Phase 1 structures)** — every structure gets table-driven tests with `go test`, plus a fuzz test against the standard library:
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
   Run with `go test -fuzz=FuzzHashMap ./dsa`. Do the same for your sorts and heap against `slices.Sort`.
4. **Hard problems (H) wait** for the Hard Pass (weeks 9-10). First pass is Easy and Medium only.
5. **Zero-Allocation Pass — later, not now.** Solve cleanly first. After week 14, rewrite your favourite solutions to pre-size slices (`make([]T, 0, n)`), avoid allocating inside loops, and check with `go test -bench . -benchmem`.

---

## 🧱 Phase 1: Foundations & Structures From Scratch (Weeks 1-2)

No LeetCode yet. Use only raw Go primitives and pointers — no external libraries or `container/*` packages. Every structure gets a `_test.go` file. Block 4 this phase: explain out loud how the structure works and the Big-O of each operation.

### Week 1 — Big-O, Recursion, Sorting, Linear Structures
**Learn:** *Grokking Algorithms* ch. 1-5; VisuAlgo "Sorting", "Linked List", "Hash Table".

* **Big-O:** time and space for loops, nested loops, recursion, and amortized operations (why slice `append` is amortized O(1)). Know O(1), O(log N), O(N), O(N log N), O(N²), O(2ᴺ), O(N!) cold.
* **Recursion:** base case, recursive case, call stack. The Fibonacci recursion tree is the seed of both backtracking and DP.
* **Sorting:** merge sort needs O(N) scratch space; quicksort is in-place but has an O(N²) worst case. Go's `slices.Sort` uses pdqsort, a quicksort hybrid that avoids it.
* **Dynamic Array:** a backing array + length + capacity. `.Append()` allocates a new array at double capacity when full and copies values over — exactly what Go does on slice growth.
* **Singly Linked List:** `.InsertAtHead()`, `.DeleteValue()`, `.Reverse()`. Nodes are scattered across the heap, so traversal causes cache misses — why arrays usually beat lists in practice.
* **Doubly Linked List:** `prev` pointers, `head` and `tail`, insert/delete at both ends, delete-by-node. Every relink must update 4 pointers consistently. (Foundation for LRU Cache.)
* **Hash Map:** buckets + a hash function (e.g. FNV-1a), collisions via chaining or open addressing, resize past a load factor. O(1) on average because the table is kept sparse on purpose.
* **Stack:** on your dynamic array — `.Push()`, `.Pop()`, `.Peek()`. Any recursive algorithm can be rewritten with an explicit stack.
* **Queue:** a ring buffer over a fixed array (head/tail wrapping with modulo). Allocates nothing after creation; used in event loops, channels, network buffers.
* **Deque:** extend the ring buffer to push/pop at both ends. Needed for monotonic-deque problems (239 Sliding Window Maximum) and 0-1 BFS.

- [ ] **Mon:** Big-O — write down the complexity of 10 small Go snippets. Recursion: factorial, power, reverse a string.
- [ ] **Tue:** Fibonacci naive → memoized; draw the recursion tree for `fib(5)`. Insertion Sort + Merge Sort.
- [ ] **Wed:** Quick Sort (Lomuto or Hoare partition); fuzz all three sorts against `slices.Sort`. Dynamic Array.
- [ ] **Thu:** Singly Linked List + Doubly Linked List.
- [ ] **Fri:** Hash Map: Put/Get/Delete, collisions, resizing. Fuzz against Go's `map`.
- [ ] **Sat:** Stack + Queue (ring buffer) + Deque.

### Week 2 — Trees, Heap, Trie, Union-Find, Graphs
**Learn:** VisuAlgo "BST" and "Binary Heap"; NeetCode's Trie and Union-Find videos.

* **Binary Search Tree:** `.Insert()`, `.Search()`, `.Delete()` (leaf, one child, two children), in-order traversal. An unbalanced BST degrades to a linked list — why real systems use balanced trees (red-black, B-trees).
* **Min Heap:** a flat slice with `Left = 2i+1`, `Right = 2i+2`, `Parent = (i-1)/2`. `.Push()` (sift up), `.Pop()` (sift down), heapify in O(N). A tree stored in contiguous memory with no pointers.
* **Trie:** nodes with `children [26]*Node` and `end bool`; `.Insert()`, `.Search()`, `.StartsWith()`. Fixed arrays are fast but waste memory; a map per node is the opposite trade-off.
* **Union-Find:** `parent []int` + `rank []int`; `.Find()` with path compression, `.Union()` by rank. Near-O(1) operations; essential for connectivity problems and Kruskal's MST.
* **Graph:** adjacency list (`[][]int` or `map[int][]int`) built from an edge list, plus BFS and DFS with a `visited` set. Grids are graphs too: traverse with direction vectors `{{1,0},{-1,0},{0,1},{0,-1}}` and bounds checks.

- [ ] **Mon:** BST: Insert, Search, in-order traversal, Delete.
- [ ] **Tue:** Min Heap: Push/Pop/heapify. Fuzz: push N values, pop all, result must be sorted.
- [ ] **Wed:** Trie + Union-Find.
- [ ] **Thu:** Tree traversals both ways: recursive and iterative (your stack for DFS, your queue for BFS). Graph: adjacency list + BFS/DFS, and the same on a 2-D grid.
- [ ] **Fri:** Go tools you'll use on LeetCode: `slices`, `sort`, `container/heap`, `strings.Builder`, `math.MaxInt`. Write a tiny example of each.
- [ ] **Sat:** 🏁 **Checkpoint:** from blank files, in 90 min: Dynamic Array, Hash Map, Heap — all tests pass.

---

## 🔄 Phase 2: Linear Patterns (Weeks 3-4)

LeetCode starts. Solutions go in `leetcode/`.

### Week 3 — Arrays & Hashing, Two Pointers, Sliding Window, Stack
**Learn:** NeetCode intros for each topic; prefix sums.

* **Arrays & Hashing:** trade memory for time — a hash map/set or a precomputed (prefix sum) array replaces nested loops.
* **Two Pointers:** walk two indices inward or side-by-side over a (usually sorted) array for linear time, with no extra allocation.
* **Sliding Window:** grow the right edge, shrink the left, update in O(1) per element instead of recomputing every subarray. Same idea as rate limiters and moving averages.
* **Stack:** LIFO for matching and nesting.

- [ ] **Mon:** 217 Contains Duplicate (E) · 242 Valid Anagram (E) · 1 Two Sum (E)
- [ ] **Tue:** 49 Group Anagrams (M) · 238 Product of Array Except Self (M) · 128 Longest Consecutive Sequence (M)
- [ ] **Wed:** 303 Range Sum Query - Immutable (E) · 560 Subarray Sum Equals K (M) · 125 Valid Palindrome (E)
- [ ] **Thu:** 167 Two Sum II (M) · 15 3Sum (M) · 11 Container With Most Water (M)
- [ ] **Fri:** 121 Best Time to Buy and Sell Stock (E) · 3 Longest Substring Without Repeating Characters (M) · 424 Longest Repeating Character Replacement (M)
- [ ] **Sat:** 567 Permutation in String (M) · 20 Valid Parentheses (E) · 155 Min Stack (M)

### Week 4 — Monotonic Stack, Binary Search, Linked List
**Learn:** binary search loop invariants and "binary search on the answer"; Floyd's cycle detection; the dummy-head-node trick.

* **Monotonic Stack:** keep the stack sorted to find the next greater/smaller element. Each element is pushed and popped once — O(N²) becomes O(N).
* **Binary Search:** halve the search space — on sorted arrays, and on a *range of possible answers*. Get the invariant right (`lo <= hi` vs `lo < hi`, `mid = lo + (hi-lo)/2`); off-by-one bugs are the most common.
* **Linked List:** rewire pointers in place; fast & slow pointers find cycles and midpoints in O(1) space. A wrong pointer in Go gives a nil-pointer panic or a silently lost node — a dummy head removes most edge cases.

- [ ] **Mon:** 150 Evaluate Reverse Polish Notation (M) · 739 Daily Temperatures (M) · 853 Car Fleet (M)
- [ ] **Tue:** 704 Binary Search (E) · 74 Search a 2D Matrix (M) · 875 Koko Eating Bananas (M)
- [ ] **Wed:** 153 Find Minimum in Rotated Sorted Array (M) · 33 Search in Rotated Sorted Array (M) · 981 Time Based Key-Value Store (M)
- [ ] **Thu:** 206 Reverse Linked List (E) · 21 Merge Two Sorted Lists (E) · 141 Linked List Cycle (E)
- [ ] **Fri:** 19 Remove Nth Node From End of List (M) · 143 Reorder List (M) · 92 Reverse Linked List II (M)
- [ ] **Sat:** 2 Add Two Numbers (M) · 287 Find the Duplicate Number (M) · 146 LRU Cache (M — your doubly linked list + a map)
- [ ] 🏁 **Checkpoint** (Sat, block 4): 2 unseen NeetCode 150 problems from Phase 2 topics, 45 min. If you solve neither, spend 2 extra days on Phase 2 review.

---

## 🌳 Phase 3: Trees, Tries, Heaps, Backtracking, Graphs (Weeks 5-6)

### Week 5 — Trees, Tries, Heaps
**Learn:** NeetCode "Trees" intro (DFS pre/in/post-order vs BFS level order); Go's `container/heap` interface.

* **Trees:** DFS (recursion or explicit stack) for depth/path questions; BFS (queue) for level questions. Recursion depth equals tree height — a degenerate tree with 10⁵ nodes means 10⁵ stack frames, so know the iterative versions too.
* **Tries:** share common prefixes across many strings.
* **Heaps:** min/max in O(log N). A size-K heap finds the top K in O(N log K) using O(K) memory no matter how big the input is.

- [ ] **Mon:** 226 Invert Binary Tree (E) · 104 Maximum Depth of Binary Tree (E) · 100 Same Tree (E)
- [ ] **Tue:** 543 Diameter of Binary Tree (E) · 110 Balanced Binary Tree (E) · 102 Binary Tree Level Order Traversal (M)
- [ ] **Wed:** 199 Binary Tree Right Side View (M) · 235 Lowest Common Ancestor of a BST (M) · 98 Validate Binary Search Tree (M)
- [ ] **Thu:** 230 Kth Smallest Element in a BST (M) · 105 Construct Binary Tree from Preorder and Inorder (M) · 208 Implement Trie (M)
- [ ] **Fri:** 211 Design Add and Search Words (M) · 703 Kth Largest Element in a Stream (E) · 1046 Last Stone Weight (E)
- [ ] **Sat:** 347 Top K Frequent Elements (M — also try bucket sort) · 973 K Closest Points to Origin (M) · 215 Kth Largest Element in an Array (M — also try quickselect)

### Week 6 — Backtracking, Graphs
**Learn:** draw the decision tree for Subsets on paper before coding. Then NeetCode "Graphs" intro: adjacency list vs grid, `visited` set, Kahn's algorithm.

* **Backtracking:** build candidates step by step and undo the last choice when a branch is dead or done. **The Go Rule:** use one shared `path` slice — `append` before recursing, `path = path[:len(path)-1]` after. **When you record a result, copy it** (`slices.Clone(path)`), or later appends overwrite the same backing array and corrupt your saved results. This is the #1 Go bug on these problems.
* **Graphs:** model as nodes + edges (adjacency list or grid), traverse with a `visited` set. Topological sort orders dependencies and detects cycles — how build systems (Go modules, Make) order compilation.

Backtracking gets only 2 problems a day — it's where most people get stuck.

- [ ] **Mon:** 621 Task Scheduler (M) · 78 Subsets (M)
- [ ] **Tue:** 90 Subsets II (M) · 39 Combination Sum (M)
- [ ] **Wed:** 46 Permutations (M) · 79 Word Search (M)
- [ ] **Thu:** 131 Palindrome Partitioning (M) · 200 Number of Islands (M) · 695 Max Area of Island (M)
- [ ] **Fri:** 133 Clone Graph (M) · 994 Rotting Oranges (M — multi-source BFS) · 417 Pacific Atlantic Water Flow (M)
- [ ] **Sat:** 130 Surrounded Regions (M) · 207 Course Schedule (M) · 210 Course Schedule II (M — Kahn's algorithm)
- [ ] 🏁 **Checkpoint** (Sat, block 4): 2 unseen NeetCode 150 problems from Phase 3 topics, 45 min.

---

## 🧩 Phase 4: Advanced Graphs, DP, Greedy & Specialty Topics (Weeks 7-8)

Bit manipulation problems are short, so they're mixed in alongside the heavier DP days.

### Week 7 — Advanced Graphs, 1-D DP, Bits
**Learn:** Dijkstra (with your heap), Prim's MST, Bellman-Ford (VisuAlgo "SSSP" / "MST"); *Grokking* ch. 9; XOR and `n & (n-1)` tricks.

* **Advanced Graphs:** weighted shortest paths (Dijkstra, Bellman-Ford) and minimum spanning trees (Prim, Kruskal with your Union-Find).
* **1-D DP:** define `dp[i]`, write the recurrence, fill bottom-up. Method for every DP problem: **recursion → memoization → bottom-up table → reduce memory** (most 1-D DP only needs the last 1-2 values, so O(N) → O(1)).
* **Knapsack types:** Coin Change is *unbounded* (reuse items); Partition Equal Subset Sum is *0/1* (each item once). The difference is one loop's direction.
* **Bits:** XOR, AND, shifts, masks. Know Go's `math/bits`.

- [ ] **Mon:** 684 Redundant Connection (M — your Union-Find) · 743 Network Delay Time (M — Dijkstra) · 136 Single Number (E)
- [ ] **Tue:** 1584 Min Cost to Connect All Points (M — MST) · 787 Cheapest Flights Within K Stops (M — Bellman-Ford) · 191 Number of 1 Bits (E)
- [ ] **Wed:** 70 Climbing Stairs (E) · 746 Min Cost Climbing Stairs (E) · 198 House Robber (M) · 338 Counting Bits (E)
- [ ] **Thu:** 213 House Robber II (M) · 5 Longest Palindromic Substring (M) · 91 Decode Ways (M)
- [ ] **Fri:** 322 Coin Change (M — unbounded knapsack) · 139 Word Break (M) · 268 Missing Number (E)
- [ ] **Sat:** 300 Longest Increasing Subsequence (M) · 416 Partition Equal Subset Sum (M — 0/1 knapsack) · 190 Reverse Bits (E)

### Week 8 — 2-D DP, Greedy, Intervals, Math
**Learn:** Kadane's algorithm; NeetCode "2-D DP", "Greedy" and "Intervals" intros.

* **2-D DP:** state depends on two indices (two strings, a grid, or item × capacity). When each row only needs the previous row, the table shrinks to one row.
* **Greedy:** take the locally best choice — and be able to argue *why* it's globally optimal.
* **Intervals:** sort by start (or end), then merge or count overlaps in one pass.
* **Math & Geometry:** matrix index manipulation and number tricks.

- [ ] **Mon:** 62 Unique Paths (M) · 1143 Longest Common Subsequence (M) · 518 Coin Change II (M)
- [ ] **Tue:** 494 Target Sum (M) · 309 Best Time to Buy and Sell Stock with Cooldown (M) · 202 Happy Number (E)
- [ ] **Wed:** 53 Maximum Subarray (M — Kadane) · 55 Jump Game (M) · 45 Jump Game II (M) · 371 Sum of Two Integers (M)
- [ ] **Thu:** 134 Gas Station (M) · 846 Hand of Straights (M) · 763 Partition Labels (M)
- [ ] **Fri:** 56 Merge Intervals (M) · 57 Insert Interval (M) · 435 Non-overlapping Intervals (M) · 43 Multiply Strings (M)
- [ ] **Sat:** 48 Rotate Image (M) · 54 Spiral Matrix (M) · 73 Set Matrix Zeroes (M) · 50 Pow(x, n) (M)
- [ ] 🏁 **Checkpoint** (Sat, block 4): 2 unseen NeetCode 150 problems from Phase 4 topics, 45 min.

---

## 🔥 Hard Pass (Weeks 9-10)

Give each Hard **40 min** before looking at the solution. Leftover block-1 time: unseen NeetCode 150 problems, topics mixed.

### Week 9
- [ ] **Mon:** 42 Trapping Rain Water
- [ ] **Tue:** 76 Minimum Window Substring
- [ ] **Wed:** 239 Sliding Window Maximum (monotonic deque)
- [ ] **Thu:** 84 Largest Rectangle in Histogram
- [ ] **Fri:** 23 Merge k Sorted Lists
- [ ] **Sat:** 124 Binary Tree Maximum Path Sum · 297 Serialize and Deserialize Binary Tree

### Week 10
- [ ] **Mon:** 212 Word Search II
- [ ] **Tue:** 295 Find Median from Data Stream (two heaps)
- [ ] **Wed:** 51 N-Queens
- [ ] **Thu:** 127 Word Ladder
- [ ] **Fri:** 778 Swim in Rising Water · 332 Reconstruct Itinerary
- [ ] **Sat:** 72 Edit Distance · 1851 Minimum Interval to Include Each Query

(Optional, LeetCode Premium: 269 Alien Dictionary.)

**Start applying around weeks 10-12.** Early interviews are practice too.

---

## 🎤 Interview Prep (Weeks 11-14)

The blocks change focus:

| Block | Time | What |
| :--- | :--- | :--- |
| **1 — Solve** | 1.5 h | 2-3 unseen NeetCode 150 problems, **topics mixed** — pick at random so you don't know the pattern in advance. When the 150 runs out, move to the NeetCode 250. |
| **2 — Review** | 45 min | Same as before: 7-day and 30-day re-solves. |
| **3 — Learn** | 45 min | **System design** (for mid/senior roles): *System Design Interview* vol. 1 (Alex Xu), a chapter every 2 days. Write 5-6 **behavioural stories** in STAR format (Situation, Task, Action, Result). |
| **4 — Perform** | 1 h | **Mock interviews** 2× a week (a friend, or free peer platforms like Pramp). Other days: one unseen Medium out loud, 25-min timer. |

Plus the **LeetCode Weekly Contest** every week (it runs on Sunday — swap that week's rest day).

**Ready when:** you solve ~70% of unseen Mediums in ≤ 30 min, name the pattern in the first 5 min, and consistently solve 2 of 4 contest problems.

### After Week 14
- Keep blocks 1-2 going at 1-2 h/day while interviewing.
- Do the Zero-Allocation Pass (Rule 5) on your 20 favourite solutions.
- Before each company's interview: their company-tagged LeetCode problems (Premium) or public interview reports.
