package main

import (
	"fmt"
	"sort"
)

var ms []int

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	ms = make([]int, m)
	for i := range ms {
		fmt.Scan(&ms[i])
	}
	sort.Ints(ms)
	var count int
	remains(&count, n, len(ms))
	fmt.Println(count)
}

func remains(c *int, n, l int) {
	for i := 0; i < l; i++ {
		if ms[i] < n {
			remains(c, n-ms[i], i+1)
		}
		if ms[i] == n {
			*c++
			return
		}
	}
}
