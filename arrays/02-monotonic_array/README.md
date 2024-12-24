# Problem 2: Monotonic Array

## Problem Statement

Write a function that takes an array of integers and determines whether the array is monotonic. An array is monotonic if it is either monotone increasing or monotone decreasing.

*   **Monotone Increasing:** All elements from left to right are non-decreasing (i.e., each element is greater than or equal to the previous one).
*   **Monotone Decreasing:** All elements from left to right are non-increasing (i.e., each element is less than or equal to the previous one).

The function should return `true` if the array is monotonic, and `false` otherwise.

## Constraints

*   Numbers can be negative, zero, or positive.
*   Numbers are not necessarily distinct.
*   Empty array is a valid input.
*   Numbers are integers.

## Function Signatures


### JavaScript
```javascript
// JavaScript
function isMonotonic(array)
```

### Go
```go
// Go
func isMonotonic(array []int) bool
```
## Example
```
Input: [1, 2, 3, 5, 13]
Output: true

Input: [3, 2, 1]
Output: true

Input: [1, 2, 1]
Output: false

Input: []
Output: true

Input: [1]
Output: true

Input: [1,1,2,3,3,4,4]
Output: true

Input: [6,5,4,4,3,2,1]
Output: true

Input: [1,3,2]
Output: false

Input: [1,2,4,2]
Output: false
```

## Approach : Single Pass Check

1. Initialize `isIncreasing` and `isDecreasing` to `true`.
2. Iterate through `array` (starting from the second element):
    *   If `array[i] > array[i-1]`, set `isDecreasing = false`.
    *   If `array[i] < array[i-1]`, set `isIncreasing = false`.
3. Return `isIncreasing || isDecreasing`.

*   Time Complexity: O(n)
*   Space Complexity: O(1)
