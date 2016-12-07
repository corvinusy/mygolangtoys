package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	// sort it
	sort.Ints(a)
	// count
	var count int
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a) && a[j]-a[i] <= k; j++ {
			if a[j]-a[i] == k {
				count++
			}
		}
	}
	fmt.Println(count)
}
