# Problem: Fibonacci Number

## Problem Statement

In the Fibonacci sequence, each subsequent term is obtained by adding the preceding two terms. This is true for all numbers except the first two numbers of the Fibonacci series, as they do not have two preceding numbers. The first two terms in the Fibonacci series are 0 and 1.

The Fibonacci sequence is defined as follows:

*   F(0) = 0
*   F(1) = 1
*   F(n) = F(n-1) + F(n-2) for n > 1

Write a function that finds F(n) given n, where n is an integer greater than or equal to 0.

## Constraints

*   `0 <= n <= 30` (A reasonable constraint for basic implementations; can be adjusted depending on the need for handling larger numbers)

## Function Signatures

### JavaScript

```javascript
// JavaScript
function fib(n)
```

### Go

```go
// Go
func fib(n int) int
```

## Example

```
Input: n = 0
Output: 0

Input: n = 1
Output: 1

Input: n = 2
Output: 1 (F(2) = F(1) + F(0) = 1 + 0 = 1)

Input: n = 3
Output: 2 (F(3) = F(2) + F(1) = 1 + 1 = 2)

Input: n = 4
Output: 3 (F(4) = F(3) + F(2) = 2 + 1 = 3)

Input: n = 10
Output: 55
```

## Approaches

There are several ways to calculate Fibonacci numbers:

### 1. Recursion (Naive)

This is the most straightforward implementation based directly on the mathematical definition: F(n) = F(n-1) + F(n-2). However, it's highly inefficient for larger values of `n` due to redundant calculations. The same subproblems are calculated many times, leading to exponential time complexity.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a cache (e.g., a map or array). Before calculating F(n), we check if it's already in the cache. If it is, we return the cached value; otherwise, we calculate it, store it in the cache, and then return it.

### 3. Tabulation (Bottom-Up Dynamic Programming)

This is an iterative approach that builds the solution from the bottom up. We start with the base cases (F(0) = 0 and F(1) = 1) and then iteratively calculate F(2), F(3), ..., F(n) by using the previously calculated values.

### 4. Space Optimization

This is a further optimization of the tabulation approach. Instead of storing all the calculated Fibonacci numbers in an array, we only need to keep track of the two most recent values. This reduces the space complexity from O(n) to O(1).

## Complexity Analysis

Here's a comparison of the time and space complexity of the different approaches:

| Approach              | Time Complexity | Space Complexity |
|-----------------------|-----------------|-----------------|
| 1. Recursion          | O(2^n)          | O(n)            |
| 2. Memoization        | O(n)            | O(n)            |
| 3. Tabulation         | O(n)            | O(n)            |
| 4. Space Optimization | O(n)            | O(1)            |

*   **Recursion:** The exponential time complexity O(2^n) is due to the overlapping subproblems. The space complexity O(n) is due to the call stack depth.
*   **Memoization:** The time complexity is reduced to O(n) because each subproblem is calculated only once. The space complexity is O(n) to store the memoized values.
*   **Tabulation:** The time complexity is O(n) because we iterate from 0 to n. The space complexity is O(n) to store the DP array.
*   **Space Optimization:** The time complexity remains O(n), but the space complexity is reduced to O(1) because we only store a constant number of variables.