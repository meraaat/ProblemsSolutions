# Problem: Maximum Length of Pair Chain

## Problem Statement

You are given an array of `n` pairs `pairs`, where `pairs[i] = [left_i, right_i]` and `left_i < right_i`.

A pair `p2 = [c, d]` follows a pair `p1 = [a, b]` if `b < c`. A chain of pairs can be formed in this fashion.

Return the length of the longest chain that can be formed. You do not need to use up all the given intervals. You can select pairs in any order.

## Constraints

*   `1 <= pairs.length <= 1000`
*   `-1000 <= left_i < right_i <= 1000`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function findLongestChain(pairs)
```

### Go

```go
// Go
func findLongestChain(pairs [][]int) int
```

## Example

```
Example 1:
pairs = [[1,2],[2,3],[3,4]]
Output: 2
Explanation: The longest chain is [1,2] -> [3,4].

Example 2:
pairs = [[1,2],[7,8],[4,5]]
Output: 3
Explanation: The longest chain is [1,2] -> [4,5] -> [7,8].

Example 3:
pairs = [[1,2],[2,9],[4,5]]
Output: 2
Explanation: The longest chain is [1,2] -> [4,5].

Example 4:
pairs = [[1,2]]
Output: 1

Example 5:
pairs = []
Output: 0

```
## Approaches

This problem can be solved using several approaches: recursion, memoization (top-down DP), tabulation (bottom-up DP), and a greedy approach.

### 1. Recursion (Naive)

This approach explores all possible combinations of pairs, leading to exponential time complexity due to overlapping subproblems. It involves sorting the pairs by their left values first.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a memo (cache). This avoids redundant calculations and reduces the time complexity to O(n^2). It also requires sorting the pairs by their left values.

### 3. Tabulation (Bottom-Up Dynamic Programming)

1.  **Sort the pairs** based on their *left* values.
2.  Use a DP array `dp` where `dp[i]` represents the length of the longest chain ending at `pairs[i]`.
3.  For each `i`, iterate through `j` from 0 to `i-1`. If `pairs[j][1] < pairs[i][0]`, then `dp[i] = max(dp[i], dp[j] + 1)`.
4.  The base case is `dp[i] = 1` for all `i`.
5.  The result is the maximum value in `dp`.

### 4. Greedy Approach (Optimal)

This approach is the most efficient.

1.  **Sort the pairs** based on their *right* values.
2.  Initialize `currentEnd` to negative infinity and `chainLength` to 0.
3.  Iterate through the sorted pairs. If `pair[0] > currentEnd`, it means this pair can be added to the chain. Update `currentEnd` to `pair[1]` and increment `chainLength`.

## Complexity Analysis

| Approach           | Time Complexity | Space Complexity |
|--------------------|-----------------|-----------------|
| 1. Recursion       | O(2^n)          | O(n)            |
| 2. Memoization     | O(n^2)          | O(n)            |
| 3. Tabulation      | O(n^2)          | O(n)            |
| 4. Greedy         | O(n log n)      | O(1) or O(n) if sorting is considered|

*   **Recursion:** Exponential time due to exploring all subsets. O(n) space for the call stack.
*   **Memoization/Tabulation:** O(n^2) time due to the nested loops. O(n) space to store the memo/DP array.
*   **Greedy:** O(n log n) due to sorting. The space complexity is O(1) if sorting is done in place or O(n) if a new array is created while sorting.