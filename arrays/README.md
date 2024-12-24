# Coding Interview Problems Collection

A curated collection of coding interview problems with detailed solutions in JavaScript and Go.

## Problem 1: Sorted Squared Array

### Problem Statement
Write a function that takes an array of integers in which each subsequent value is not less than the previous value (non-decreasing array). The function should return a new array containing the squares of each number sorted in ascending order.

### Constraints
- Input array is in non-decreasing order
- Numbers can be negative, zero, or positive
- Numbers are not necessarily distinct
- Empty array is a valid input
- Numbers are integers

### Function Signatures

```javascript
// JavaScript
function sortedSquaredArray(array)
```

```go
// Go
func SortedSquaredArray(array []int) []int
```

### Example
```
Input: [-4, -1, 0, 3, 10]
Output: [0, 1, 9, 16, 100]

Input: [-7, -3, 2, 3, 11]
Output: [4, 9, 9, 49, 121]

Input: []
Output: []
```

### Approach 1: Square and Sort
1. Square each number in the array
2. Sort the resulting array
- Time Complexity: O(n log n)
- Space Complexity: O(n)

### Approach 2: Two Pointers
1. Create result array of same length
2. Use two pointers (left and right)
3. Compare absolute values and place squares accordingly
- Time Complexity: O(n)
- Space Complexity: O(n)