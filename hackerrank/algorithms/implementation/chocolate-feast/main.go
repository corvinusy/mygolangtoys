package main

import "fmt"

func main() {
	var t, n, c, m int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&n, &c, &m)
		fmt.Println(n/c + solve(m, n/c))
	}
}

func solve(m, w int) int {
	var total int
	for x := w / m; x > 0; x = w / m {
		total += x
		w = w%m + x
	}
	return total
}
