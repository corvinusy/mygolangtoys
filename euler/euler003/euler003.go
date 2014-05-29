package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	time1 := time.Now()
	const trg = 600851475143
	var subtrg int64 = int64(math.Sqrt(trg))
	var i int64 = 1
	var res int64

	for i = 1; i <= subtrg; i += 2 {
		if is_prime(i) && (trg%i == 0) {
			res = i
		}
	}
	time2 := time.Since(time1)
	fmt.Println(res, " ", time2)
}

func is_prime(n int64) bool {
	var i int64
	var sqn = int64(math.Sqrt(float64(n)))

	for i = 3; i <= sqn; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	//	fmt.Printf("prime n = %v\n", n)
	return true
}
