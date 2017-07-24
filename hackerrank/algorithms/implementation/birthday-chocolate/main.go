package main

import (
	"fmt"
)

func main() {
	var n, m, d int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := range s {
		fmt.Scan(&s[i])
	}
	fmt.Scan(&d, &m)

	fmt.Println(totalPieces(s, d, m))
}

func totalPieces(s []int, d, m int) int {
	total := 0

	for i := 0; i < len(s); i++ {
		if sumOf(s, i, m) == d {
			total++
		}
	}
	return total
}

func sumOf(s []int, start, dist int) int {
	if start+dist > len(s) {
		return -1
	}
	sum := 0
	for i := start; i < start+dist; i++ {
		sum += s[i]
	}
	return sum
}
