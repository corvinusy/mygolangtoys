package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		var n, m, s int
		fmt.Scanf("%d %d %d", &n, &m, &s)
		s--
		m %= n
		r := (s + m) % n
		if r == 0 {
			r = n
		}
		fmt.Println(r)
	}
}
