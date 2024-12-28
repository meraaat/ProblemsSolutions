package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		input []int
	}{
		{
			input: []int{1, 2, 3},
		},
		{
			input: []int{0},
		},
		{
			input: []int{},
		},
		{
			input: []int{1, 2},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: %v\n", testCase.input)
		inputCopy := make([]int, len(testCase.input))
		copy(inputCopy, testCase.input)
		result := subsets(inputCopy)
		fmt.Printf("Result: %v\n", result)

		fmt.Println()
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}

	var helper func(nums []int, i int, subsets []int)

	helper = func(nums []int, i int, subsets []int) {
		// Base case
		if i == len(nums) {
			subsetsCopy := make([]int, len(subsets))
			copy(subsetsCopy, subsets)
			result = append(result, subsetsCopy)
			return
		}

		// Recursive case

		// Don't Add
		helper(nums, i+1, subsets)

		// Add
		helper(nums, i+1, append(subsets, nums[i]))

	}
	helper(nums, 0, []int{})
	return result
}
