package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var q, t float64
	for i := 1; i < n-1; i++ {
		for j := i; j < n-i; j++ {
			t = math.Sin(float64(i)) + math.Sin(float64(j)) + math.Sin(float64(n-i-j))
			if t > q {
				q = t
			}
		}
	}
	fmt.Printf("%.9f\n", q)
}
