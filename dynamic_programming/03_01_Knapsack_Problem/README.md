# Problem: 0/1 Knapsack Problem

## Problem Statement

You are provided with a set of *N* items, each with a specified weight and value. Your objective is to pack these items into a backpack with a weight limit of *W*, maximizing the total value of items in the backpack.

Specifically, you have two arrays:

*   `val[0..N-1]`: representing the values of the items.
*   `wt[0..N-1]`: indicating their weights.

Additionally, you have a weight limit *W* for the backpack. The challenge is to determine the most valuable combination of items where the total weight does not exceed *W*. Note that each item is unique and indivisible, meaning it must be either taken as a whole or left entirely (hence "0/1 Knapsack").

## Constraints

*   `1 <= N <= 1000` (Number of items)
*   `1 <= W <= 1000` (Weight limit of the knapsack)
*   `1 <= val[i] <= 1000` (Value of each item)
*   `1 <= wt[i] <= 10` (Weight of each item)

## Function Signatures

### JavaScript

```javascript
// JavaScript
function knapsack(W, wt, val, n)
```

### Go

```go
// Go
func knapsack(W int, wt []int, val []int, n int) int
```

## Example
```
Example 1:
N = 3, W = 8
val = [3, 9, 5]
wt = [4, 5, 2]
Output: 14
Explanation: We can pick items with values 9 and 5 and weights 5 and 2 respectively. Total weight = 7 <= 8 and total value = 14.

Example 2:
N = 3, W = 4
val = [1, 2, 3]
wt = [4, 5, 1]
Output: 3
Explanation: We can only pick the last item with value 3 and weight 1.

Example 3:
N = 2, W = 10
val = [60, 100]
wt = [10, 20]
Output: 100
Explanation: We can pick the item with value 100 and weight 20 because the weight limit is 10.
```
## Approaches

This problem can be solved using dynamic programming. Here are the common approaches:

### 1. Recursion (Naive)

This approach explores all possible combinations of items by recursively considering whether to include each item or not. It has exponential time complexity due to overlapping subproblems.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a memo (cache). This avoids redundant calculations and reduces the time complexity to pseudo-polynomial (dependent on W).

### 3. Tabulation (Bottom-Up Dynamic Programming)

Tabulation builds the solution iteratively using a 2D DP table. `dp[i][w]` represents the maximum value that can be obtained using items up to index `i` with a maximum weight of `w`.

### 4. Space Optimization (using 1D arrays)

This approach optimizes the tabulation method by using only two 1D arrays (often named `prev` and `curr`) of size `W+1`. This reduces the space complexity to O(W). The `prev` array stores the results from the previous row (item), and the `curr` array stores the results for the current row. After each item is processed, the `prev` and `curr` arrays are swapped to simulate moving down the DP table.

## Complexity Analysis

| Approach              | Time Complexity | Space Complexity |
|-----------------------|-----------------|-----------------|
| 1. Recursion          | O(2^N)          | O(N)            |
| 2. Memoization        | O(N*W)          | O(N*W)          |
| 3. Tabulation         | O(N*W)          | O(N*W)          |
| 4. Space Optimization | O(N*W)          | O(W)            |

*   **Recursion:** Exponential time due to exploring all subsets. O(N) space for the call stack.
*   **Memoization/Tabulation:** Pseudo-polynomial time complexity O(N*W), as it depends on the value of W. O(N*W) space to store the memo/DP table.
*   **Space Optimization:** Pseudo-polynomial time complexity O(N*W). O(W) space due to using only two 1D arrays.