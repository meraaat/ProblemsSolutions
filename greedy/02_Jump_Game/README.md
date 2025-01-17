# Problem: Jump Game

## Problem Statement

You are given an integer array `nums`. You are initially positioned at the array's first index, and each element `nums[i]` represents your maximum jump length at that position `i`.

Return `true` if you can reach the last index, or `false` otherwise.

## Constraints

*   `1 <= nums.length <= 10^4`
*   `0 <= nums[i] <= 10^5`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function canJump(nums)
```

### Go

```go
// Go
func canJump(nums []int) bool
```

## Example

```
Example 1:
nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Example 3:
nums = [1,3,4,1,1,2,1]
Output: true

Example 4:
nums = [1,3,4,2,1,1,0,1,1]
Output: false

Example 5:
nums = [0]
Output: true

Example 6:
nums = [2,0,0]
Output: true

```

## Approaches

This problem can be solved using several approaches. The most common and efficient ones are Greedy and Dynamic Programming.

### 1. Greedy Approach (Optimal)

This is the most efficient approach with O(n) time complexity.

1.  Initialize a variable `reachable` to 0. This variable keeps track of the farthest index that can be reached so far.

2.  Iterate through the `nums` array from left to right.

3.  At each index `i`, check if `i > reachable`. If it is, it means we cannot reach the current index, so return `false`.

4.  Otherwise, update `reachable` as `reachable = max(reachable, i + nums[i])`. This extends the reachable range.

5.  After the loop, if `reachable` is greater than or equal to `nums.length - 1`, it means we can reach the last index, so return `true`.

### 2. Dynamic Programming (Less Efficient)

This approach uses a DP array to store whether each index is reachable. While it works, the greedy approach is more efficient.

1.  Create a boolean DP array `dp` of the same length as `nums`. `dp[i]` will be `true` if index `i` is reachable, and `false` otherwise.

2.  Initialize `dp[0]` to `true`.

3.  Iterate through the `nums` array. For each index `i` where `dp[i]` is `true`, iterate from `j = 1` to `nums[i]` (the maximum jump length) and set `dp[i + j]` to `true` if `i + j` is within the array bounds.

4.  Return `dp[nums.length - 1]`.

## Complexity Analysis

| Approach        | Time Complexity | Space Complexity |
|-----------------|-----------------|-----------------|
| Greedy          | O(n)            | O(1)            |
| Dynamic Programming | O(n^2)          | O(n)            |

*   **Greedy:** Linear time because we iterate through the array once. Constant space because we only use a few variables.
*   **Dynamic Programming:** O(n^2) time due to the nested loops. O(n) space for the DP array.

