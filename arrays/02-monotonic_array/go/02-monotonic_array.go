package main

import "fmt"

func main() {
	testCases := [][]int{
		{},
		{1},
		{3, 2, 1},
		{1, 2, 1},
		{1, 3, 2},
		{1, 2, 4, 2},
		{1, 2, 3, 5, 13},
		{1, 1, 2, 3, 3, 4, 4},
		{6, 5, 4, 4, 3, 2, 1},
	}

	for i, testcase := range testCases {
		fmt.Printf("TestCase %v: %v\n", i, isMonotonic(testcase))
	}
}

func UnOptimizedIsMonotonic(array []int) bool {
	if len(array) <= 1 {
		return true
	}

	first, last := array[0], array[len(array)-1]
	if first > last {
		for i := 1; i < len(array)-1; i++ {
			if array[i] < array[i+1] {
				return false
			}
		}

	} else if first < last {
		for i := 1; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				return false
			}
		}

	} else {
		// first == last
		for i := 1; i < len(array)-1; i++ {
			if array[i+1] != array[i] {
				return false
			}
		}
	}

	return true

}

// isMonotonic - T: O(n) | S: O(1)
func isMonotonic(array []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isIncreasing = false
		}
		if array[i] > array[i-1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
