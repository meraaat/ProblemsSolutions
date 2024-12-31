package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		n      int
		expect [][]string
	}{
		{
			n: 4,
			expect: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			n:      1,
			expect: [][]string{{"Q"}},
		},
		{
			n:      2,
			expect: [][]string{},
		},
		{
			n:      3,
			expect: [][]string{},
		},
		{
			n: 5,
			expect: [][]string{
				{".Q...", "...Q.", ".....", ".Q...", "...Q."},
				{".Q...", ".....", "..Q..", "...Q.", "Q...."},
				{"..Q..", ".....", ".Q...", "...Q.", "Q...."},
				{"..Q..", "...Q.", "Q....", ".....", ".Q..."},
				{"...Q.", "Q....", "..Q..", ".....", ".Q..."},
				{"...Q.", ".Q...", ".....", "Q....", "..Q.."},
				{"Q....", "...Q.", ".Q...", ".....", "..Q.."},
				{"Q....", "..Q..", ".....", ".Q...", "...Q."},
			},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: n = %d\n", testCase.n)

		result := solveNQueens(testCase.n)

		fmt.Printf("Result:\n")
		for _, row := range result {
			fmt.Println(row)
		}
		fmt.Printf("Expect:\n")
		for _, row := range testCase.expect {
			fmt.Println(row)
		}

		if reflect.DeepEqual(result, testCase.expect) {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			current := make([]string, n)
			for i := range board {
				current[i] = string(board[i])
			}
			result = append(result, current)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)
	return result
}

func isValid(board [][]byte, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check top-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check top-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
