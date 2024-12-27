package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []struct {
		input interface{}
	}{
		{[]interface{}{1}},
		{[]interface{}{[]interface{}{1}}},
		{[]interface{}{[]interface{}{1, 2}, 3}},
		{[]interface{}{2, 3, []interface{}{4, 1, 2}}},
		{[]interface{}{[]interface{}{[]interface{}{1}}}},
		{[]interface{}{1, []interface{}{2, []interface{}{3}}}},
		{[]interface{}{1, 2, []interface{}{7, []interface{}{3, 4}, 2}}},
	}

	for i, testCase := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		result := peculiarArraySum(testCase.input)
		fmt.Printf("Input: %v\n", testCase.input)
		fmt.Printf("Result: %d\n", result)
		fmt.Println()
	}
}


func peculiarArraySum(arr interface{}) int {
	return peculiarArraySumHelper(arr, 1)
}

func peculiarArraySumHelper(arr interface{}, level int) int {
	sum := 0
	switch v := arr.(type) {
	case []interface{}:
		for _, val := range v {
			sum += peculiarArraySumHelper(val, level+1)
		}
		return int(math.Pow(float64(sum), float64(level)))
	case int:
		return v
	}
	return 0 // Should not reach here given the constraints
}
