# Problem: N-Queens

## Problem Statement

The n-queens puzzle is the problem of placing *n* queens on an *n* x *n* chessboard such that no two queens attack each other.

Given an integer *n*, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

## Constraints

*   `1 <= n <= 9`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function solveNQueens(n)
```

### Go

```go
// Go
func solveNQueens(n int) [][]string
```

## Example
```
Input: n = 4
Output: [
 [".Q..","...Q","Q...","..Q."],
 ["..Q.","Q...","...Q",".Q.."]
]

Input: n = 1
Output: [["Q"]]

## Approach: Backtracking

Backtracking is the standard approach for solving the N-Queens problem. We try placing queens one row at a time, and for each placement, we check if it's safe (i.e., no other queen attacks it). If it's safe, we recursively try placing the next queen in the next row. If we reach the end of the board (all queens placed), we've found a solution. If a placement is not safe, we backtrack and try a different column.

### Steps:

1.  **Represent the Board:** We can represent the board as a 2D array of strings or characters.

2.  **Base Case:** If we have placed all *n* queens (reached the *n*th row), we have found a solution. Convert the board representation to the desired output format (array of strings) and add it to the result.

3.  **Recursive Step:** For each column in the current row:
    *   **Check if Safe:** Check if placing a queen in the current row and column is safe (no other queen in the same column, diagonal, or anti-diagonal).
    *   **Place the Queen:** If it's safe, place the queen (`'Q'`) in the cell.
    *   **Recursive Call:** Recursively call the function to place the next queen in the next row.
    *   **Backtrack:** If the recursive call doesn't lead to a solution, remove the queen (`'.'`) from the cell (undo the choice).

### Safety Check

To check if a position is safe, we need to check:

1.  **Column:** No other queen in the same column.
2.  **Diagonal:** No other queen on the same diagonal (where `row - col` is constant).
3.  **Anti-Diagonal:** No other queen on the same anti-diagonal (where `row + col` is constant).

## Complexity Analysis

*   **Time Complexity:** The time complexity is difficult to express precisely. In the worst case, we might explore a large portion of the search space, leading to exponential time complexity. However, the constraints and the pruning from the safety checks significantly reduce the actual time taken.

*   **Space Complexity:** O(n^2) to store the board and O(n) for the recursion stack. The output space is proportional to the number of solutions, which can also be significant in some cases.```