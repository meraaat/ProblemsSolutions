# Problem: K-th Symbol in Grammar

## Problem Statement

We build a table of `n` rows (1-indexed). We start by writing `0` in the 1st row. Now in every subsequent row, we look at the previous row and replace each occurrence of `0` with `01`, and each occurrence of `1` with `10`.

For example, for `n = 3`, the 1st row is `0`, the 2nd row is `01`, and the 3rd row is `0110`.

Given two integers `n` and `k`, return the `k`th (1-indexed) symbol in the `n`th row of a table of `n` rows.

## Constraints

*   `n` will be an integer greater than or equal to 1 (n >= 1).
*   `k` will be within the bounds of the length of the `n`th row (1 <= k <= 2^(n-1)).

## Function Signatures

### JavaScript
```javascript
// JavaScript
function kthGrammar(n, k)
```

### Go
```go
// Go
func KthGrammar(n int, k int) int
```

## Example
```
Input: n = 1, k = 1
Output: 0

Input: n = 2, k = 1
Output: 0

Input: n = 2, k = 2
Output: 1

Input: n = 3, k = 1
Output: 0

Input: n = 3, k = 2
Output: 1

Input: n = 3, k = 3
Output: 1

Input: n = 3, k = 4
Output: 0

Input: n = 4, k = 5
Output: 1

Input: n = 4, k = 8
Output: 0
```

## Approach 1: Recursive Approach (Based on Parent)

This approach leverages the recursive nature of the grammar generation.

1. The first row (n=1) is always "0".
2. Each subsequent row is generated by replacing '0' with "01" and '1' with "10" in the previous row. This means each row has double the length of the previous row.
3. To find the k-th symbol in row n:
    *   If k is in the first half of row n (k <= 2^(n-2)), the symbol is the same as the k-th symbol in row n-1.
    *   If k is in the second half of row n (k > 2^(n-2)), the symbol is the *inverse* (0 becomes 1, and 1 becomes 0) of the (k - 2^(n-2))-th symbol in row n-1.

*   Time Complexity: O(n) in the worst-case scenario where we recurse down to n=1.
*   Space Complexity: O(n) due to the call stack of the recursion.

## Approach 2: Iterative Approach (Bit Manipulation)

This approach uses bit manipulation for a more efficient solution.

1. The value at the k-th position in row n is determined by the parity (even or odd) of the number of set bits (1s) in the binary representation of (k-1).
2. Count the set bits in the binary representation of (k-1).
3. If the count is even, the k-th symbol is 0. If the count is odd, the k-th symbol is 1.

*   Time Complexity: O(log k) or O(n) in the worst case where k is close to 2^(n-1). This is because we iterate through the bits of k.
*   Space Complexity: O(1) as we use constant extra space.