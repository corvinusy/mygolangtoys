package main

import (
	"fmt"
	"math"
)

//const LIMIT = 5e7

const LIMIT = 5e7

func main() {

	primes := create_primes_atkin(LIMIT)

	fmt.Println(primes[0], len(primes))

	resmap := make(map[int]bool, 0)

	UPPER1 := int(math.Sqrt(float64(LIMIT))) + 1
	for i1 := 0; primes[i1] <= UPPER1; i1++ {

		UPPER2 := int(math.Cbrt(float64(LIMIT-primes[i1]*primes[i1]))) + 1
		for i2 := 0; primes[i2] <= UPPER2; i2++ {

			UPPER3 := int(math.Sqrt(math.Sqrt(float64(LIMIT-primes[i1]*primes[i1]-primes[i2]*primes[i2]*primes[i2])))) + 1
			for i3 := 0; primes[i3] <= UPPER3; i3++ {
				sum := primes[i1]*primes[i1] + primes[i2]*primes[i2]*primes[i2] +
					primes[i3]*primes[i3]*primes[i3]*primes[i3]
				if sum < LIMIT {
					resmap[sum] = true
				} else {
					break
				}
			}
		}
	}
	fmt.Println("result =", len(resmap))

}

/*-----------------------------------------------------------------------------*/
func create_primes_atkin(limit int) []int {

	var i, x, y, n int

	sqr_lim := int(math.Sqrt(float64(limit)))
	sieve_nums := make([]bool, limit+1)
	primes := make([]int, 0)

	for i = 5; i <= limit; i++ {
		sieve_nums[i] = false
	}

	sieve_nums[2] = true
	sieve_nums[3] = true

	for x = 1; x <= sqr_lim; x++ {
		for y = 1; y <= sqr_lim; y++ {

			n = 4*x*x + y*y
			if (n <= limit) && ((n%12 == 1) || (n%12 == 5)) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - x*x
			if (n <= limit) && (n%12 == 7) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - 2*y*y
			if (x > y) && (n <= limit) && (n%12 == 11) {
				sieve_nums[n] = !sieve_nums[n]
			}
		}
	}

	for i = 5; i <= sqr_lim; i++ {
		if sieve_nums[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				sieve_nums[j] = false
			}
		}
	}

	for i = 0; i <= limit; i++ {
		if sieve_nums[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
