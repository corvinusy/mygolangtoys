// Paul Graham exercise:
// create function that take n and return function taking i returning n + i

package main

import "fmt"

type fnBar func(int) int

func main() {
	n := 5
	bar := foo(n)
	i := 3
	fmt.Println(bar(i))
}

func foo(n int) fnBar {
	return func(i int) int {
		return n + i
	}
}
