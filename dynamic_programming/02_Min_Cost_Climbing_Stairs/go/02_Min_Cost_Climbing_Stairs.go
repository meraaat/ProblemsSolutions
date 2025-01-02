package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []struct {
		cost   []int
		expect int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{1, 2, 3}, 2},
		{[]int{0, 2, 2, 1}, 2},
		{[]int{0, 1, 2, 3, 4, 5}, 6},
	}

	minCostFuncs := map[string]func([]int) int{
		"Recursive":      minCostClimbingStairsRecursive,
		"Memoization":    minCostClimbingStairsMemo,
		"Tabulation":     minCostClimbingStairsTab,
		"SpaceOptimized": minCostClimbingStairsSpace,
	}

	for name, minCostFunc := range minCostFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: minCost(%d)\n", i+1, tc.cost)
			res := minCostFunc(tc.cost)
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

// Recursive (Naive) approach. This has exponential time complexity due to overlapping subproblems.
func minCostClimbingStairsRecursive(cost []int) int {
	n := len(cost)

	
	var minCost func(int) int
	minCost = func(i int) int {
		if i < 0 {
			return 0 
		}
		if i <= 1 {
			return cost[i]
		}
		return cost[i] + int(math.Min(float64(minCost(i-1)), float64(minCost(i-2))))
	}
	return int(math.Min(float64(minCost(n-1)), float64(minCost(n-2))))
}

// Memoization (Top-Down DP). This has O(n) time complexity and O(n) space complexity.
func minCostClimbingStairsMemo(cost []int) int {
	n := len(cost)
	memo := make(map[int]int) 

	var minCost func(int) int
	minCost = func(i int) int {
		if i < 0 {
			return 0 
		}
		if i <= 1 {
			return cost[i]
		}
		if val, ok := memo[i]; ok {
			return val
		}
		memo[i] = cost[i] + int(math.Min(float64(minCost(i-1)), float64(minCost(i-2)))) // Calculate and memoize
		return memo[i]
	}

	return int(math.Min(float64(minCost(n-1)), float64(minCost(n-2))))
}

// Tabulation (Bottom-Up DP). This has O(n) time complexity and O(n) space complexity.
func minCostClimbingStairsTab(cost []int) int {
	n := len(cost)
	dp := make([]int, n) 
	dp[0] = cost[0]     
	dp[1] = cost[1]     
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2]))) // DP relation
	}
	
	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}

// Space Optimized DP. This has O(n) time complexity and O(1) space complexity.
func minCostClimbingStairsSpace(cost []int) int {
	n := len(cost)
	if n <= 1 {
		return 0 // or cost[0] if you consider it as 1 step
	}
	one := cost[0] // Cost to reach the (i-2)th step
	two := cost[1] // Cost to reach the (i-1)th step

	for i := 2; i < n; i++ {
		temp := cost[i] + int(math.Min(float64(one), float64(two))) // Calculate cost to reach the ith step
		one = two                                                   // Update one to the previous two value
		two = temp                                                  // Update two to the current value
	}
	// Return the minimum cost to reach the top (either from n-1 or n-2)
	return int(math.Min(float64(one), float64(two)))
}
