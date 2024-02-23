package main

import (
	"fmt"
)

// Recursive
// Time complexity: O(n)
// Space complexity: O(n)
func Sum_to_n_a(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n + Sum_to_n_a(n-1)
}

// Simple for loop
// Time complexity: O(n)
// Space complexity: O(1)
func Sum_to_n_b(n int) int {
	sum := 0
	for curr := 1; curr <= n; curr++ {
		sum += curr
	}
	return sum
}

// Arithmetic series formula (best)
// Time complexity: O(1)
// Space complexity: O(1)
func Sum_to_n_c(n int) int {
	return int(float64(n) * ((float64(n) + 1) / 2))
}

func main() {
	fmt.Println(Sum_to_n_a(5))
	fmt.Println(Sum_to_n_b(5))
	fmt.Println(Sum_to_n_c(5))
}
