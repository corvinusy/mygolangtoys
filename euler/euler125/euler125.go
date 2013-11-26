package main

import (
	"fmt"
	"math"
)

const LIMIT = 1e8

func main() {

	result := 0

	for i := 2; i < LIMIT ; i++ {
		if is_palindrome(i) && is_quadseq(i) {
			result += i
		}
	}
	fmt.Println(result)
}
/*-----------------------------------------------------------------------------*/
func is_palindrome(n int) bool {

	var ds [10]int

	nl := 0

	for n > 0 {
		ds[nl] = n % 10
		n /= 10 
		nl++
	}

	for i :=0 ; i < (nl >> 1) ; i++ {
		if ds[i] != ds[nl-i-1] {
			return false
		}
	}
	return true
}
/*-----------------------------------------------------------------------------*/
func is_quadseq (n int) bool {

	sqn := int(math.Sqrt(float64(n)))
	sum := 0

	for i := 1; i <= sqn; i++ {
		sum = 0
		for j := i; sum <= n; j++ {
			sum += j*j
			if sum == n && j-i > 0 {
				return true
			}
		}
	}
	return false
}
/*-----------------------------------------------------------------------------*/
