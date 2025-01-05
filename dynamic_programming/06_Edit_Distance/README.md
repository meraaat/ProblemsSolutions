# Problem: Edit Distance

## Problem Statement

Given two strings, `word1` and `word2`, return the minimum number of operations required to convert `word1` to `word2`.

You have the following three operations permitted on a word:

*   Insert a character
*   Delete a character
*   Replace a character

## Constraints

*   `0 <= word1.length, word2.length <= 500`
*   `word1` and `word2` consist of lowercase English letters.

## Function Signatures

### JavaScript

```javascript
// JavaScript
function minDistance(word1, word2)
```

### Go

```go
// Go
func minDistance(word1 string, word2 string) int
```

## Example
```
Example 1:
word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')

Example 3:
word1 = "", word2 = ""
Output: 0
Explanation: No operations are needed since the strings are empty

Example 4:
word1 = "a", word2 = ""
Output: 1
Explanation: Delete 'a' from word1

Example 5:
word1 = "", word2 = "b"
Output: 1
Explanation: Insert 'b' into word1


```
## Approaches

This problem can be solved using dynamic programming.

### 1. Recursion (Naive)

This approach directly implements the recursive definition of Edit Distance. It explores all possible operations, leading to exponential time complexity due to overlapping subproblems.

### 2. Memoization (Top-Down Dynamic Programming)

Memoization optimizes the recursive approach by storing the results of subproblems in a memo (cache). This avoids redundant calculations and reduces the time complexity to O(m\*n).

### 3. Tabulation (Bottom-Up Dynamic Programming)

We use a 2D array `dp` where `dp[i][j]` represents the minimum number of operations to convert `word1[0...i-1]` to `word2[0...j-1]`.

The recurrence relation is as follows:

*   If `word1[i-1] == word2[j-1]`: `dp[i][j] = dp[i-1][j-1]` (No operation needed if characters match)
*   If `word1[i-1] != word2[j-1]`: `dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])`
    *   `dp[i-1][j]` represents deletion from `word1`.
    *   `dp[i][j-1]` represents insertion into `word1`.
    *   `dp[i-1][j-1]` represents replacement.

The base cases are:

*   `dp[0][j] = j`
*   `dp[i][0] = i`

### 4. Space Optimization

We can optimize the tabulation approach by using only two 1D arrays (representing the current and previous rows of the DP table). This reduces the space complexity from O(m\*n) to O(min(m, n)).

## Complexity Analysis

| Approach              | Time Complexity | Space Complexity |
|-----------------------|-----------------|-----------------|
| 1. Recursion          | O(3^min(m,n))      | O(min(m,n))          |
| 2. Memoization        | O(m*n)          | O(m*n)          |
| 3. Tabulation         | O(m*n)          | O(m*n)          |
| 4. Space Optimization | O(m*n)          | O(min(m, n))     |

*   **Recursion:** Exponential time due to exploring all possibilities. The actual complexity is closer to O(3^min(m,n)) than O(2^(m+n)). The space is O(min(m,n)) due to recursion stack.
*   **Memoization/Tabulation:** O(m\*n) time where m and n are the lengths of `word1` and `word2` respectively. O(m\*n) space to store the memo/DP table.
*   **Space Optimization:** O(m\*n) time. O(min(m, n)) space because we only need to store two rows (or columns, depending on the implementation).