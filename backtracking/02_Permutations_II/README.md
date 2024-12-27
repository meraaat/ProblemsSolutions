# Problem: Permutations II (with Duplicates)

## Problem Statement

Given a collection of numbers, `nums`, that might contain duplicates, return all possible *unique* permutations in any order.

## Constraints

*   The input array `nums` can contain duplicate numbers.

## Function Signatures

### Javascript
```javascript
// JavaScript
function permuteUnique(nums)
```
### Go
```go
// Go
func permuteUnique(nums []int) [][]int
```
## Example
```
Input: nums = [1, 1, 2]
Output: [[1, 1, 2], [1, 2, 1], [2, 1, 1]]

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Input: nums = [1,1,1]
Output: [[1,1,1]]

Input: nums = [1]
Output: [[1]]

Input: nums = []
Output: []
```
## Approach : Backtracking with Duplicate Handling

This approach builds upon the backtracking method for generating permutations but adds a crucial step to handle duplicates.

1.  **Sort the Input:** Sort the input array `nums` in ascending order. This is essential for efficiently detecting and skipping duplicate permutations.

2.  **Base Case:** If the current permutation has the same length as the input `nums`, it's a complete permutation. Add it to the result.

3.  **Recursive Step:** Iterate through the input `nums`.
    *   **Duplicate Check:** Before adding a number to the current permutation, check if it's the same as the previous number in the `remainingNums` and if the previous number hasn't been used in the current recursion level. This condition is `i > 0 && remainingNums[i] == remainingNums[i-1]`. If it is the same and the previous number hasn't been used it means it has been already used in the previous recursive call and we should skip this number to avoid duplicate permutations.
    *   If the current number is not a duplicate (or it's the first number), add it to the current permutation.
    *   Recursively call the permutation function.
    *   **Backtrack:** After the recursive call, remove the last added number from the current permutation.

*   Time Complexity: O(n!), where n is the length of `nums`, in the worst case (when there are no duplicates). In cases with many duplicates, the time complexity can be significantly less because many branches of the recursion tree are pruned.
*   Space Complexity: O(n) due to the recursion depth and the space used to store the current permutation.