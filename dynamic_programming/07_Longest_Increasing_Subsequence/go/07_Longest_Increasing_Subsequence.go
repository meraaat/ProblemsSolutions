package main

import (
	"fmt"
	"math"
	"sort"
)

// Recursive
func lengthOfLISRecursive(nums []int) int {
	n := len(nums)
	var lis func(int) int
	lis = func(i int) int {
		if i == 0 {
			return 1
		}
		maxLen := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxLen = int(math.Max(float64(maxLen), float64(1+lis(j))))
			}
		}
		return maxLen
	}

	maxLength := 0
	for i := 0; i < n; i++ {
		maxLength = int(math.Max(float64(maxLength), float64(lis(i))))
	}
	return maxLength
}

// Memoization
func lengthOfLISMemo(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var lis func(int) int
	lis = func(i int) int {
		if i == 0 {
			return 1
		}
		if memo[i] != -1 {
			return memo[i]
		}
		maxLen := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxLen = int(math.Max(float64(maxLen), float64(1+lis(j))))
			}
		}
		memo[i] = maxLen
		return maxLen
	}

	maxLength := 0
	for i := 0; i < n; i++ {
		maxLength = int(math.Max(float64(maxLength), float64(lis(i))))
	}
	return maxLength
}

// Tabulation
func lengthOfLISTab(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	maxLength := 0
	for _, length := range dp {
		maxLength = int(math.Max(float64(maxLength), float64(length)))
	}
	return maxLength
}

// Patience Sorting with Binary Search
func lengthOfLISPatience(nums []int) int {
	piles := []int{}
	for _, num := range nums {
		i := sort.SearchInts(piles, num)
		if i == len(piles) {
			piles = append(piles, num)
		} else {
			piles[i] = num
		}
	}
	return len(piles)
}

func main() {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 2, 1, 4, 3, 4, 5}, 5},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 1}, 1},
		{[]int{1, 3, 2, 4, 5}, 4},
	}

	lisFuncs := map[string]func([]int) int{
		"Recursive":       lengthOfLISRecursive,
		"Memoization":     lengthOfLISMemo,
		"Tabulation":      lengthOfLISTab,
		"PatienceSorting": lengthOfLISPatience,
	}

	for name, lisFunc := range lisFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: nums=%v\n", i+1, tc.nums)
			res := lisFunc(tc.nums)
			fmt.Printf("Got: %d\n", res)
			fmt.Printf("Expected: %d\n", tc.expect)
			if res != tc.expect {
				fmt.Println("FAIL")
			} else {
				fmt.Println("PASS")
			}
			fmt.Println()
		}
		fmt.Println("--------------------")
	}
}
