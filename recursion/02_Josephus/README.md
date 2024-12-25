# Problem: Josephus Problem (Circular Game)

## Problem Statement

There are `n` friends playing a game. They are sitting in a circle, numbered from 1 to `n` in clockwise order. Moving clockwise from the `i`-th friend brings you to the `(i+1)`-th friend for `1 <= i < n`, and moving clockwise from the `n`-th friend brings you to the 1st friend.

The rules are:

1.  Start at the 1st friend.
2.  Count the next `k` friends clockwise (including the starting friend). The counting wraps around the circle.
3.  The last friend counted leaves the circle.
4.  If more than one friend remains, repeat from step 2, starting from the friend immediately clockwise of the one who just lost.
5.  The last remaining friend wins.

Given `n` (number of friends) and `k` (counting step), return the winner's number.

## Constraints

*   `1 <= n <= 500` (Example constraint, adjust as needed)
*   `1 <= k <= n` (Example constraint, adjust as needed)

## Function Signatures

### JavaScript
```javascript
// JavaScript
function findTheWinner(n, k)
```

### Go
```go
// Go
func findTheWinner(n int, k int) int
```

## Example
```
Input: n = 4, k = 2
Output: 3

Explanation:
1. Start at 1. Count 1, 2. Friend 2 leaves.
2. Start at 3. Count 3, 4. Friend 4 leaves.
3. Start at 1. Count 1, 3. Friend 1 leaves.
4. Friend 3 wins.

Input: n = 5, k = 2
Output: 3

Explanation:
1. Start at 1. Count 1, 2. Friend 2 leaves.
2. Start at 3. Count 3, 4. Friend 4 leaves.
3. Start at 5. Count 5, 1. Friend 1 leaves.
4. Start at 3. Count 3, 5. Friend 5 leaves.
5. Friend 3 wins.

```

## Approach 1: Recursive Approach

This approach uses recursion to simulate the game.

1.  **Base Case:** If `n == 1`, there's only one friend left, so they are the winner. Return 1.
2.  **Recursive Step:**
    *   The first friend to be removed in the initial round is `(k - 1) % n`. This is because we start at index 0 and count k elements, and we use modulo to wrap around.
    *   The remaining `n-1` friends form a smaller circle. We recursively call `findTheWinner(n-1, k)`.
    *   The result of the recursive call is the winner's position in the smaller circle (indexed from 1 to n-1). To map this back to the original circle (indexed 1 to n), we need to shift the result by `k` and take the modulo `n`. The formula for this mapping is `(findTheWinner(n-1, k) + k - 1) % n + 1`.

*   Time Complexity: O(n) as we recurse n times in the worst case.
*   Space Complexity: O(n) due to the recursion call stack.

## Approach 2: Iterative Approach (Optimized)

This approach uses a mathematical formula to directly calculate the survivor's position without simulating the game step by step.

1.  Initialize `survivor = 0`. This represents the 0-based index of the survivor.
2.  Iterate from `i = 2` to `n` (inclusive). `i` represents the current number of people in the circle.
3.  In each iteration, update the survivor's position using the formula: `survivor = (survivor + k) % i`.
4.  After the loop, `survivor` holds the 0-based index of the final survivor. Add 1 to convert it to a 1-based index.

*   Time Complexity: O(n) as the loop iterates n-1 times.
*   Space Complexity: O(1) as we use only constant extra space.