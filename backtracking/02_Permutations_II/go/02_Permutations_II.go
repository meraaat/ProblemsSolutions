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
			input: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
			},
		},
		{
			input:    []int{1, 1, 1},
			expected: [][]int{{1, 1, 1}},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: %v\n", testCase.input)
		inputCopy := make([]int, len(testCase.input))
		copy(inputCopy, testCase.input)
		result := permuteUnique(inputCopy)
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
func permuteUnique(nums []int) [][]int {

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

		seen := make(map[int]bool)
		// Recursive case
		for j := i; j < n; j++ {
			if !seen[nums[j]] {

				seen[nums[j]] = true
				nums[i], nums[j] = nums[j], nums[i] // swap
				helper(i + 1)
				nums[i], nums[j] = nums[j], nums[i] // Backtrack (swap back)
			}
		}
	}

	helper(0)
	return result
}
