package main

import (
	"fmt"
)

// Partition Fuction P
// P(1, 1) = 1
// P(n, k) = P(n, k-1) + P(n - k, k)
// P(n, n) = 1
// P(n, 0) = 0
// P(n, 1) = n
// P(n, k > n) = 0
// P(n) = Sum(P(n, k) ) for k = 1..n

const LIMIT = 100

func main() {

	// create dynamic array for parts
	parts := make([][]int, LIMIT+1)

	result := 0

	for i := 0; i <= LIMIT; i++ {
		parts[i] = make([]int, i+1)
	}
	// init 1st member
	parts[1][1] = 1

	// get parts and make sum of them
	for i := 2; i <= LIMIT; i++ {
		result += get_parts(LIMIT, i, parts)
	}

	fmt.Println(result)

}

/*-----------------------------------------------------------------------------*/
func get_parts(i, j int, parts [][]int) int {

	switch {
	case i < j, j == 0:
		return 0

	case parts[i][j] != 0:
		return parts[i][j]

	case j == 1:
		parts[i][j] = 1
		return 1

	case i == j:
		parts[i][j] = 1
		return 1

	default:
		parts[i][j] = get_parts(i-1, j-1, parts) + get_parts(i-j, j, parts)
		return parts[i][j]
	}
}
