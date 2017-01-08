package main

import (
	"fmt"
)

func main() {
	var n, d, k int
	fmt.Scan(&n, &d)
	a := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		a[k] = true
	}
	var triplets int
	for k := range a {
		if a[k+d] && a[k+d+d] {
			triplets++
		}
	}
	fmt.Println(triplets)
}
