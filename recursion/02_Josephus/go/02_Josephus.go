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
		fmt.Printf("TestCase %v: \n", i)
		fmt.Printf("findTheWinner1: %v\n", findTheWinner0(testcase.n, testcase.k))
		fmt.Printf("findTheWinner2: %v\n", findTheWinner1(testcase.n, testcase.k))
		fmt.Printf("findTheWinner3: %v\n", findTheWinnerRecursive(testcase.n, testcase.k))
		fmt.Printf("findTheWinner4: %v\n", findTheWinner2(testcase.n, testcase.k))
	}
}

// Approach 0: T O(n^2) | S O(n)
func findTheWinner0(n int, k int) int {

	// Creating array from 1 to n
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	// Define the recursion helper function
	var helper func([]int, int) int
	helper = func(arr []int, startIndex int) int {
		// Base case
		if len(arr) == 1 {
			return arr[0]
		}

		// Recursion case
		indexToRemove := (startIndex + k - 1) % len(arr)
		arr = append(arr[:indexToRemove], arr[indexToRemove+1:]...)
		return helper(arr, indexToRemove)
	}

	return helper(arr, 0)
}

// Approach 1: T O(n) | S O(n)
func findTheWinner1(n int, k int) int {

	var josephus func(int) int
	josephus = func(n int) int {
		//Base case
		if n == 1 {
			return 0
		}
		// Recursive case
		return (josephus(n-1) + k) % n

	}

	return josephus(n) + 1
}

// Approach 1: T O(n) | S O(n) / cleaner
func findTheWinnerRecursive(n int, k int) int {
	if n == 1 {
		return 1
	}
	return (findTheWinnerRecursive(n-1, k)+k-1)%n + 1
}

// Approach 2: T O(n) | S O(1) / iterative
func findTheWinner2(n int, k int) int {

	var survivor int

	for i := 2; i <= n; i++ {
		survivor = (survivor + k) % i
	}

	return survivor + 1
}
