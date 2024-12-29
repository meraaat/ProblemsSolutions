# Problem: Subsets II (Power Set with Duplicates)

## Problem Statement

Given an integer array `nums` that may contain duplicates, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the solution in any order.

## Constraints

*   `nums` may contain duplicate elements.

## Function Signatures

### JavaScript

```javascript
// JavaScript
function subsetsWithDup(nums)
```

### GO
```go
// Go
func subsetsWithDup(nums []int) [][]int
```

## Example 
```
Input: nums = [1, 2, 2]
Output: [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]]

Input: nums = [0]
Output: [[], [0]]

Input: nums = []
Output: [[]]
```

## Approach: Backtracking with Duplicate Handling

This problem is similar to generating subsets of an array with unique elements, but with the added constraint of handling duplicates. Backtracking is still a suitable approach, but we need to modify the logic to avoid generating duplicate subsets.

### Steps:

1.  **Sort the Input:** Sort the input array `nums` in ascending order. This is crucial for efficiently identifying and skipping duplicate elements.

2.  **Base Case:** When we've considered all elements in `nums`, add the current subset to the result.

3.  **Recursive Step:** For each element at index `i` in `nums`, we have two choices:
    *   **Include the element:** Add `nums[i]` to the current subset and recursively call the function for the next element (`i + 1`).
    *   **Exclude the element:** This is where we handle duplicates. If the current element `nums[i]` is the same as the previous element `nums[i - 1]` *and* the previous element was *not* included in the current subset (meaning we just backtracked from it), then we skip the current element to avoid generating duplicate subsets. We only proceed with the recursive call for the next element (`i + 1`) if either `nums[i]` is different from `nums[i-1]` or if the previous element *was* included.

4.  **Backtracking:** After exploring both the include and exclude cases, remove the last added element from the current subset (if we added it).

## Complexity Analysis

*   **Time Complexity:** \(O(2^n)\), where \(n\) is the length of `nums`. In the worst-case scenario (no duplicates), we have \(2^n\) possible subsets. The duplicate handling doesn't change the overall exponential nature.

*   **Space Complexity:** \(O(n)\) for the recursion depth, as the maximum depth of the call stack is \(n\). The space to store the output is \(O(n * 2^n)\), which is not included in the space complexity analysis. The sorting takes \(O(n log n)\) space in some implementations, but can be done in place in others.