package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	var n int
	for i := 0; i < t; i++ {
		h := 1
		fmt.Scan(&n)
		for y := 0; y < n; y++ {
			if y%2 == 1 {
				h++
			} else {
				h *= 2
			}
		}
		fmt.Println(h)
	}
}
