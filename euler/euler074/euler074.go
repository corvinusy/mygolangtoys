package main

import (
	"fmt"
)

// dirty bruteforce
const LIMIT = 1e6

func main() {

	var res, n, i uint64

	res = 0

	for i = 1; i <= LIMIT; i++ {

		chain := make(map[uint64]bool)

		for n = i; !chain[n]; n = next(n) {
			chain[n] = true
		}

		//		fmt.Println(i, len(chain))
		if len(chain) >= 60 {
			res++
		}
	}
	fmt.Println(res)
}

/*-----------------------------------------------------------------------------*/
func next(n uint64) uint64 {
	var res uint64 = 0
	facts := [...]uint64{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

	for n > 0 {
		res += facts[n%10]
		n /= 10
	}
	return res
}
