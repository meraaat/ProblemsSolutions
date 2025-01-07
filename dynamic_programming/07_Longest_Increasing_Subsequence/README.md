# Problem: Longest Increasing Subsequence (LIS)

## Problem Statement

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

## Constraints

*   `1 <= nums.length <= 2500`
*   `-10^4 <= nums[i] <= 10^4`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function lengthOfLIS(nums)
```

### Go

```go
// Go
func lengthOfLIS(nums []int) int
```

## Example

```
Example 1:
nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
nums = [0,1,0,3,2,3]
Output: 4
Explanation: The longest increasing subsequence is [0,1,2,3], therefore the length is 4.

Example 3:
nums = [7,7,7,7,7,7,7]
Output: 1
Explanation: The longest increasing subsequence is [7], therefore the length is 1.

Example 4:
nums = [1,2,1,4,3,4,5]
Output: 5
Explanation: The longest increasing subsequence is [1,2,3,4,5], therefore the length is 5.

```

## Approaches

This problem can be solved using several approaches, including dynamic programming and a more efficient approach using patience sorting with binary search.

### 1. Recursion (Naive)

This approach explores all possible subsequences, leading to exponential time complexity due to overlapping subproblems.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a memo (cache). This avoids redundant calculations and reduces the time complexity to O(n^2).

### 3. Tabulation (Bottom-Up Dynamic Programming)

We can use a 1D DP array `dp` where `dp[i]` stores the length of the longest increasing subsequence ending at index `i`.

The recurrence relation is as follows:

`dp[i] = max(dp[j] + 1)` for all `j < i` where `nums[j] < nums[i]`.

The base case is `dp[i] = 1` for all `i`.

The final result is the maximum value in the `dp` array.

### 4. Patience Sorting with Binary Search (More Efficient)

This approach uses a "piles" concept. We iterate through the `nums` array and maintain a list of piles. For each number:

1.  If the number is greater than all the top elements of existing piles, we create a new pile with that number.
2.  Otherwise, we find the pile with the smallest top element that is greater than or equal to the current number using binary search and replace that top element with the current number.

The length of the LIS is the number of piles at the end. This method has a time complexity of O(n log n).

## Complexity Analysis

| Approach                               | Time Complexity | Space Complexity |
|----------------------------------------|-----------------|-----------------|
| 1. Recursion                           | O(2^n)          | O(n)            |
| 2. Memoization                         | O(n^2)          | O(n)            |
| 3. Dynamic Programming (Tabulation)    | O(n^2)          | O(n)            |
| 4. Patience Sorting with Binary Search | O(n log n)      | O(n)            |

*   **Recursion:** Exponential time due to exploring all subsets. O(n) space for the call stack.
*   **Memoization/Tabulation:** O(n^2) time due to the nested loops. O(n) space to store the memo/DP array.
*   **Patience Sorting with Binary Search:** O(n log n) time due to the binary search in each iteration. O(n) space to store the piles.