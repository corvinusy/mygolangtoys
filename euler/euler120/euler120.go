package main

import (
	"fmt"
)

/*
((a-1)**n + (a+1)**n ) % a*a
*/

const LIMIT = 1e5

func main() {

	var result, r, rmax, a, n uint64

	result = 0

	for a = 3; a <= 1e3; a++ {
		rmax = 0
		for n = 1; n < LIMIT; n++ {
			r = (exp_a_n_mod(a-1, n, a*a) + exp_a_n_mod(a+1, n, a*a)) % (a * a)
			if r > rmax {
				rmax = r
			}
		}
		fmt.Println(a, rmax)
		result += rmax
	}

	fmt.Println(result)

}

/*-----------------------------------------------------------------------------*/
func exp_a_n_mod(a, n, mod uint64) uint64 { // fast (a ** n ) % mod

	var result uint64 = 1

	power := a

	for n != 0 {
		if n&1 != 0 {
			result = (result * power) % mod
		}
		power = (power * power) % mod
		n /= 2
	}
	return result
}

/*-----------------------------------------------------------------------------*/
