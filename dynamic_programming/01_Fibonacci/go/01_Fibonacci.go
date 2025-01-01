package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		n      int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
	}

	fibFuncs := map[string]func(int) int{
		"Recursive":      fibRecursive,
		"Memoization":    fibMemoization,
		"Tabulation":     fibTabulation,
		"SpaceOptimized": fibSpaceOptimized,
	}

	for name, fibFunc := range fibFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: fib(%d)\n", i+1, tc.n)
			res := fibFunc(tc.n)
			fmt.Printf("Got: %d\n", res)
			fmt.Printf("Expected: %d\n", tc.expect)
			if res != tc.expect {
				fmt.Println("FAIL")
			} else {
				fmt.Println("PASS")
			}
			fmt.Println()
		}
		fmt.Println("--------------------") // Separator between function tests
	}
}

// 1. Recursion
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// 2. Memoization
func fibMemoization(n int) int {
	memo := make(map[int]int)
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if val, ok := memo[n]; ok {
			return val
		}
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}
	return fib(n)
}

// 3. Tabulation
func fibTabulation(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 4. Space Optimization
func fibSpaceOptimized(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
