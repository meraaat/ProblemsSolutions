package main

import (
	"fmt"
	"math"
)

// Recursive
func longestCommonSubsequenceRecursive(text1 string, text2 string) int {
	var lcs func(i, j int) int
	lcs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if text1[i] == text2[j] {
			return 1 + lcs(i-1, j-1)
		}
		return max(lcs(i-1, j), lcs(i, j-1))
	}
	return lcs(len(text1)-1, len(text2)-1)
}

func longestCommonSubsequenceRecursive2(text1, text2 string) int {

	n := len(text1)
	m := len(text2)

	var helper func(int, int) int
	helper = func(index1, index2 int) int {

		// Base case
		if index1 >= n || index2 >= m {
			return 0
		}

		// Recursive case
		if text1[index1] == text2[index2] {
			return 1 + helper(index1+1, index2+1)
		}

		return max(helper(index1+1, index2), helper(index1, index2+1))
	}

	return helper(0, 0)
}

// Memoization
func longestCommonSubsequenceMemo(text1 string, text2 string) int {
	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var lcs func(i, j int) int
	lcs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if memo[i+1][j+1] != -1 {
			return memo[i+1][j+1]
		}
		if text1[i] == text2[j] {
			memo[i+1][j+1] = 1 + lcs(i-1, j-1)
			return memo[i+1][j+1]
		}
		memo[i+1][j+1] = int(math.Max(float64(lcs(i-1, j)), float64(lcs(i, j-1))))
		return memo[i+1][j+1]
	}
	return lcs(len(text1)-1, len(text2)-1)
}

func longestCommonSubsequenceMemo2(text1, text2 string) int {
	n := len(text1)
	m := len(text2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var helper func(int, int) int
	helper = func(index1, index2 int) int {

		// Base case
		if index1 >= n || index2 >= m {
			return 0
		}

		if memo[index1][index2] != -1 {
			return memo[index1][index2]
		}

		// Recursive case
		if text1[index1] == text2[index2] {
			memo[index1][index2] = 1 + helper(index1+1, index2+1)
		} else {
			memo[index1][index2] = max(helper(index1+1, index2), helper(index1, index2+1))
		}

		return memo[index1][index2]
	}

	return helper(0, 0)
}

// Tabulation
func longestCommonSubsequenceTab(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

// Space Optimized
func longestCommonSubsequenceSpace(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	if m < n {
		text1, text2 = text2, text1
		m, n = n, m
	}

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		prev, curr = curr, prev
	}
	return prev[n]
}

func max[T ~int | ~float32](nums ...T) T {

	if len(nums) == 0 {
		var zero T
		return zero
	}

	maximum := nums[0]

	for i := range nums {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	return maximum
}

func main() {
	testCases := []struct {
		text1  string
		text2  string
		expect int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"abcdef", "acef", 4},
		{"", "", 0},         // Empty strings
		{"abc", "", 0},      // One empty string
		{"", "def", 0},      // One empty string
		{"fish", "fosh", 3}, //Similar strings
	}

	lcsFuncs := map[string]func(string, string) int{
		"Recursive":      longestCommonSubsequenceRecursive,
		"Recursive2":     longestCommonSubsequenceRecursive2,
		"Memoization":    longestCommonSubsequenceMemo,
		"Memoization2":   longestCommonSubsequenceMemo2,
		"Tabulation":     longestCommonSubsequenceTab,
		"SpaceOptimized": longestCommonSubsequenceSpace,
	}

	for name, lcsFunc := range lcsFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: text1=\"%s\", text2=\"%s\"\n", i+1, tc.text1, tc.text2)
			res := lcsFunc(tc.text1, tc.text2)
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
