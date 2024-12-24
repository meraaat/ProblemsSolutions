package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := [][]int{
		{},
		{0, 5},
		{1, 3, 5},
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
	}

	for i, e := range testCases {
		fmt.Printf("TestCase %v:\n", i)
		fmt.Printf("Sorted Squared Array: %v\n", SortedSquaredArray(e))
		fmt.Printf("Tow Pointer: %v\n", SquareAndSortOptimized(e))
	}

}

// Approach 1:
// Sort and Square - T: O(n log n) | S: O(n)
func SortedSquaredArray(array []int) []int {

	result := make([]int, len(array))

	for i, num := range array {
		result[i] = num * num
	}

	sort.Ints(result)
	return result
}

// Using multiplication instead of math.Pow
/*
func square(input int) int {
	return int(math.Pow(float64(input), 2))
}
*/

// Approach 2:
// Two Pointers - T: O(n) | S: O(n)
func SquareAndSortOptimized(array []int) []int {

	result := make([]int, len(array))
	left, right := 0, len(array)-1

	for i := len(array) - 1; i >= 0; i-- {
		leftSquare := array[left] * array[left]
		rightSquare := array[right] * array[right]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}

	return result
}
