# Problem: Unbounded Knapsack

## Problem Statement

You are given a set of *N* items, each with a weight and a value, represented by the arrays `wt` and `val` respectively, and a knapsack with weight limit *W*. The task is to fill the knapsack in such a way that you can get the maximum profit.

**Note:** Each item can be taken any number of times (this is the key difference from the 0/1 Knapsack problem).

## Constraints

*   `1 <= N <= 1000` (Number of items)
*   `1 <= W <= 1000` (Weight limit of the knapsack)
*   `1 <= val[i] <= 1000` (Value of each item)
*   `1 <= wt[i] <= 10` (Weight of each item)

## Function Signatures

### JavaScript

```javascript
// JavaScript
function unboundedKnapsack(W, wt, val, n)
```

### Go

```go
// Go
func unboundedKnapsack(W int, wt []int, val []int, n int) int
```

## Example
```
Example 1:
N = 2, W = 6
val = [5, 10]
wt = [2, 2]
Output: 30
Explanation: We can take the second item (value 10, weight 2) three times. Total weight = 2 * 3 = 6. Total value = 10 * 3 = 30.

Example 2:
N = 3, W = 8
val = [1, 4, 5]
wt = [1, 3, 4]
Output: 11
Explanation: We can take the first item (value 1, weight 1) three times and the third item (value 5, weight 4) once. Total weight = 3*1 + 1*4 = 7 <= 8. Total value = 3*1 + 1*5 = 8. Another option is to take the second item twice and first item twice. Total weight = 2*3 + 2*1 = 8. Total value = 2*4 + 2*1 = 10. Taking the third item twice leads to total weight 8 and value 10. Taking the first item 8 times gives a total value of 8. The best solution is taking the third item once and the first item three times, with a total value of 8. Taking the second item twice and the first item twice gives a total value of 10. Taking the third item twice gives a total value of 10. Taking the first item eight times gives a total value of 8. Taking the second item once and the third item once gives a total value of 9. Taking the first item twice and the second item once gives a total value of 6.

Example 3:
N = 3, W = 4
val = [1, 2, 3]
wt = [4, 5, 1]
Output: 3
Explanation: We can only pick the third item with value 3 and weight 1 four times. Total Weight = 4*1 = 4 and Total value = 4*3 = 12.

```
## Approaches

This problem can be efficiently solved using dynamic programming with tabulation.

### Tabulation (Bottom-Up Dynamic Programming)

Tabulation builds the solution iteratively using a 1D DP array (unlike the 2D array used in the 0/1 Knapsack tabulation). `dp[w]` represents the maximum value that can be obtained with a maximum weight of `w`.

The recurrence relation is:

`dp[w] = max(dp[w], val[i] + dp[w - wt[i]])` for all items `i` where `wt[i] <= w`.

## Complexity Analysis

| Approach      | Time Complexity | Space Complexity |
|---------------|-----------------|-----------------|
| Tabulation | O(N*W)          | O(W)            |

*   **Tabulation:** Pseudo-polynomial time complexity O(N*W), as it depends on the value of W. O(W) space to store the DP table.