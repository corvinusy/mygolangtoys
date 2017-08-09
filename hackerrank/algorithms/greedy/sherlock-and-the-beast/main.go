package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		repl()
	}
}

func repl() {
	var k, n int
	fmt.Scan(&n)
	switch n {
	case 1, 2, 4, 7:
		fmt.Println("-1")
		return
	case 5, 10:
		fmt.Println(strings.Repeat("3", n))
		return
	}
	for k = n; k > 2; k -= 5 {
		if k%3 == 0 {
			fmt.Println(strings.Repeat("5", k) + strings.Repeat("3", n-k))
			return
		}
	}

}
