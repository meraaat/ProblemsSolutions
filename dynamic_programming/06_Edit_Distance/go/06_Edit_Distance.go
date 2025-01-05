package main

import (
	"fmt"
	"math"
)

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
				curr[j] = 1 + int(math.Min(math.Min(float64(prev[j]), float64(curr[j-1])), float64(prev[j-1])))
			}
		}
		prev, curr = curr, prev
	}
	return prev[n]
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
		{"kitten", "sitting", 3},
		{"apple", "aple", 1},
		{"intention", "intention", 0},
	}

	minDistanceFuncs := map[string]func(string, string) int{
		"Recursive":      minDistanceRecursive,
		"Memoization":    minDistanceMemo,
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
