package main

import (
	"fmt"
	"math"
)

func main() {
	var x, n int
	fmt.Scan(&x, &n)
	fmt.Println(countSumOfPowers(x, n, 1))
}

func countSumOfPowers(x, n, accu int) int {
	y := x - int(math.Pow(float64(accu), float64(n)))
	switch {
	case y == 0:
		return 1
	case y < 0:
		return 0
	default:
		return countSumOfPowers(y, n, accu+1) + countSumOfPowers(x, n, accu+1)
	}
}
