package main

import (
	"fmt"
)

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}

		var total, lo, hi int
		for lo < len(a) {
			hi = lo + findMaxIdx(a[lo:])
			if hi == lo {
				lo++
				continue
			}
			total += findMaxProfitToLast(a[lo : hi+1])
			lo = hi + 1
		}
		fmt.Println(total)
	}
}

func findMaxIdx(a []int) int {
	var m, k int
	for i := range a {
		if m < a[i] {
			m = a[i]
			k = i
		}
	}
	return k
}

func findMaxProfitToLast(a []int) int {
	var profit int
	for i := range a {
		profit += a[len(a)-1] - a[i]
	}
	return profit
}
