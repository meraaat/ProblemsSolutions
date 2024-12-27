# Problem: Peculiar Array Sum

## Problem Statement

Let's define a peculiar type of array where each element is either an integer or another peculiar array. Assume that a peculiar array is never empty. Write a function that takes a peculiar array as input and finds the sum of its elements. If an element is a peculiar array, you must convert it to its equivalent value before summing it with other elements.

The equivalent value of a peculiar array is the sum of its elements raised to the power of the nesting level (how far nested it is). The outermost array has a nesting level of 1.

## Constraints

*   Peculiar arrays are never empty.
*   Elements within the peculiar arrays are either integers or other peculiar arrays.

## Function Signatures

```javascript
// JavaScript
function peculiarArraySum(arr)
```

### Go
```go
// Go
func peculiarArraySum(arr interface{}) int // Go uses interface{} for mixed types
```

## Example
```Input: [2, 3, [4, 1, 2]]
Output: 54
Explanation: 2 + 3 + (4 + 1 + 2)^2 = 2 + 3 + 7^2 = 2 + 3 + 49 = 54

Input: [1, 2, [7, [3, 4], 2]]
Output: 50
Explanation: 1 + 2 + (7 + (3 + 4)^3 + 2)^2 = 1 + 2 + (7 + 7^3 + 2)^2 = 1 + 2 + (7 + 343 + 2)^2 = 1 + 2 + 352^2 = 1 + 2 + 123904 = 123907

Input: [1]
Output: 1

Input: [[1]]
Output: 1

Input: [[1,2], 3]
Output: 12
Explanation: (1+2)^2 + 3 = 3^2 + 3 = 9+3 = 12
```

## Approach : Recursive Approach

This approach uses recursion to traverse the nested arrays.

1.  **Base Case:** If an element is an integer, return the integer.
2.  **Recursive Step:** If an element is an array:
    *   Recursively calculate the sum of the elements in the array.
    *   Raise the sum to the power of the current nesting level.
    *   Return the result.
3.  The main function initializes the nesting level to 1 and calls the recursive helper function.

*   Time Complexity: O(N), where N is the total number of integers within the nested arrays.
*   Space Complexity: O(D), where D is the maximum nesting depth of the arrays (due to the recursive call stack).
