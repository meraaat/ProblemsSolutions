package main

import "fmt"

func main() {
	n := 3
	totalMoves := toh(n, 1, 3, 2)
	fmt.Printf("Total moves for %d disks: %d\n", n, totalMoves)
	
	n = 2
	totalMoves = toh(n, 1, 3, 2)
	fmt.Printf("Total moves for %d disks: %d\n", n, totalMoves)
}

// Approach : T O(n^2) | S O(n)
func toh(n int, from int, to int, aux int) int {
	if n == 1 {
		fmt.Printf("Move disk 1 from rod %d to rod %d\n", from, to)
		return 1
	}
	moves := 0
	moves += toh(n-1, from, aux, to)
	fmt.Printf("Move disk %d from rod %d to rod %d\n", n, from, to)
	moves++
	moves += toh(n-1, aux, to, from)
	return moves
}
