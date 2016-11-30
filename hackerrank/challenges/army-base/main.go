package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(math.Ceil(float64(n)/2) * math.Ceil(float64(m)/2))
}
