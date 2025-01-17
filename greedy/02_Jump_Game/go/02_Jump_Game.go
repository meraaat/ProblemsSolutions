package main

import "fmt"

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+nums[i])
	}
	return reachable >= len(nums)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		nums   []int
		expect bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{1, 3, 4, 1, 1, 2, 1}, true},
		{[]int{1, 3, 4, 2, 1, 1, 0, 1, 1}, false},
		{[]int{0}, true},
		{[]int{2, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false}, // New test case
		{[]int{1, 1, 1, 1}, true},
		{[]int{1, 1, 0, 1}, false},
		{[]int{5, 0, 0, 0, 0, 0, 1}, false},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d: nums=%v\n", i+1, tc.nums)
		res := canJump(tc.nums)
		fmt.Printf("Got: %t\n", res)
		fmt.Printf("Expected: %t\n", tc.expect)
		if res != tc.expect {
			fmt.Println("FAIL")
		} else {
			fmt.Println("PASS")
		}
		fmt.Println()
	}
}
