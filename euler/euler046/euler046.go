package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	
	var (
		i uint64
	)
	
	const LIMIT = 1e6

	primes := make([]uint64, 0, LIMIT)
	prmap := make(map[uint64]bool, LIMIT)

	// add 1 to primes
	prmap[1] = true
	primes = append(primes, 1)

	time1 := time.Now();

	create_primes_atkin(LIMIT, &primes, &prmap)

	for i = 0; i < 30; i++ {
		fmt.Printf("%d ", primes[i])
	}

	for i = 9; i < LIMIT; i +=2 {
		if prmap[i] {
			continue
		}
		if !is_satisfies(i, primes) {
			fmt.Println("\n",i)
			return
		}
	}
	
	time3 := time.Since(time1)
	fmt.Println("\n", time3)

}
/*-----------------------------------------------------------------------------*/
func is_satisfies(num uint64, primes []uint64) bool {
	var n, m uint64
	
	for n = 0; primes[n] < num; n++ {
		for m = 1; (primes[n] + m * m * 2) <= num ; m++ {

			if num == (primes[n] + m * m * 2) {
				return true
			}
		}
	}
	return false
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit uint64, primes *([]uint64), prmap *(map[uint64]bool))  {

	var sqr_lim uint64 = uint64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n uint64;

	for i = 5; i <= limit ; i++ {
		sieve_nums[i] = false;
	}

	sieve_nums[2] = true
	sieve_nums[3] = true
	
	for x = 1; x <= sqr_lim; x++ {
		for y = 1; y <= sqr_lim; y++ {

			n = 4 * x * x + y * y
			if (n <= limit) && ( (n % 12 == 1) || (n % 12 == 5) ) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - x * x
			if (n <= limit) && (n % 12 == 7) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - 2 * y * y
			if (x > y) && (n <= limit) && (n % 12 == 11) {
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
			*primes = append(*primes, i)
			(*prmap)[i] = true
		} 
	}

	return
}
/*-----------------------------------------------------------------------------*/
