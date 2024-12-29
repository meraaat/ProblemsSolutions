package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		input []int
	}{
		{
			input: []int{1, 2, 2},
		},
		{
			input: []int{0},
		},
		{
			input: []int{},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: %v\n", testCase.input)
		inputCopy := make([]int, len(testCase.input))
		copy(inputCopy, testCase.input)
		result := subsetsWithDup(inputCopy)
		fmt.Printf("Result: %v\n", result)

		fmt.Println()
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}

	var helper func(i int, subsets []int)

	helper = func(i int, subsets []int) {
		// Base case
		if i == len(nums) {
			subsetsCopy := make([]int, len(subsets))
			copy(subsetsCopy, subsets)
			result = append(result, subsetsCopy)
			return
		}
		// Recursive case
		// Include
		helper(i+1, append(subsets, nums[i]))

		// Exclude
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		helper(i+1, subsets)

	}
	helper(0, []int{})
	return result
}
