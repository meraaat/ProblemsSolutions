package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		input  [][]byte
		expect [][]byte
	}{
		{
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expect: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
		{
			input: [][]byte{
				{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
				{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
				{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
				{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
				{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
				{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
			},
			expect: [][]byte{
				{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
				{'7', '8', '3', '9', '5', '2', '1', '4', '6'},
				{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
				{'3', '5', '7', '6', '9', '4', '2', '1', '8'},
				{'2', '6', '4', '3', '1', '8', '5', '9', '7'},
				{'1', '9', '8', '5', '2', '7', '3', '6', '4'},
				{'9', '7', '5', '8', '6', '3', '4', '2', '1'},
				{'8', '4', '2', '5', '9', '1', '7', '3', '6'},
				{'6', '3', '1', '2', '7', '5', '9', '8', '4'},
			},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Println("Input:")
		for _, row := range testCase.input {
			fmt.Println(string(row))
		}
		inputCopy := make([][]byte, 9)
		for j := 0; j < 9; j++ {
			inputCopy[j] = make([]byte, 9)
			copy(inputCopy[j], testCase.input[j])
		}
		solveSudoku(inputCopy)
		fmt.Println("Result:")
		for _, row := range inputCopy {
			fmt.Println(string(row))
		}
		fmt.Println("Expected:")
		for _, row := range testCase.expect {
			fmt.Println(string(row))
		}

		if reflect.DeepEqual(inputCopy, testCase.expect) {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}

		fmt.Println()
	}
}

func solveSudoku(board [][]byte) {
    var solve func(row, col int) bool
    solve = func(row, col int) bool {
        if row == 9 {
            return true // Reached the end of the board
        }
        if col == 9 {
            return solve(row+1, 0) // Move to the next row
        }

        if board[row][col] != '.' {
            return solve(row, col+1) // Skip filled cells
        }

        for num := byte('1'); num <= byte('9'); num++ {
            if isValid(board, row, col, num) {
                board[row][col] = num
                if solve(row, col+1) {
                    return true // Solution found!
                }
                board[row][col] = '.' // Backtrack: Reset the cell
            }
        }
        return false // No valid number found for this cell
    }
    solve(0, 0) // Start solving from top-left (0, 0)
}

func isValid(board [][]byte, row, col int, num byte) bool {
    for i := 0; i < 9; i++ {
        if board[row][i] == num || // Check row
            board[i][col] == num || // Check column
            board[row/3*3+i/3][col/3*3+i%3] == num { // Check 3x3 sub-box
            return false // Number already exists in row, column, or sub-box
        }
    }
    return true // Number is valid in this cell
}

