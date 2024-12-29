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
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expect:     [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expect:     [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			expect:     [][]int{},
		},
		{
			candidates: []int{1},
			target:     1,
			expect:     [][]int{{1}},
		},
		{
			candidates: []int{1, 2},
			target:     4,
			expect:     [][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: candidates = %v, target = %d\n", testCase.candidates, testCase.target)

		result := combinationSum(testCase.candidates, testCase.target)

		fmt.Printf("Result: %v\n", result)
		fmt.Printf("Expect: %v\n", testCase.expect)

		// Sort inner slices and outer slice for comparison
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

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var combinationSumRecursive func(start int, current []int, currentSum int)
	combinationSumRecursive = func(start int, current []int, currentSum int) {
		// Base cases
		if currentSum > target {
			return
		}

		if currentSum == target {
			currentCopy := make([]int, len(current))
			copy(currentCopy, current)
			result = append(result, currentCopy)
			return
		}

		// Recursive case
		for i := start; i < len(candidates); i++ {
			combinationSumRecursive(i, append(current, candidates[i]), currentSum+candidates[i])
		}
	}

	combinationSumRecursive(0, []int{}, 0)
	return result
}
