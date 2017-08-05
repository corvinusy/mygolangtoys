package main

import "fmt"

func main() {
	var q, n, m int
	for fmt.Scan(&q); q > 0; q-- {
		fmt.Scan(&n, &m)
		if m == 1 || n%2 == 0 {
			fmt.Println("2")
		} else {
			fmt.Println("1")
		}
	}
}
