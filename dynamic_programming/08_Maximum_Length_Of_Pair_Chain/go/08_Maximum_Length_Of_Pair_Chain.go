package main

import (
	"fmt"
	"math"
	"sort"
)

// Pair struct for easier sorting
type Pair struct {
	Left  int
	Right int
}

// Convert [][]int to []Pair
func toPairs(pairs [][]int) []Pair {
	ps := make([]Pair, len(pairs))
	for i, p := range pairs {
		ps[i] = Pair{Left: p[0], Right: p[1]}
	}
	return ps
}

// Convert []Pair back to [][]int for testing output
// func fromPairs(ps []Pair) [][]int {
// 	pairs := make([][]int, len(ps))
// 	for i, p := range ps {
// 		pairs[i] = []int{p.Left, p.Right}
// 	}
// 	return pairs
// }

// Recursive
func findLongestChainRecursive(pairs [][]int) int {
	ps := toPairs(pairs)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Left < ps[j].Left
	})
	var findChain func(int) int
	findChain = func(index int) int {
		if index == len(ps) {
			return 0
		}
		maxLength := 1
		for i := index + 1; i < len(ps); i++ {
			if ps[index].Right < ps[i].Left {
				maxLength = int(math.Max(float64(maxLength), float64(1+findChain(i))))
			}
		}
		return maxLength
	}

	maxLength := 0
	for i := 0; i < len(ps); i++ {
		maxLength = int(math.Max(float64(maxLength), float64(findChain(i))))
	}
	return maxLength
}

// Memoization
func findLongestChainMemo(pairs [][]int) int {
	ps := toPairs(pairs)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Left < ps[j].Left
	})
	memo := make([]int, len(ps))
	for i := range memo {
		memo[i] = -1
	}

	var findChain func(int) int
	findChain = func(index int) int {
		if index == len(ps) {
			return 0
		}
		if memo[index] != -1 {
			return memo[index]
		}
		maxLength := 1
		for i := index + 1; i < len(ps); i++ {
			if ps[index].Right < ps[i].Left {
				maxLength = int(math.Max(float64(maxLength), float64(1+findChain(i))))
			}
		}
		memo[index] = maxLength
		return maxLength
	}
	maxLength := 0
	for i := 0; i < len(ps); i++ {
		maxLength = int(math.Max(float64(maxLength), float64(findChain(i))))
	}
	return maxLength
}

// Tabulation
func findLongestChainTab(pairs [][]int) int {
	ps := toPairs(pairs)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Left < ps[j].Left
	})
	n := len(ps)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if ps[j].Right < ps[i].Left {
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

// Greedy
func findLongestChainGreedy(pairs [][]int) int {
	ps := toPairs(pairs)
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Right < ps[j].Right
	})

	currentEnd := math.MinInt
	chainLength := 0

	for _, p := range ps {
		if p.Left > currentEnd {
			currentEnd = p.Right
			chainLength++
		}
	}
	return chainLength
}

func main() {
	testCases := []struct {
		pairs  [][]int
		expect int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
		{[][]int{{1, 2}, {7, 8}, {4, 5}}, 3},
		{[][]int{{1, 2}, {2, 9}, {4, 5}}, 2},
		{[][]int{{1, 2}}, 1},
		{[][]int{}, 0},
		{[][]int{{1, 2}, {1, 3}}, 1}, // Test case where left values are the same
		{[][]int{{1, 2}, {3, 4}, {2, 8}}, 2},
		{[][]int{{1, 2}, {2, 3}, {4, 5}, {1, 9}}, 3},
		{[][]int{{-10, -8}, {8, 9}, {-5, 0}, {6, 10}, {-6, -4}, {1, 7}, {9, 10}, {-4, 7}}, 4},
	}

	findLongestChainFuncs := map[string]func([][]int) int{
		"Recursive":   findLongestChainRecursive,
		"Memoization": findLongestChainMemo,
		"Tabulation":  findLongestChainTab,
		"Greedy":      findLongestChainGreedy,
	}

	for name, findLongestChainFunc := range findLongestChainFuncs {
		fmt.Printf("Testing %s:\n", name)
		for i, tc := range testCases {
			fmt.Printf("Test Case %d: pairs=%v\n", i+1, tc.pairs)
			res := findLongestChainFunc(tc.pairs)
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
