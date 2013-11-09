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

	time1 := time.Now();

	create_primes_atkin(LIMIT, &primes, &prmap)

	SIZE := uint64(len(primes)-1)

	for i = 644; i < primes[SIZE]; i++ {
		if divcount(i, primes) != 4 {
			continue
		}
		if divcount(i+1, primes) != 4 {
			continue
		}
		if divcount(i+2, primes) != 4 {
			continue
		}
		if divcount(i+3, primes) == 4 {
			fmt.Println(i)
			time3 := time.Since(time1)
			fmt.Println("\n", time3)
			return
		}
	}
}
/*-----------------------------------------------------------------------------*/
func divcount(num uint64, primes []uint64) uint64 {
	var n, count uint64
	
	count = 0
	for n = 0; primes[n] < num/6; n++ {
		if num % primes[n] == 0 {
			count++
		}
	}
	return count
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
