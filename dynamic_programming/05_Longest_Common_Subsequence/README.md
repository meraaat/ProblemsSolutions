# Problem: Longest Common Subsequence (LCS)

## Problem Statement

Given two strings, `text1` and `text2`, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".

A common subsequence of two strings is a subsequence that is common to both strings.

## Constraints

*   `1 <= text1.length, text2.length <= 1000`
*   `text1` and `text2` consist of lowercase English characters.

## Function Signatures

### JavaScript

```javascript
// JavaScript
function longestCommonSubsequence(text1, text2)
```

### Go

```go
// Go
func longestCommonSubsequence(text1 string, text2 string) int
```

## Example
```
Example 1:
text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:
text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:
text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

Example 4:
text1 = "abcdef", text2 = "acef"
Output: 4
Explanation: The longest common subsequence is "acef" and its length is 4.

```
## Approaches

This problem can be solved using several dynamic programming approaches.

### 1. Recursion (Naive)

This approach directly implements the recursive definition of LCS. It explores all possible subsequences, leading to exponential time complexity due to overlapping subproblems.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a memo (cache). This avoids redundant calculations and reduces the time complexity.

### 3. Tabulation (Bottom-Up Dynamic Programming)

Tabulation builds the solution iteratively using a 2D DP table. `dp[i][j]` stores the length of the longest common subsequence of `text1[0...i-1]` and `text2[0...j-1]`.

The recurrence relation is as follows:

*   If `text1[i-1] == text2[j-1]`: `dp[i][j] = dp[i-1][j-1] + 1`
*   If `text1[i-1] != text2[j-1]`: `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`

The base cases are `dp[0][j] = 0` and `dp[i][0] = 0` for all `i` and `j`.

### 4. Space Optimization

We can optimize the tabulation approach by observing that to calculate `dp[i][j]`, we only need the values from the current row (`i`) and the previous row (`i-1`). Therefore, we can use two 1D arrays to store these rows, reducing the space complexity.

## Complexity Analysis

| Approach              | Time Complexity | Space Complexity |
|-----------------------|-----------------|-----------------|
| 1. Recursion          | O(2^(m+n))      | O(m+n)           |
| 2. Memoization        | O(m*n)          | O(m*n)           |
| 3. Tabulation         | O(m*n)          | O(m*n)           |
| 4. Space Optimization | O(m*n)          | O(min(m, n))     |

*   **Recursion:** Exponential time due to exploring all subsets. O(m+n) space for the call stack.
*   **Memoization/Tabulation:** O(m\*n) time where m and n are the lengths of `text1` and `text2` respectively. O(m\*n) space to store the memo/DP table.
*   **Space Optimization:** O(m\*n) time. O(min(m, n)) space because we only need to store two rows (or columns, depending on the implementation), and we choose the smaller dimension to optimize space usage further.