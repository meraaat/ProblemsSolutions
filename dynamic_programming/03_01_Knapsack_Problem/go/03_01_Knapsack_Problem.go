package main

import (
	"fmt"
)

// Recursive (Naive)
func knapsackRecursive(W int, wt []int, val []int, n int) int {

	var helper func(int, int) int
	helper = func(index, remWeight int) int {
		// Base case: No items left or no capacity remaining
		if index >= n || remWeight == 0 {
			return 0
		}

		// Recursive case
		exclude := helper(index+1, remWeight)
		include := 0

		if wt[index] <= remWeight {
			include = val[index] + helper(index+1, remWeight-wt[index])
		}

		if include > exclude {
			return include
		}

		return exclude
	}

	return helper(0, W)

}

// Memoization (Top-Down DP)
func knapsackMemo(W int, wt []int, val []int, n int) int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, W+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var helper func(int, int) int
	helper = func(index, remWeight int) int {
		// Base case: No items left or no capacity remaining
		if index == n || remWeight == 0 {
			return 0
		}

		if memo[index][remWeight] != -1 {
			return memo[index][remWeight]
		}

		exclude := helper(index+1, remWeight)

		include := 0
		if wt[index] <= remWeight {
			include = val[index] + helper(index+1, remWeight-wt[index])
		}

		// Store the maximum of including or excluding the item
		memo[index][remWeight] = max(include, exclude)
		return memo[index][remWeight]
	}

	return helper(0, W)
}

// Tabulation (Bottom-Up DP)
func knapsackTab(W int, wt []int, val []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= W; w++ {
			exclude := dp[i-1][w]

			include := 0
			if wt[i-1] <= w {
				include = val[i-1] + dp[i-1][w-wt[i-1]]
			}

			dp[i][w] = max(exclude, include)
		}
	}
	return dp[n][W]
}

// Space Optimized Knapsack
func knapsackSpaceOptimized(W int, wt []int, val []int, n int) int {
	prev := make([]int, W+1) // Stores results for the previous row (i-1)
	curr := make([]int, W+1) // Stores results for the current row (i)

	for i := 1; i <= n; i++ { // Iterate through items
		for w := 1; w <= W; w++ { // Iterate through weights
			exclude := prev[w] // Value if we exclude the current item

			include := 0      // Value if we include the current item
			if wt[i-1] <= w { // Check if the current item's weight fits
				include = val[i-1] + prev[w-wt[i-1]]
			}

			curr[w] = max(exclude, include) // Choose the maximum value
		}
		prev, curr = curr, prev // Swap prev and curr for the next iteration
	}
	return prev[W] // The result is in prev[W] after all iterations
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
		{10, []int{10, 20}, []int{60, 100}, 2, 60},
		{50, []int{10, 20, 30}, []int{60, 100, 120}, 3, 220},
		{10, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5, 10}, // Added test case
	}

	knapsackFuncs := map[string]func(W int, wt []int, val []int, n int) int{
		"Recursive":      knapsackRecursive,
		"Memoization":    knapsackMemo,
		"Tabulation":     knapsackTab,
		"SpaceOptimized": knapsackSpaceOptimized,
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

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
