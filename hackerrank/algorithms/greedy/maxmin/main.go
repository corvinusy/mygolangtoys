package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(getMinimalUnfairness(a, k-1))
}

func getMinimalUnfairness(a []int, k int) int {
	minUnf := a[k] - a[0]
	for i := 1; i < len(a)-k; i++ {
		if a[k+i]-a[i] < minUnf {
			minUnf = a[k+i] - a[i]
		}
	}
	return minUnf
}
