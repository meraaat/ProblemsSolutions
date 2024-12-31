function solveNQueens(n) {
  const result = [];
  const board = Array(n)
    .fill(null)
    .map(() => Array(n).fill("."));

  const isSafe = (row, col) => {
    for (let i = 0; i < n; i++) {
      if (board[row][i] === "Q" || board[i][col] === "Q") {
        return false;
      }
      if (row - i >= 0 && col - i >= 0 && board[row - i][col - i] === "Q") {
        return false;
      }
      if (row - i >= 0 && col + i < n && board[row - i][col + i] === "Q") {
        return false;
      }
    }
    return true;
  };

  const solve = (row) => {
    if (row === n) {
      result.push(board.map((r) => r.join("")));
      return;
    }

    for (let col = 0; col < n; col++) {
      if (isSafe(row, col)) {
        board[row][col] = "Q";
        solve(row + 1);
        board[row][col] = "."; // Backtrack
      }
    }
  };

  solve(0);
  return result;
}

const testCases = [
  {
    n: 4,
    expected: [
      [".Q..", "...Q", "Q...", "..Q."],
      ["..Q.", "Q...", "...Q", ".Q.."],
    ],
  },
  {
    n: 1,
    expected: [["Q"]],
  },
  {
    n: 2,
    expected: [],
  },
  {
    n: 3,
    expected: [],
  },
  {
    n: 5,
    expected: [
      [".Q...", "...Q.", ".....", ".Q...", "...Q."],
      [".Q...", ".....", "..Q..", "...Q.", "Q...."],
      ["..Q..", ".....", ".Q...", "...Q.", "Q...."],
      ["..Q..", "...Q.", "Q....", ".....", ".Q..."],
      ["...Q.", "Q....", "..Q..", ".....", ".Q..."],
      ["...Q.", ".Q...", ".....", "Q....", "..Q.."],
      ["Q....", "...Q.", ".Q...", ".....", "..Q.."],
      ["Q....", "..Q..", ".....", ".Q...", "...Q."],
    ],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(`Input: n = ${testCase.n}`);
    const result = solveNQueens(testCase.n);
    console.log("Result:");
    result.forEach((row) => console.log(row));
    console.log("Expected:");
    testCase.expected.forEach((row) => console.log(row));
    if (JSON.stringify(result) === JSON.stringify(testCase.expected)) {
      console.log("PASS");
    } else {
      console.log("FAIL");
    }
    console.log();
  });
})();

var solveNQueens = function (n) {
  let res = [];
  // Create initial empty board
  let board = new Array(n).fill().map(() => new Array(n).fill("."));

  function isValid(row, col, board) {
    // Check column for other queens
    for (let x = 0; x < row; x++) {
      if (board[x][col] === "Q") return false;
    }

    // Check top-left diagonal
    for (let r = row, c = col; r >= 0 && c >= 0; r--, c--) {
      if (board[r][c] === "Q") return false;
    }

    // Check top-right diagonal
    for (let r = row, c = col; r >= 0 && c < n; r--, c++) {
      if (board[r][c] === "Q") return false;
    }

    return true;
  }

  function convertBoard(board) {
    return board.map((row) => row.join(""));
  }

  function positionNextQueen(board, row) {
    // Base case: all queens are placed
    if (row === n) {
      res.push(convertBoard(board));
      return;
    }

    // Try placing a queen in each column of the current row
    for (let col = 0; col < n; col++) {
      if (isValid(row, col, board)) {
        board[row][col] = "Q"; // Place the queen
        positionNextQueen(board, row + 1); // Move to next row
        board[row][col] = "."; // Backtrack: remove queen
      }
    }
  }

  positionNextQueen(board, 0); // Start from first row
  return res;
};
