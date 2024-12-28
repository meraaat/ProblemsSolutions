# Problem: Subsets (Power Set)

## Problem Statement

Given an integer array `nums` of unique elements, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the solution in any order.

## Constraints

*   All elements in `nums` are unique.

## Function Signatures

### Javascript
```javascript
// JavaScript
function subsets(nums)
```
### Go
```go
//Go
func subsets(nums int[]) [][]int
```

## Example

```
Input: nums = [1, 2, 3]
Output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]

Input: nums = [0]
Output: [[], [0]]

Input: nums = []
Output: [[]]

```

## Approach 1: Backtracking

Backtracking is a suitable approach for generating subsets.

1.  **Base Case:** When we've considered all elements in `nums` (reached the end of the input array), the current subset is complete. Add it to the result.

2.  **Recursive Step:** For each element at index `i` in `nums`, we have two choices:
    *   **Include the element:** Add `nums[i]` to the current subset and recursively call the function for the next element (`i + 1`).
    *   **Exclude the element:** Do not add `nums[i]` to the current subset and recursively call the function for the next element (`i + 1`).

3. Backtracking: After exploring both the include and exclude cases, we remove the last added element from the current subset (if we added it). This is essential to explore other subset combinations.

*   Time Complexity: O(2^n), where n is the length of `nums`. There are 2^n possible subsets.
*   Space Complexity: O(n) due to the recursion depth. The space to store the output is not included in the space complexity analysis.