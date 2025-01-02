# Problem: Min Cost Climbing Stairs

## Problem Statement

You are given an integer array `cost` where `cost[i]` is the cost of the *i*th step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor. The length of the `cost` array is guaranteed to be greater than or equal to 2.

## Constraints

*   `2 <= cost.length <= 1000`
*   `0 <= cost[i] <= 999`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function minCostClimbingStairs(cost)
```

### Go

```go
// Go
func minCostClimbingStairs(cost []int) int
```

## Example

```
Input: cost = [1,2,3]
Output: 2
Explanation: You start at index 1.
- Pay cost[1] = 2 to climb to the top.
The total cost is 2.

Input: cost = [10,15,20]
Output: 15
Explanation: You start at index 1. 
- Pay cost[1] = 15 to climb to the top.
The total cost is 15.

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You can start at index 0 and take steps of size 1, 1, 1, 1, 1 to reach the top.
Total cost will be 1 + 1 + 1 + 1 + 1 = 5. You can also start at index 0 and take steps of 1, 2, 1, 2, 1 to reach the top with cost 6.

```
## Approaches

This problem is a classic dynamic programming problem. Here are the common approaches:

### 1. Recursion (Naive)

This approach directly follows the recursive relationship: The minimum cost to reach step `i` is the minimum of the cost to reach step `i-1` plus `cost[i-1]` and the cost to reach `i-2` plus `cost[i-2]`. However, this leads to overlapping subproblems and exponential time complexity.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems. If we've already calculated the minimum cost to reach step `i`, we store it in a cache (e.g., an array or map) and return the cached value if we encounter the same subproblem again.

### 3. Tabulation (Bottom-Up Dynamic Programming)

Tabulation builds the solution iteratively. We create a DP array `dp` where `dp[i]` stores the minimum cost to reach step `i`. We initialize `dp[0] = cost[0]` and `dp[1] = cost[1]`. Then, we calculate the rest using the recurrence relation: `dp[i] = cost[i] + min(dp[i-1], dp[i-2])`. The minimum cost to reach the top is `min(dp[n-1], dp[n-2])`, where n is the length of cost.

### 4. Space Optimization

This optimizes the tabulation approach. We only need to store the minimum cost to reach the previous two steps. This reduces the space complexity to O(1).

## Complexity Analysis

| Approach              | Time Complexity | Space Complexity |
|-----------------------|-----------------|-----------------|
| 1. Recursion          | O(2^n)          | O(n)            |
| 2. Memoization        | O(n)            | O(n)            |
| 3. Tabulation         | O(n)            | O(n)            |
| 4. Space Optimization | O(n)            | O(1)            |

*   **Recursion:** Exponential time due to overlapping subproblems. O(n) space for the call stack.
*   **Memoization:** O(n) time as each subproblem is solved only once. O(n) space for the memoization cache.
*   **Tabulation:** O(n) time for the loop. O(n) space for the DP array.
*   **Space Optimization:** O(n) time for the loop. O(1) constant space.
