# Problem: Combination Sum

## Problem Statement

Given an array of *distinct* integers `candidates` and a target integer `target`, return a list of all *unique* combinations of `candidates` where the chosen numbers sum to `target`. You may return the combinations in any order.

The *same* number may be chosen from `candidates` an *unlimited number of times*. Two combinations are unique if the frequency of at least one of the chosen numbers is different. (You will not be given an empty `candidates` array)

It is also given that all the integers in the candidates array are non-negative.

## Constraints

*   `1 <= candidates.length <= 30`
*   `1 <= candidates[i] <= 200`
*   All elements of `candidates` are distinct.
*   `1 <= target <= 500`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function combinationSum(candidates, target)
```

### Go
```go
// Go
func combinationSum(candidates []int, target int) [][]int
```

## Example

```
Input: candidates = [2, 3, 6, 7], target = 7
Output: [[2, 2, 3], [7]]

Input: candidates = [2, 3, 5], target = 8
Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]

Input: candidates = [2], target = 1
Output: []
```

## Approach: Backtracking

Backtracking is a suitable approach for this problem. We explore different combinations by recursively trying each candidate and either including it in the current combination or skipping it. Because we can use the same number multiple times, we don't increment the starting index in the recursive call when we include a number.

### Steps:

1.  **Base Cases:**
    *   If the current sum equals the `target`, add a copy of the current combination to the result.
    *   If the current sum exceeds the `target`, we can stop exploring this branch (no need to continue).

2.  **Recursive Step:** For each candidate at index `i` (from the current `start` index to handle duplicates if the input was not distinct, but it is in this problem):
    *   **Include the candidate:** Add `candidates[i]` to the current combination.
    *   **Recursive Call:** Call the backtracking function again. *Crucially*, we do *not* increment the starting index (`i+1`) in the recursive call. This allows us to reuse the same candidate multiple times.
    *   **Backtrack:** Remove the last added candidate from the current combination to explore other possibilities.

## Complexity Analysis

*   **Time Complexity:** The time complexity is difficult to express with a simple Big-O notation due to the nature of the problem. In the worst-case scenario, where the target is large and the candidates are small, the number of combinations can grow exponentially. A loose upper bound could be considered O(N^(T/M + 1)), where N is the number of candidates, T is the target, and M is the minimum value in candidates. This is because in the worst case, we might explore all possible combinations of using the smallest candidate to reach the target.

*   **Space Complexity:** The space complexity is O(T), where T is the target value. This is due to the maximum depth of the recursion stack, which can be at most T (in the scenario where we only use the smallest candidate repeatedly to reach the target). The space to store the result is not considered in the auxiliary space complexity.