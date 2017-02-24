package main

import (
	"fmt"
	"math"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var (
		t            int
		a, b         int
		aRoot, bRoot int64
		result       int64
	)

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		aRoot = int64(math.Floor(math.Sqrt(float64(a))))
		bRoot = int64(math.Floor(math.Sqrt(float64(b))))
		result = bRoot - aRoot
		if math.Floor(math.Sqrt(float64(a))) == math.Sqrt(float64(a)) {
			result++
		}
		fmt.Println(result)
	}
}
