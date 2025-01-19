package main

import (
	"fmt"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	n := len(costs) / 2

	// Sort by the difference between cost to A and cost to B
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	minCost := 0
	for i := 0; i < n; i++ {
		minCost += costs[i][0]   // Cost to city A for the first n people
		minCost += costs[i+n][1] // Cost to city B for the remaining n people
	}

	return minCost
}

func main() {
	testCases := []struct {
		costs  [][]int
		expect int
	}{
		{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
		{[][]int{{5, 20}, {30, 100}, {400, 30}, {50, 10}}, 75},
		{[][]int{{20, 700}, {400, 50}, {900, 600}, {200, 150}, {800, 100}, {500, 450}}, 1470},
		{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
		{[][]int{{10, 90}, {80, 20}, {30, 70}, {60, 40}}, 110},
		{[][]int{{10, 10}, {10, 10}}, 20},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d: costs=%v\n", i+1, tc.costs)
		res := twoCitySchedCost(tc.costs)
		fmt.Printf("Got: %d\n", res)
		fmt.Printf("Expected: %d\n", tc.expect)
		if res != tc.expect {
			fmt.Println("FAIL")
		} else {
			fmt.Println("PASS")
		}
		fmt.Println()
	}
}
