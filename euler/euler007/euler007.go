package main

import (
	"fmt"
	"math"
)

func main() {
	
	var count int64 = 2
	var i int64

	for i = 5; ; i += 2 {
		if is_prime(i) {
			count++
		}

		if count == 10001 {
			fmt.Println(i)
			break
		}
	}
}

func is_prime (n int64) bool {
	var i int64;
	for i = 3; i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n % i == 0 {
			return false;
			break;
		}
	}
	return true;
}
