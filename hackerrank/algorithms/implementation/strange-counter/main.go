package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	fmt.Println(defineCounterState(t))
}

func defineCounterState(t int) int {
	c := 3
	x := 1
	for x+c <= t {
		x += c
		c *= 2
		fmt.Println(x, c)
	}
	return c - t + x
}
