package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	testCases := []struct {
		candidates []int
		target     int
		expect     [][]int
	}{
		{
			candidates: []int{3, 5, 2, 1, 3},
			target:     7,
			expect:     [][]int{{1, 3, 3}, {2, 5}},
		},
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expect:     [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expect:     [][]int{{1, 2, 2}, {5}},
		},
		{
			candidates: []int{1, 1, 1, 1, 1},
			target:     3,
			expect:     [][]int{{1, 1, 1}},
		},
		{
			candidates: []int{1, 2},
			target:     4,
			expect:     [][]int{},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: candidates = %v, target = %d\n", testCase.candidates, testCase.target)

		result := combinationSum2(testCase.candidates, testCase.target)

		fmt.Printf("Result: %v\n", result)
		fmt.Printf("Expect: %v\n", testCase.expect)

		for j := range result {
			sort.Ints(result[j])
		}
		for j := range testCase.expect {
			sort.Ints(testCase.expect[j])
		}

		sort.Slice(result, func(i, j int) bool {
			return reflect.DeepEqual(result[i], result[j])
		})
		sort.Slice(testCase.expect, func(i, j int) bool {
			return reflect.DeepEqual(testCase.expect[i], testCase.expect[j])
		})

		if reflect.DeepEqual(result, testCase.expect) {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}

		fmt.Println()
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)

	var backtracks func(index, currentSum int, current []int)
	backtracks = func(index, currentSum int, current []int) {
		// Base case
		if currentSum == target {
			currentCopy := make([]int, len(current))
			copy(currentCopy, current)
			result = append(result, currentCopy)
			return
		}
		if currentSum > target || index > len(candidates)-1 {
			return
		}

		// Recursive case
		seen := make(map[int]bool)
		for i := index; i < len(candidates); i++ {
			if !seen[candidates[i]] {
				seen[candidates[i]] = true
				backtracks(i+1, currentSum+candidates[i], append(current, candidates[i]))
			}
		}

	}

	backtracks(0, 0, []int{})
	return result
}
