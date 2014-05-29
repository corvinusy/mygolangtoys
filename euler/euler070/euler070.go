package main

import (
	"fmt"
	"time"
)

const LIMIT = 1e7 + 1

func main() {

	var (
		ratio, res_ratio float64 = 1, LIMIT
		res_n            int     = 1
	)

	var tots [LIMIT + 1]int

	t1 := time.Now()

	// prepare totient_sieve
	for i := 1; i <= LIMIT; i++ {
		tots[i] = i
	}

	for i := 1; i <= LIMIT; i++ {
		for j := 2 * i; j < LIMIT; j += i {
			tots[j] -= tots[i]
		}
	}

	// find min ratio for permutations
	for i := 2; i < len(tots)-1; i++ {
		if is_permutation(i, tots[i]) {
			ratio = float64(i) / float64(tots[i])
			if ratio < res_ratio {
				res_ratio = ratio
				res_n = i
			}
		}
	}

	t2 := time.Since(t1)

	fmt.Println(res_n, tots[res_n], res_ratio, t2)

	return
}

/*-----------------------------------------------------------------------------*/
func is_permutation(n1, n2 int) bool {
	d := [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for n1 > 0 {
		d[n1%10]++
		n1 /= 10
	}

	for n2 > 0 {
		d[n2%10]--
		n2 /= 10
	}

	for _, b := range d {
		if b != 0 {
			return false
		}
	}

	return true
}
