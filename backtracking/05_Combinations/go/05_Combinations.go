package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	testCases := []struct {
		n      int
		k      int
		expect [][]int
	}{
		{
			n:      4,
			k:      2,
			expect: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:      1,
			k:      1,
			expect: [][]int{{1}},
		},
		{
			n:      5,
			k:      3,
			expect: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
		{
			n:      2,
			k:      1,
			expect: [][]int{{1}, {2}},
		},
		{
			n:      2,
			k:      2,
			expect: [][]int{{1, 2}},
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Input: n = %d, k = %d\n", testCase.n, testCase.k)

		result := combine(testCase.n, testCase.k)

		fmt.Printf("Result: %v\n", result)
		fmt.Printf("Expect: %v\n", testCase.expect)

		// Sort the inner slices for comparison
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

func combine(n int, k int) [][]int {
	result := [][]int{}

	var helper func(start int, current []int)
	helper = func(start int, current []int) {
		// Base case
		if len(current) == k {
			currentCopy := make([]int, len(current))
			copy(currentCopy, current)
			result = append(result, currentCopy)
			return
		}

		// Recursive case
		need := k - len(current)
		for i := start; i <= n-need+1; i++ {
			helper(i+1, append(current, i))

		}
	}

	helper(1, []int{})
	return result
}
