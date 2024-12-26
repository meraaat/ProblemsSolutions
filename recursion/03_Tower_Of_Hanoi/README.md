# Problem: Tower of Hanoi

## Problem Statement

Given N disks and three rods (numbered 1, 2, and 3), move all disks from rod 1 to rod 3 following these rules:

1. Only one disk can be moved at a time
2. A larger disk cannot be placed on top of a smaller disk
3. Disks can only be moved between rods
4. Initially, all disks are on rod 1 in ascending order (smallest on top)

The goal is to move the entire stack to rod 3, maintaining the ascending order, while printing each move.

## Constraints

* `1 <= N <= 15` (practical constraint due to exponential growth)
* Disks are numbered 1 to N (1 being the smallest)
* All disks must start on rod 1
* Final state must have all disks on rod 3
* Each move must be valid according to size constraints

## Function Signatures

### JavaScript
```javascript
function toh(N, from, to, aux) {
    // Returns number of moves
}
```

### Go
```go
func toh(N int, from int, to int, aux int) int {
    // Returns number of moves
}
```

## Example

```
Input: N = 2
Output: 
Move disk 1 from rod 1 to rod 2
Move disk 2 from rod 1 to rod 3
Move disk 1 from rod 2 to rod 3

Input: N = 3
Output:
Move disk 1 from rod 1 to rod 3
Move disk 2 from rod 1 to rod 2
Move disk 1 from rod 3 to rod 2
Move disk 3 from rod 1 to rod 3
Move disk 1 from rod 2 to rod 1
Move disk 2 from rod 2 to rod 3
Move disk 1 from rod 1 to rod 3
```

## Approach 1: Recursive Solution

The Tower of Hanoi can be solved elegantly using recursion by breaking it down into smaller subproblems:

1. **Base Case:** If N = 1, simply move the disk from source to destination
2. **Recursive Steps:**
   * Move N-1 disks from source to auxiliary rod
   * Move the largest disk from source to destination
   * Move the N-1 disks from auxiliary to destination rod

* Time Complexity: O(2^n) - each disk requires 2 recursive calls
* Space Complexity: O(n) for the recursion stack

## Approach 2: Iterative Solution

For N disks, the optimal solution requires 2^N - 1 moves. We can solve this iteratively using these patterns:

1. For odd N: Counter-clockwise moves (1→3, 3→2, 2→1)
2. For even N: Clockwise moves (1→2, 2→3, 3→1)

* Time Complexity: O(2^n) - optimal number of moves required
* Space Complexity: O(1) - constant extra space

## Mathematical Properties

1. Minimum moves required: 2^N - 1
2. Disk movement pattern:
   * Smallest disk moves every 2 steps
   * Disk k moves every 2^k steps
3. Each disk moves in a consistent direction throughout the solution