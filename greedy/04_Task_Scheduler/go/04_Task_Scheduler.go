package main

import (
	"fmt"
	"math"
)

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	freq := make([]int, 26)
	for _, task := range tasks {
		freq[task-'A']++
	}

	maxFreq := 0
	maxFreqCount := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
			maxFreqCount = 1
		} else if f == maxFreq {
			maxFreqCount++
		}
	}

	return int(math.Max(float64(len(tasks)), float64((maxFreq-1)*(n+1)+maxFreqCount)))
}

func main() {
	testCases := []struct {
		tasks  []byte
		n      int
		expect int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C'}, 2, 8},
		{[]byte{'A', 'A', 'B', 'B'}, 3, 6},
		{[]byte{'C', 'E', 'C', 'E', 'E', 'C'}, 2, 8},
		{[]byte{'C', 'E', 'C', 'E', 'E', 'C'}, 3, 10}, // Example from the prompt
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D', 'E'}, 2, 9},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'D', 'E'}, 3, 11},
		{[]byte{'A', 'A', 'A', 'A', 'B', 'B', 'C', 'C', 'D'}, 2, 11},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d: tasks=%s, n=%d\n", i+1, string(tc.tasks), tc.n)
		res := leastInterval(tc.tasks, tc.n)
		fmt.Printf("Got: %d\n", res)
		fmt.Printf("Expected: %d\n", tc.expect)
		if res != tc.expect {
			fmt.Println("FAIL")
		} else {
			fmt.Println("PASS")
		}
		fmt.Println()
	}
}
