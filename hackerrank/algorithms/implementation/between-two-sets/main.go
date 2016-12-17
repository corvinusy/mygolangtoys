package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}
	fmt.Println(countBetweens(a, b))
}

func countBetweens(a, b []int) int {
	var (
		count int
		ok    bool
	)

	sort.Ints(a)
	sort.Ints(b)
	lim := b[len(b)-1]
	for x := a[len(a)-1]; x <= lim; x++ {
		ok = true
		for i := range a {
			if x%a[i] != 0 {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		for i := range b {
			if b[i]%x != 0 {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}
	return count
}
