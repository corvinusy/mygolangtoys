package main

import (
	"fmt"
	"math"
)


// 1_2_3_4_5_6_7_8_9_0
// 1 2 3 4 5 6 7688900
// 830
const LIMIT = 1e9

func main() {

	start := uint64(math.Sqrt(float64(10203040506070809)))
	end := uint64(math.Sqrt(float64(19293949596979899)))

	for i := start; i <= end; i++ {
		if is_satisfies(i*i) {
			fmt.Println(i, i*i)
			return
		}
	}

}
/*-----------------------------------------------------------------------------*/
func is_satisfies(n uint64) bool {

	t := uint64(9)
	
	for n > 0 && t > 0 {
		if n % 10 != t { return false }
		n /= 100
		t--
	}
	return true
}
