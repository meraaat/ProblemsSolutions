const testCases = [
  {
    input: [
      ["5", "3", ".", ".", "7", ".", ".", ".", "."],
      ["6", ".", ".", "1", "9", "5", ".", ".", "."],
      [".", "9", "8", ".", ".", ".", ".", "6", "."],
      ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
      ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
      ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
      [".", "6", ".", ".", ".", ".", "2", "8", "."],
      [".", ".", ".", "4", "1", "9", ".", ".", "5"],
      [".", ".", ".", ".", "8", ".", ".", "7", "9"],
    ],
    expected: [
      ["5", "3", "4", "6", "7", "8", "9", "1", "2"],
      ["6", "7", "2", "1", "9", "5", "3", "4", "8"],
      ["1", "9", "8", "3", "4", "2", "5", "6", "7"],
      ["8", "5", "9", "7", "6", "1", "4", "2", "3"],
      ["4", "2", "6", "8", "5", "3", "7", "9", "1"],
      ["7", "1", "3", "9", "2", "4", "8", "5", "6"],
      ["9", "6", "1", "5", "3", "7", "2", "8", "4"],
      ["2", "8", "7", "4", "1", "9", "6", "3", "5"],
      ["3", "4", "5", "2", "8", "6", "1", "7", "9"],
    ],
  },
  {
    input: [
      [".", ".", "9", "7", "4", "8", ".", ".", "."],
      ["7", ".", ".", ".", ".", ".", ".", ".", "."],
      [".", "2", ".", "1", ".", "9", ".", ".", "."],
      [".", ".", "7", ".", ".", ".", "2", "4", "."],
      [".", "6", "4", ".", "1", ".", "5", "9", "."],
      [".", "9", "8", ".", ".", ".", "3", ".", "."],
      [".", ".", ".", "8", ".", "3", ".", "2", "."],
      [".", ".", ".", ".", ".", ".", ".", ".", "6"],
      [".", ".", ".", "2", "7", "5", "9", ".", "."],
    ],
    expected: [
      ["5", "1", "9", "7", "4", "8", "6", "3", "2"],
      ["7", "8", "3", "9", "5", "2", "1", "4", "6"],
      ["4", "2", "6", "1", "3", "9", "8", "7", "5"],
      ["3", "5", "7", "6", "9", "4", "2", "1", "8"],
      ["2", "6", "4", "3", "1", "8", "5", "9", "7"],
      ["1", "9", "8", "5", "2", "7", "3", "6", "4"],
      ["9", "7", "5", "8", "6", "3", "4", "2", "1"],
      ["8", "4", "2", "5", "9", "1", "7", "3", "6"],
      ["6", "3", "1", "2", "7", "5", "9", "8", "4"],
    ],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log("Input:");
    testCase.input.forEach((row) => console.log(row.join("")));

    // Create a deep copy of the input
    const inputCopy = testCase.input.map((row) => [...row]);

    solveSudoku(inputCopy);

    console.log("Result:");
    inputCopy.forEach((row) => console.log(row.join("")));

    console.log("Expected:");
    testCase.expected.forEach((row) => console.log(row.join("")));

    if (JSON.stringify(inputCopy) === JSON.stringify(testCase.expected)) {
      console.log("PASS");
    } else {
      console.log("FAIL");
    }
    console.log();
  });
})();

function solveSudoku(board) {
  function isValid(board, row, col, num) {
    // Check row and column
    for (let x = 0; x < 9; x++) {
      if (board[x][col] === num || board[row][x] === num) {
        return false;
      }
    }

    // Check 3x3 box
    let r = 3 * Math.floor(row / 3);
    let c = 3 * Math.floor(col / 3);
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        if (board[r + i][c + j] === num) {
          return false;
        }
      }
    }

    return true;
  }

  function helper(board) {
    for (let row = 0; row < 9; row++) {
      for (let col = 0; col < 9; col++) {
        if (board[row][col] === ".") {
          for (let num = 1; num <= 9; num++) {
            let char = num.toString();
            if (isValid(board, row, col, char)) {
              board[row][col] = char;
              if (helper(board)) {
                return true;
              }
              board[row][col] = ".";
            }
          }
          return false;
        }
      }
    }
    return true;
  }
  helper(board);
}
