# Problem: Permutations

## Problem Statement

Given an array `nums` of distinct integers, return all possible permutations. You can return the answer in any order.

## Constraints

*   `nums` will have at least one element.
*   All integers in `nums` are distinct.

## Function Signatures

### Javascript
```javascript
// JavaScript
function permute(nums)
```

### Go
```go
// Go
func permute(nums []int) [][]int
```

## Example
```
Input: nums = [1, 2, 3]
Output: [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Input: nums = [1]
Output: [[1]]
```

## Approach: Backtracking

This is a common and efficient approach for generating permutations.

1.  **Base Case:** If the current permutation has the same length as the input `nums`, it's a complete permutation. Add it to the result.

2.  **Recursive Step:** Iterate through the input `nums`.
    *   For each number, check if it's already included in the current permutation.
    *   If not, add it to the current permutation.
    *   Recursively call the permutation function to generate permutations for the remaining numbers.
    *   **Backtrack:** After the recursive call returns, remove the last added number from the current permutation. This is crucial for exploring other possibilities.

*   Time Complexity: O(n!), where n is the length of `nums`. There are n! possible permutations.
*   Space Complexity: O(n) due to the recursion depth and the space used to store the current permutation. The space required to store the output is not considered in the space complexity analysis, as it is the expected output.