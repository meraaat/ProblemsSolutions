package main

import (
	"fmt"
)

// Tabulation (Bottom-Up DP)
func unboundedKnapsack(W int, wt []int, val []int, n int) int {
	dp := make([]int, W+1)

	for w := 0; w <= W; w++ {
		for i := 0; i < n; i++ {
			if wt[i] <= w {
				dp[w] = max(dp[w], val[i]+dp[w-wt[i]])
			}
		}
	}
	return dp[W]
}

// Tabulation (Bottom-Up DP)
func unboundedKnapsack2(W int, wt []int, val []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= W; w++ {
			exclude := dp[i-1][w]

			include := 0
			if wt[i-1] <= w {
				include = val[i-1] + dp[i][w-wt[i-1]]
			}

			dp[i][w] = max(exclude, include)
		}
	}
	return dp[n][W]
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		W      int
		wt     []int
		val    []int
		n      int
		expect int
	}{
		{6, []int{2, 2}, []int{5, 10}, 2, 30},
		{8, []int{1, 3, 4}, []int{1, 4, 5}, 3, 11},
		{4, []int{4, 5, 1}, []int{1, 2, 3}, 3, 12},
		{10, []int{2, 5, 7}, []int{1, 6, 18}, 3, 36},
		{5, []int{1, 2, 3}, []int{10, 15, 20}, 3, 50},
	}

	knapsackFuncs := map[string]func(W int, wt []int, val []int, n int) int{
		"Tabulation":  unboundedKnapsack,
		"Tabulation2": unboundedKnapsack2,
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
