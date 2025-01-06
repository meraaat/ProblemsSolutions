package main

import (
	"fmt"
	"math"
)

// Recursive
func minDistanceRecursive(word1 string, word2 string) int {
	var minDistance func(i, j int) int
	minDistance = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return minDistance(i-1, j-1)
		}
		return 1 + int(math.Min(math.Min(float64(minDistance(i-1, j)), float64(minDistance(i, j-1))), float64(minDistance(i-1, j-1))))
	}
	return minDistance(len(word1)-1, len(word2)-1)
}

func minDistanceRecursive2(word1, word2 string) int {

	n := len(word1)
	m := len(word2)

	var numberOfOperations func(int, int) int
	numberOfOperations = func(index1, index2 int) int {
		// Base case

		if index1 > n-1 && index2 > m-1 {
			return 0
		}

		if index1 > n-1 {
			return m - index2
		}

		if index2 > m-1 {
			return n - index1
		}

		// Recursive case
		if word1[index1] == word2[index2] {
			return numberOfOperations(index1+1, index2+1)
		}

		insertOP := 1 + numberOfOperations(index1, index2+1)
		deleteOP := 1 + numberOfOperations(index1+1, index2)
		replaceOP := 1 + numberOfOperations(index1+1, index2+1)

		return min(insertOP, deleteOP, replaceOP)
	}

	return numberOfOperations(0, 0)
}

// Memoization
func minDistanceMemo(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := range memo {
		memo[i] = make([]int, len(word2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var minDistance func(i, j int) int
	minDistance = func(i, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if memo[i+1][j+1] != -1 {
			return memo[i+1][j+1]
		}
		if word1[i] == word2[j] {
			memo[i+1][j+1] = minDistance(i-1, j-1)
			return memo[i+1][j+1]
		}
		memo[i+1][j+1] = 1 + int(math.Min(math.Min(float64(minDistance(i-1, j)), float64(minDistance(i, j-1))), float64(minDistance(i-1, j-1))))
		return memo[i+1][j+1]
	}
	return minDistance(len(word1)-1, len(word2)-1)

}

func minDistanceMemo2(word1, word2 string) int {

	n := len(word1)
	m := len(word2)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var numberOfOperations func(int, int) int
	numberOfOperations = func(index1, index2 int) int {
		// Base case
		if index1 > n-1 && index2 > m-1 {
			return 0
		}

		if index1 > n-1 {
			return m - index2
		}

		if index2 > m-1 {
			return n - index1
		}

		// Recursive case
		if memo[index1][index2] != -1 {
			return memo[index1][index2]
		}

		if word1[index1] == word2[index2] {
			memo[index1][index2] = numberOfOperations(index1+1, index2+1)
		} else {
			insertOP := 1 + numberOfOperations(index1, index2+1)
			deleteOP := 1 + numberOfOperations(index1+1, index2)
			replaceOP := 1 + numberOfOperations(index1+1, index2+1)
			memo[index1][index2] = min(insertOP, deleteOP, replaceOP)
		}

		return memo[index1][index2]
	}

	return numberOfOperations(0, 0)
}

// Tabulation
func minDistanceTab(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1])))
			}
		}
	}
	return dp[m][n]
}

// Space Optimized
func minDistanceSpace(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	if m < n {
		word1, word2 = word2, word1
		m, n = n, m
	}

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for j := 0; j <= n; j++ {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = 1 + min(prev[j], curr[j-1], prev[j-1])
			}
		}
		prev, curr = curr, prev
	}
	return prev[n]
}

func min[T ~int | ~float64](nums ...T) T {

	if len(nums) == 0 {
		var zero T
		return zero
	}

	minimum := nums[0]
	for i := range nums {
		if nums[i] <= minimum {
			minimum = nums[i]
		}
	}

	return minimum
}

func main() {
	testCases := []struct {
		word1  string
		word2  string
		expect int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "", 0},
		{"a", "", 1},
		{"", "b", 1},
		{"abc", "ac", 1},
		{"kitten", "sitting", 3},
		{"apple", "aple", 1},
		{"intention", "intention", 0},
	}

	minDistanceFuncs := map[string]func(string, string) int{
		"Recursive":      minDistanceRecursive,
		"Recursive2":     minDistanceRecursive2,
		"Memoization":    minDistanceMemo,
		"Memoization2":   minDistanceMemo2,
		"Tabulation":     minDistanceTab,
		"SpaceOptimized": minDistanceSpace,
	}

	for name, minDistanceFunc := range minDistanceFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: word1=\"%s\", word2=\"%s\"\n", i+1, tc.word1, tc.word2)
			res := minDistanceFunc(tc.word1, tc.word2)
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
