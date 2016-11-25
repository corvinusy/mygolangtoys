package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var n, k, q, tmp int
	fmt.Scanln(&n, &k, &q)
	r := ring.New(n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		r.Value = tmp
		r = r.Next()
	}

	//make rotation
	for i := 0; i < k; i++ {
		r = r.Prev()
	}

	// make array from ring
	a := make([]int, n)
	for i := range a {
		a[i] = r.Value.(int)
		r = r.Next()
	}

	// get queries
	qs := make([]int, q)
	for i := range qs {
		fmt.Scan(&qs[i])
	}

	for _, v := range qs {
		fmt.Println(a[v])
	}
}
