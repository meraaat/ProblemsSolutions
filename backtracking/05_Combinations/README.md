# Problem: Combinations

## Problem Statement

Given two integers `n` and `k`, return all possible combinations of `k` numbers chosen from the range `[1, n]`. You may return the answer in any order.

## Constraints

*   `1 <= n <= 20`
*   `1 <= k <= n`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function combine(n, k)
```

### GO
```go
// Go
func combine(n int, k int) [][]int
```

## Example

```
Input: n = 4, k = 2
Output: [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]

Input: n = 1, k = 1
Output: [[1]]
```

## Approach: Backtracking

Backtracking is a suitable approach for generating combinations. The idea is to explore each number from 1 to `n` and decide whether to include it in the current combination or not.

### Steps:

1.  **Base Case:** When the current combination has `k` elements, add it to the result.

2.  **Recursive Step:** For each number `i` from `start` to `n` (where `start` is the next number to consider to avoid duplicates like [2,1] after [1,2]):
    *   **Include the number:** Add `i` to the current combination and recursively call the function for the next number (`i + 1`).

3. **Backtracking:** After the recursive call, remove the last added number from the current combination to explore other possibilities.

### Optimization: Early Pruning

We can optimize the backtracking process by adding an early pruning condition. If the remaining numbers to choose are more than the numbers we need to pick to reach k, we can stop exploring that branch.

## Complexity Analysis

*   **Time Complexity:** The time complexity is difficult to express with a simple Big-O notation. It's related to the number of combinations, which is given by the binomial coefficient "n choose k" (nCk). In the worst case, where k is close to n/2, the number of combinations can be significant. A loose upper bound can be O(k * nCk) because we are creating copies of k sized arrays nCk times.

*   **Space Complexity:** \(O(k)\) for the recursion depth and the space used to store each combination. The space for storing the final result is proportional to the number of combinations (nCk) and the size of each combination (k), but it is not considered part of the auxiliary space complexity.