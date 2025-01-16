package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Profit float64
	Weight float64
	Ratio  float64
}

func fractionalKnapsack(W float64, items [][]int) float64 {
	n := len(items)
	if n == 0 || W == 0 {
		return 0
	}

	itemStructs := make([]Item, n)
	for i := 0; i < n; i++ {
		itemStructs[i] = Item{Profit: float64(items[i][0]), Weight: float64(items[i][1])}
		itemStructs[i].Ratio = itemStructs[i].Profit / itemStructs[i].Weight
	}

	sort.Slice(itemStructs, func(i, j int) bool {
		return itemStructs[i].Ratio > itemStructs[j].Ratio
	})

	totalProfit := 0.0
	remainingWeight := W

	for _, item := range itemStructs {
		if item.Weight <= remainingWeight {
			totalProfit += item.Profit
			remainingWeight -= item.Weight
		} else {
			totalProfit += (remainingWeight / item.Weight) * item.Profit
			remainingWeight = 0
			break
		}
	}

	return totalProfit
}

func main() {
	testCases := []struct {
		W      float64
		items  [][]int
		expect float64
	}{
		{25, [][]int{{70, 10}, {90, 20}, {150, 30}}, 137.5},
		{45, [][]int{{70, 10}, {90, 20}, {150, 30}}, 242.5},
		{10, [][]int{{10, 2}, {5, 3}, {15, 5}, {7, 7}, {6, 1}, {18, 6}}, 43},
		{10, [][]int{}, 0},
		{0, [][]int{{10, 5}}, 0},
		{50, [][]int{{60, 10}, {100, 20}, {120, 30}}, 240},
		{10, [][]int{{10, 5}, {1, 1}}, 11},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d: W=%.0f, items=%v\n", i+1, tc.W, tc.items)
		res := fractionalKnapsack(tc.W, tc.items)
		fmt.Printf("Got: %.2f\n", res)
		fmt.Printf("Expected: %.2f\n", tc.expect)
		if res != tc.expect {
			fmt.Println("FAIL")
		} else {
			fmt.Println("PASS")
		}
		fmt.Println()
	}
}
