# Problem: Sudoku Solver

## Problem Statement

Write a program to solve a Sudoku puzzle by filling the empty cells.

A Sudoku solution must satisfy all of the following rules:

1.  Each of the digits `1-9` must occur exactly once in each row.
2.  Each of the digits `1-9` must occur exactly once in each column.
3.  Each of the digits `1-9` must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

The `'.'` character indicates empty cells.

## Constraints

*   The input `board` is a 9x9 array of characters.
*   The `board` may contain digits `1-9` and the `'.'` character.
*   It is guaranteed that the input `board` has a unique solution.

## Function Signatures

### JavaScript

```javascript
// JavaScript
function solveSudoku(board)
```
### Go
```go
// Go
func solveSudoku(board [][]byte)
```

## Example

```
Input:
board = [
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]

Output:
board = [
  ["5","3","4","6","7","8","9","1","2"],
  ["6","7","2","1","9","5","3","4","8"],
  ["1","9","8","3","4","2","5","6","7"],
  ["8","5","9","7","6","1","4","2","3"],
  ["4","2","6","8","5","3","7","9","1"],
  ["7","1","3","9","2","4","8","5","6"],
  ["9","6","1","5","3","7","2","8","4"],
  ["2","8","7","4","1","9","6","3","5"],
  ["3","4","5","2","8","6","1","7","9"]
]

```
## Approach: Backtracking

Backtracking is a natural fit for solving Sudoku puzzles. We try filling each empty cell with a digit from `1` to `9` and recursively check if the resulting board is valid. If it's not valid, we backtrack and try a different digit.

### Steps:

1.  **Find an Empty Cell:** Iterate through the board to find the next empty cell (`.`).

2.  **Base Case:** If no empty cell is found, the puzzle is solved.

3.  **Recursive Step:** For each digit from `1` to `9`:
    *   **Place the Digit:** Place the digit in the empty cell.
    *   **Check Validity:** Check if the current board state is valid (respects Sudoku rules).
    *   **Recursive Call:** If the board is valid, recursively call the solver function to fill the next empty cell.
    *   **Backtrack:** If the recursive call returns `false` (meaning no solution was found from that point), undo the placement of the digit (set the cell back to `'.'`).

### Validity Check

The validity check involves three separate checks:

1.  **Row Check:** Check if the digit already exists in the current row.
2.  **Column Check:** Check if the digit already exists in the current column.
3.  **3x3 Sub-box Check:** Check if the digit already exists in the 3x3 sub-box containing the current cell.

## Complexity Analysis

*   **Time Complexity:** The worst-case time complexity is difficult to precisely determine, as it depends on the initial state of the board. In the worst case, we might try all possible combinations of digits in empty cells, which would be very high. However, Sudoku puzzles are designed to be solvable with reasonable backtracking, so the practical time complexity is much lower.

*   **Space Complexity:** O(1) as the board size is fixed (9x9). The recursion stack depth can be at most 81 (the number of cells), but this is still considered constant space.