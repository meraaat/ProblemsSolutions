package main

import "fmt"

func main() {
	testCases := []struct {
		n int
		k int
	}{
		{1, 1},
		{2, 1},
		{2, 2},
		{3, 1},
		{3, 2},
		{3, 3},
		{3, 4},
		{4, 5},
		{4, 8},
	}

	for i, testcase := range testCases {
		fmt.Printf("TestCase %v: %v\n", i, KthGrammar(testcase.n, testcase.k))
	}
}

// Recursion Approach : T O(n) | S O(n)
func KthGrammar(n int, k int) int {

	// Base case
	if n == 1 {
		return 0
	}

	// Recursive case
	length := 1 << (n - 2) // This is the length of the previous row.
	if k <= length {
		return KthGrammar(n-1, k)
	} else {
		return 1 - KthGrammar(n-1, k-length) // Inverse if in second half
	}
}

// Iterative (Bit Manipulation) Approach (Optimized)
// Iterative Approach : T O(n) | S O(1)
func KthGrammarIterative(n int, k int) int {
	k-- // 0-based indexing for bit manipulation
	count := 0
	for k > 0 {
		if k&1 == 1 { // Check if the last bit is set (1)
			count++
		}
		k >>= 1 // Right shift to check the next bit
	}

	if count&1 == 0 { // count%2 == 0
		return 0
	}
	return 1
}
