package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	c := make([]int, n)
	for i := range c {
		fmt.Scan(&c[i])
	}
	fmt.Println(getMinimalCost(c, k))
}

func getMinimalCost(c []int, k int) int {
	var cost, tour, i, j int
	sort.Ints(c)

	for i = len(c) - 1; i >= 0; i -= k {
		for j = 0; j < k; j++ {
			if i-j >= 0 {
				cost += (tour + 1) * c[i-j]
			}
		}
		tour++
	}
	return cost
}
