package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	
	var (
		a, b, n, ares, bres, cres int64
	)
	
	const LIMIT = 1e7

	primes := make([]int64, 0, LIMIT)
	prmap := make(map[int64]bool, LIMIT)

	time1 := time.Now();

	create_primes_atkin(2*LIMIT, &primes, &prmap)

	time2 := time.Since(time1);
	fmt.Println("primes created for ", time2)

	cres = 0

	for a = -999; a < 1000; a++ {
		for b = -999; b < 1000; b += 2 {
			for n = 0; prmap[n * n + a * n + b]; n++ {
				;
			}
			if cres < n {
				ares = a
				bres = b
				cres = n
			}
		}
	}


	time3 := time.Since(time1)
	fmt.Println(ares * bres, "\n", time3)

}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit int64, primes *([]int64), prmap *(map[int64]bool)) int64 {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n, count int64;

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
	
	count = 0
	for i = 0; i <= limit; i++ {
		if sieve_nums[i] {
			*primes = append(*primes, i)
			(*prmap)[i] = true
			count++
		} 
	}

	return count
}
/*-----------------------------------------------------------------------------*/
