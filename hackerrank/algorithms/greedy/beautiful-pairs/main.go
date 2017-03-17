package main

import (
	"fmt"
)

const limit = 1e3

func main() {
	var n, x int
	fmt.Scan(&n) //nolint
	a := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&x) //nolint
		a[x]++
	}
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x) //nolint
		if a[x] > 0 {
			a[x]--
			count++
		}
	}
	if count < n {
		count++
	} else {
		count--
	}
	fmt.Println(count)
}
