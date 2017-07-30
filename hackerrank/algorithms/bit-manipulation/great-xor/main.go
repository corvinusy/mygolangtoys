package main

import "fmt"

func main() {
	var q, x, n, v int64
	for fmt.Scan(&q); q > 0; q-- {
		fmt.Scan(&x)
		n = 0
		for v = x; v > 0; v >>= 1 {
			n++
		}
		fmt.Println(1<<uint(n) - 1 - x)
	}
}
