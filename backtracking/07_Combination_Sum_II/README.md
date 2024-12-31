# Problem: Combination Sum II

## Problem Statement

Given a collection of candidate numbers (`candidates`) and a target number (`target`), find all unique combinations in `candidates` where the candidate numbers sum to `target`.

Each number in `candidates` may only be used *once* in the combination. The solution set must not contain duplicate combinations.

## Constraints

*   `1 <= candidates.length <= 100`
*   `1 <= candidates[i] <= 50`
*   `1 <= target <= 30`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function combinationSum2(candidates, target)
```

### Go
```go
// Go
func combinationSum2(candidates []int, target int) [][]int
```

## Example 

```
Input: candidates = [3, 5, 2, 1, 3], target = 7
Output: [[1, 2, 3], [2, 5]]

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output: [[1,1,6],[1,2,5],[1,7],[2,6]]

Input: candidates = [2,5,2,1,2], target = 5
Output: [[1,2,2],[5]]
```
## Approach: Backtracking with Duplicate Handling

This problem is similar to Combination Sum, but with the added constraint that each number can only be used *once* in a combination, and we must avoid duplicate combinations in the result.

### Steps:

1.  **Sort the Input:** Sort the `candidates` array in non-decreasing order. This is essential for efficiently handling duplicates.

2.  **Base Cases:**
    *   If `currentSum` equals `target`, add a *copy* of the `current` combination to the `result`.
    *   If `currentSum` exceeds `target`, return (prune the search).

3.  **Recursive Step:** Iterate through the `candidates` array starting from the current `start` index:
    *   **Duplicate Handling:** If the current candidate is the same as the previous candidate (`candidates[i] == candidates[i-1]`) *and* `i > start` (to avoid skipping the first occurrence), then skip the current candidate to prevent duplicate combinations.
    *   **Include the candidate:** Add `candidates[i]` to the `current` combination.
    *   **Recursive Call:** Call the backtracking function with `start = i + 1` (because each number can only be used once).
    *   **Backtrack:** Remove the last added candidate from `current`.

## Complexity Analysis

*   **Time Complexity:** The time complexity is still exponential, similar to Combination Sum, but the duplicate handling significantly reduces the number of explored branches. The worst-case time complexity is still related to O(2^n) where n is the length of candidates but the sorting and duplicate handling make it much faster in practice.

*   **Space Complexity:** O(n) in the worst case, due to the recursion stack depth. The space to store the result is not considered in the auxiliary space complexity. The sorting takes O(n log n) space in some implementations, but can be done in place in others.