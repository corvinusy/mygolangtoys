package main

import "fmt"

func main() {
	var q, n int
	for fmt.Scan(&q); q > 0; q-- {
		fmt.Scan(&n)
		if n%7 < 2 {
			fmt.Println("Second")
		} else {
			fmt.Println("First")
		}
	}
}
