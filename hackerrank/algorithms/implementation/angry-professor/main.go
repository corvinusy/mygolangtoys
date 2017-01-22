package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)
		a := make([]int, n)
		var num int
		for i := range a {
			fmt.Scan(&a[i])
			if a[i] <= 0 {
				num++
			}
		}
		if num >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
