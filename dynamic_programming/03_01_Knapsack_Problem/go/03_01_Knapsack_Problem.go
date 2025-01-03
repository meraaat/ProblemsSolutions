package main

import (
	"fmt"
	"math"
)

// Recursive (Naive)
func knapsackRecursive(W int, wt []int, val []int, n int) int {
	if n == 0 || W == 0 {
		return 0
	}
	if wt[n-1] > W {
		return knapsackRecursive(W, wt, val, n-1)
	}
	return int(math.Max(float64(val[n-1]+knapsackRecursive(W-wt[n-1], wt, val, n-1)), float64(knapsackRecursive(W, wt, val, n-1))))
}

// Memoization (Top-Down DP)
func knapsackMemo(W int, wt []int, val []int, n int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, W+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var knap func(W int, n int) int
	knap = func(W int, n int) int {
		if n == 0 || W == 0 {
			return 0
		}
		if memo[n][W] != -1 {
			return memo[n][W]
		}
		if wt[n-1] > W {
			memo[n][W] = knap(W, n-1)
			return memo[n][W]
		}
		memo[n][W] = int(math.Max(float64(val[n-1]+knap(W-wt[n-1], n-1)), float64(knap(W, n-1))))
		return memo[n][W]
	}
	return knap(W, n)
}

// Tabulation (Bottom-Up DP)
func knapsackTab(W int, wt []int, val []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if wt[i-1] <= w {
				dp[i][w] = int(math.Max(float64(val[i-1]+dp[i-1][w-wt[i-1]]), float64(dp[i-1][w])))
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]
}

func main() {
	testCases := []struct {
		W      int
		wt     []int
		val    []int
		n      int
		expect int
	}{
		{8, []int{4, 5, 2}, []int{3, 9, 5}, 3, 14},
		{4, []int{4, 5, 1}, []int{1, 2, 3}, 3, 3},
		{10, []int{10, 20}, []int{60, 100}, 2, 100},
		{50, []int{10, 20, 30}, []int{60, 100, 120}, 3, 220},
		{10, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5, 10}, // Added test case
	}

	knapsackFuncs := map[string]func(W int, wt []int, val []int, n int) int{
		"Recursive":   knapsackRecursive,
		"Memoization": knapsackMemo,
		"Tabulation":  knapsackTab,
	}

	for name, knapsackFunc := range knapsackFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: W=%d, wt=%v, val=%v, n=%d\n", i+1, tc.W, tc.wt, tc.val, tc.n)
			res := knapsackFunc(tc.W, tc.wt, tc.val, tc.n)
			fmt.Printf("Got: %d\n", res)
			fmt.Printf("Expected: %d\n", tc.expect)
			if res != tc.expect {
				fmt.Println("FAIL")
			} else {
				fmt.Println("PASS")
			}
			fmt.Println()
		}
		fmt.Println("--------------------")
	}
}
