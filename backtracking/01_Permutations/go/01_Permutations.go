package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			input:    []int{},
			expected: [][]int{},
		},
		{
			input:    []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: %v\n", testCase.input)
		inputCopy := make([]int, len(testCase.input))
		copy(inputCopy, testCase.input)
		result := permute(inputCopy)
		fmt.Printf("Result: %v\n", result)
		fmt.Printf("Expected: %v\n", testCase.expected)
		if !reflect.DeepEqual(result, testCase.expected) {
			fmt.Printf("Test case %d failed\n", i+1)
		} else {
			fmt.Printf("Test case %d passed\n", i+1)
		}
		fmt.Println()
	}
}

// Approach : T O(n!*n) | S O(n)
func permute(nums []int) [][]int {

	result := [][]int{}
	n := len(nums)

	var helper func(i int)
	helper = func(i int) {
		// Base case
		if i == n-1 {
			permutationCopy := make([]int, n)
			copy(permutationCopy, nums)
			result = append(result, permutationCopy)
			return
		}

		// Recursive case
		for j := i; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i] // swap
			helper(i + 1)
			nums[i], nums[j] = nums[j], nums[i] // Backtrack (swap back)
		}
	}

	helper(0)
	return result
}
