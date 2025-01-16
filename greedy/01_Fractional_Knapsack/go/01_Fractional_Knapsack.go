package main

import (
	"fmt"
	"math"
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

func fractionalKnapsackSimple(W float64, arr [][]int) float64 {
	n := len(arr)
	if n == 0 || W == 0 {
		return 0
	}

	sort.Slice(arr, func(i, j int) bool {
		return float64(arr[j][0])/float64(arr[j][1]) > float64(arr[i][0])/float64(arr[i][1])
	})

	remainingWeight := W
	value := 0.0

	for i := range arr {
		if remainingWeight == 0 {
			break
		}

		weightTaken := math.Min(remainingWeight, float64(arr[i][1]))

		value += (float64(arr[i][0]) / float64(arr[i][1])) * weightTaken

		remainingWeight -= weightTaken
	}

	return value
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
		{30, [][]int{{60, 10}, {100, 20}, {120, 30}}, 160},
		{5, [][]int{{60, 10}, {100, 20}, {120, 30}}, 30},
		{20, [][]int{{60, 10}, {50, 20}, {100, 30}}, 110}, // New test case
	}

	knapsackFuncs := map[string]func(float64, [][]int) float64{
		"Struct Version": fractionalKnapsack,
		"Simple Version": fractionalKnapsackSimple,
	}

	for name, knapsackFunc := range knapsackFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: W=%.0f, items=%v\n", i+1, tc.W, tc.items)
			res := knapsackFunc(tc.W, tc.items)
			fmt.Printf("Got: %.2f\n", res)
			fmt.Printf("Expected: %.2f\n", tc.expect)
			if math.Abs(res-tc.expect) < 0.0001 {
				fmt.Println("PASS")
			} else {
				fmt.Println("FAIL")
			}
			fmt.Println()
		}
		fmt.Println("--------------------")
	}
}
