package main

import (
	"fmt"
	"sort"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var (
		n, i, c int
	)
	fmt.Scan(&n)
	a := make([]int, n)
	for i = range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	predicate := func(i int) bool { return a[i] > 0 }

	for i = sort.Search(len(a), predicate); i < len(a); i = sort.Search(len(a), predicate) {
		c = a[i]
		fmt.Println(len(a) - i)
		for j := i; j < len(a); j++ {
			a[j] -= c
		}
	}
}
