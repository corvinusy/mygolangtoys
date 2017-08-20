package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	makeSwaps(a, k)
	// print
	for i := 1; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
}

func makeSwaps(a []int, k int) {
	l := len(a)
	// make derive array
	b := make([]int, l)
	for i := 1; i < l; i++ {
		b[a[i]] = i
	}
	// do swaps
	var idx int
	for i := 1; i < l; i++ {
		if a[i] != l-i {
			idx = b[l-i]
			a[i], a[idx] = a[idx], a[i]
			b[a[idx]] = idx
			b[a[i]] = i
			k--
			if k == 0 {
				break
			}
		}
	}
}
