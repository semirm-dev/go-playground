# 📆 Day-by-Day Study Plan

Companion to [dsa.md](dsa.md). That file says **what** to learn; this one says **when**.

## ⏱️ Weekly Rhythm

Assumes **~1.5 hours/day, Mon–Fri**, plus a shorter Saturday. If a week runs over, let the schedule slip — never skip the review day.

| Day | What |
| :--- | :--- |
| **Mon–Fri** | New material (see the week's table) |
| **Saturday** | Review: re-solve last week's problems from a blank file (the 1-week re-solve). Catch up on anything unfinished. |
| **Sunday** | Rest. From week 8 on: 2-3 of the 1-month re-solves that are due. |

### 🔁 Daily Session (~90 min)
1. **10 min — Warm-up:** re-solve one old problem from memory.
2. **15 min — Learn:** before the first problem of a new topic, watch/read its explanation (see "Learn first" in each week).
3. **55 min — Solve:** the day's problems, using the 20-minute rule from [dsa.md](dsa.md).
4. **10 min — Log:** add the problem to your review log, and write one sentence on *why* this pattern applies.

### 📚 Resources
- **[NeetCode.io](https://neetcode.io/roadmap)** — a video explanation for every problem in this plan.
- **[VisuAlgo](https://visualgo.net)** — animations of every data structure and sorting algorithm. Watch these before implementing.
- **Book: *Grokking Algorithms* (Aditya Bhargava)** — the most beginner-friendly intro to Big-O, recursion, sorting, hashing, graphs, and DP.
- **Go:** [Go Slices: usage and internals](https://go.dev/blog/slices-intro), and the [fuzzing tutorial](https://go.dev/doc/tutorial/fuzz).

---

## 🧱 Phase 1: Foundations & Structures (Weeks 1-3)

No LeetCode yet. Everything goes in `dsa/` with a `_test.go` file next to it.

### Week 1 — Big-O, Recursion, Sorting
**Learn first:** *Grokking Algorithms* ch. 1-4; VisuAlgo "Sorting".

| Day | Task |
| :--- | :--- |
| Mon | Big-O: read, then write down the complexity of 10 small Go snippets (loops, nested loops, halving loops). |
| Tue | Big-O for recursion + amortized `append`. Start recursion: factorial, power, reverse a string. |
| Wed | Fibonacci: naive → memoized. Draw the recursion tree by hand for `fib(5)`. |
| Thu | Insertion Sort + Merge Sort, with table-driven tests. |
| Fri | Quick Sort. Fuzz test all three sorts against `slices.Sort`. |
| Sat | Review: re-implement Merge Sort from a blank file. |

### Week 2 — Arrays, Lists, Hash Map
**Learn first:** VisuAlgo "Linked List" and "Hash Table"; *Grokking* ch. 5.

| Day | Task |
| :--- | :--- |
| Mon | Dynamic Array with manual growth. Test that capacity doubles. |
| Tue | Singly Linked List: InsertAtHead, DeleteValue, Reverse. |
| Wed | Doubly Linked List: head/tail, insert/delete both ends, delete-by-node. |
| Thu | Hash Map, part 1: hash function + buckets + Put/Get with chaining. |
| Fri | Hash Map, part 2: Delete + resizing on load factor. Fuzz against Go's `map`. |
| Sat | Review: re-write Reverse (singly list) and the hash map's resize from memory. |

### Week 3 — Stack, Queue, Trees, Heap, Trie, Union-Find
**Learn first:** VisuAlgo "BST" and "Binary Heap"; NeetCode's Trie and Union-Find videos.

| Day | Task |
| :--- | :--- |
| Mon | Stack (on your dynamic array) + Queue (ring buffer). |
| Tue | BST: Insert, Search, in-order traversal. |
| Wed | BST: Delete (leaf, one child, two children). |
| Thu | Min Heap: Push/Pop/heapify. Fuzz: pushing N values then popping them all must give sorted order. |
| Fri | Trie + Union-Find. |
| Sat | Review: re-write the heap and BST Delete from a blank file. |

---

## 🔄 Phase 2: Linear Patterns (Weeks 4-7)

LeetCode starts here. Solutions go in `leetcode/`. Hard (H) problems are deferred to the Hard Pass at the end.

### Week 4 — Arrays & Hashing, Two Pointers
**Learn first:** NeetCode "Arrays & Hashing" and "Two Pointers" intros; read about prefix sums.

| Day | Problems |
| :--- | :--- |
| Mon | 217 Contains Duplicate · 242 Valid Anagram |
| Tue | 1 Two Sum · 49 Group Anagrams |
| Wed | 238 Product of Array Except Self · 128 Longest Consecutive Sequence |
| Thu | 303 Range Sum Query · 560 Subarray Sum Equals K |
| Fri | 125 Valid Palindrome · 167 Two Sum II |
| Sat | Review week 4 |

### Week 5 — Two Pointers, Sliding Window, Stack
**Learn first:** NeetCode "Sliding Window" and "Stack" intros.

| Day | Problems |
| :--- | :--- |
| Mon | 15 3Sum · 11 Container With Most Water |
| Tue | 121 Best Time to Buy and Sell Stock · 3 Longest Substring Without Repeating Characters |
| Wed | 424 Longest Repeating Character Replacement · 567 Permutation in String |
| Thu | 20 Valid Parentheses · 155 Min Stack |
| Fri | 150 Evaluate Reverse Polish Notation · 739 Daily Temperatures (monotonic stack) |
| Sat | Review week 5 |

### Week 6 — Stack, Binary Search, Linked List
**Learn first:** NeetCode "Binary Search" intro — focus on the loop invariant and "binary search on the answer".

| Day | Problems |
| :--- | :--- |
| Mon | 853 Car Fleet · 704 Binary Search |
| Tue | 74 Search a 2D Matrix · 875 Koko Eating Bananas |
| Wed | 153 Find Minimum in Rotated Sorted Array · 33 Search in Rotated Sorted Array |
| Thu | 981 Time Based Key-Value Store · 206 Reverse Linked List |
| Fri | 21 Merge Two Sorted Lists · 141 Linked List Cycle |
| Sat | Review week 6 |

### Week 7 — Linked List + Phase 2 Checkpoint
**Learn first:** Floyd's cycle detection; the dummy-head-node trick.

| Day | Problems |
| :--- | :--- |
| Mon | 19 Remove Nth Node From End · 143 Reorder List |
| Tue | 92 Reverse Linked List II · 2 Add Two Numbers |
| Wed | 287 Find the Duplicate Number |
| Thu | 146 LRU Cache (reuse your doubly linked list + a map) |
| Fri | Catch-up day |
| Sat | **🏁 Checkpoint:** 2 unseen NeetCode 150 problems from Phase 2 topics, 45 min timed. If you can't solve either, repeat one week of Phase 2 review before moving on. |

---

## 🌳 Phase 3: Trees, Heaps, Backtracking, Graphs (Weeks 8-12)

### Week 8 — Trees
**Learn first:** NeetCode "Trees" intro; DFS (pre/in/post-order) vs BFS (level order). Write both recursively once, then iteratively once.

| Day | Problems |
| :--- | :--- |
| Mon | 226 Invert Binary Tree · 104 Maximum Depth |
| Tue | 100 Same Tree · 543 Diameter of Binary Tree |
| Wed | 110 Balanced Binary Tree · 102 Level Order Traversal |
| Thu | 199 Right Side View · 235 LCA of a BST |
| Fri | 98 Validate BST · 230 Kth Smallest in a BST |
| Sat | Review week 8 |

### Week 9 — Trees, Tries, Heaps
**Learn first:** Go's `container/heap` interface (you built your own; now learn the stdlib one).

| Day | Problems |
| :--- | :--- |
| Mon | 105 Construct Binary Tree from Preorder and Inorder |
| Tue | 208 Implement Trie · 211 Design Add and Search Words |
| Wed | 703 Kth Largest in a Stream · 1046 Last Stone Weight |
| Thu | 347 Top K Frequent Elements · 973 K Closest Points to Origin |
| Fri | 215 Kth Largest Element in an Array · 621 Task Scheduler |
| Sat | Review week 9 |

### Week 10 — Backtracking
**Learn first:** NeetCode "Backtracking" intro. Draw the decision tree for Subsets on paper before coding. Re-read the "Go Rule" on copying `path` in [dsa.md](dsa.md).

Only one problem a day this week: backtracking is where most people get stuck, so take the time to understand each one.

| Day | Problems |
| :--- | :--- |
| Mon | 78 Subsets |
| Tue | 90 Subsets II · 39 Combination Sum |
| Wed | 46 Permutations |
| Thu | 79 Word Search |
| Fri | 131 Palindrome Partitioning |
| Sat | Review week 10 |

### Week 11 — Graphs
**Learn first:** NeetCode "Graphs" intro; adjacency list vs grid; DFS/BFS with a `visited` set; Kahn's algorithm for topological sort.

| Day | Problems |
| :--- | :--- |
| Mon | 200 Number of Islands · 695 Max Area of Island |
| Tue | 133 Clone Graph · 994 Rotting Oranges |
| Wed | 417 Pacific Atlantic Water Flow · 130 Surrounded Regions |
| Thu | 207 Course Schedule |
| Fri | 210 Course Schedule II |
| Sat | Review week 11 |

### Week 12 — Union-Find, Advanced Graphs + Phase 3 Checkpoint
**Learn first:** Dijkstra (with your heap), Prim's MST, Bellman-Ford. VisuAlgo "SSSP" and "MST" are excellent here.

| Day | Problems |
| :--- | :--- |
| Mon | 684 Redundant Connection (your Union-Find) |
| Tue | 743 Network Delay Time (Dijkstra) |
| Wed | 1584 Min Cost to Connect All Points (MST) |
| Thu | 787 Cheapest Flights Within K Stops (Bellman-Ford) |
| Fri | Catch-up day |
| Sat | **🏁 Checkpoint:** 2 unseen NeetCode 150 problems from Phase 3 topics, 45 min timed. |

---

## 🧩 Phase 4: DP, Greedy & Specialty Topics (Weeks 13-16)

### Week 13 — 1-D Dynamic Programming
**Learn first:** NeetCode "1-D DP" intro; *Grokking* ch. 9. Method for every DP problem: **recursion → memoization → bottom-up table → reduce memory**.

| Day | Problems |
| :--- | :--- |
| Mon | 70 Climbing Stairs · 746 Min Cost Climbing Stairs |
| Tue | 198 House Robber · 213 House Robber II |
| Wed | 5 Longest Palindromic Substring · 91 Decode Ways |
| Thu | 322 Coin Change |
| Fri | 139 Word Break |
| Sat | Review week 13 |

### Week 14 — 1-D → 2-D Dynamic Programming
**Learn first:** 0/1 vs unbounded knapsack — the one-line difference in the loop order.

| Day | Problems |
| :--- | :--- |
| Mon | 300 Longest Increasing Subsequence |
| Tue | 416 Partition Equal Subset Sum (0/1 knapsack) |
| Wed | 62 Unique Paths · 1143 Longest Common Subsequence |
| Thu | 518 Coin Change II · 494 Target Sum |
| Fri | 309 Best Time to Buy and Sell Stock with Cooldown |
| Sat | Review week 14 |

### Week 15 — Greedy, Intervals
**Learn first:** NeetCode "Greedy" and "Intervals" intros; Kadane's algorithm.

| Day | Problems |
| :--- | :--- |
| Mon | 53 Maximum Subarray · 55 Jump Game |
| Tue | 45 Jump Game II · 134 Gas Station |
| Wed | 846 Hand of Straights · 763 Partition Labels |
| Thu | 56 Merge Intervals · 57 Insert Interval |
| Fri | 435 Non-overlapping Intervals · 136 Single Number |
| Sat | Review week 15 |

### Week 16 — Bit Manipulation, Math + Phase 4 Checkpoint
**Learn first:** XOR properties, `n & (n-1)`, shifts; Go's `math/bits`.

| Day | Problems |
| :--- | :--- |
| Mon | 191 Number of 1 Bits · 338 Counting Bits |
| Tue | 190 Reverse Bits · 268 Missing Number |
| Wed | 371 Sum of Two Integers · 202 Happy Number |
| Thu | 48 Rotate Image · 54 Spiral Matrix · 73 Set Matrix Zeroes |
| Fri | 50 Pow(x, n) · 43 Multiply Strings |
| Sat | **🏁 Checkpoint:** 2 unseen NeetCode 150 problems from Phase 4 topics, 45 min timed. |

---

## 🔥 Weeks 17-19: Hard Pass

One Hard problem a day. Give yourself 40 min before looking at the solution (instead of 20).

| Week | Mon | Tue | Wed | Thu | Fri |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 17 | 42 Trapping Rain Water | 76 Minimum Window Substring | 239 Sliding Window Maximum | 84 Largest Rectangle in Histogram | 23 Merge k Sorted Lists |
| 18 | 124 Binary Tree Max Path Sum | 297 Serialize/Deserialize Binary Tree | 212 Word Search II | 295 Find Median from Data Stream | 51 N-Queens |
| 19 | 127 Word Ladder | 778 Swim in Rising Water | 332 Reconstruct Itinerary | 72 Edit Distance | 1851 Minimum Interval to Include Each Query |

---

## 🚀 After Week 19

- Work through the rest of the **NeetCode 150**, mixing topics (not one topic at a time). Picking the right pattern when the topic isn't given is the real skill.
- Do the **Zero-Allocation Pass** from [dsa.md](dsa.md) on your 20 favourite solutions.
- Keep the Saturday review and Sunday 1-month re-solves going.
- Once the NeetCode 150 feels comfortable: weekly **LeetCode contests** for timed practice.
