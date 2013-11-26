package main

import (
	"fmt"
	"math"
	"time"
)

const HIGHLIMIT = 600
const HIGHPRIME = 78498

func main() {
	
	var (
		sum int64 = 0
		res_prime int64 = 953
		res_adds int = 21
		i, j, n int
	)

	primes := make([]int64, 1e6)
	prime_set := make(map[int64]bool, 1e6)
	time1 := time.Now();

	create_primes(2e7, &primes, &prime_set)
	time3 := time.Since(time1)
	fmt.Println(time3)

	for i = 168; i < HIGHPRIME; i ++ {
		for j = 0; j < i; j++ {
			sum = 0
			for n = j; (n < j + HIGHLIMIT) && (n < i); n++ {
				sum += primes[n]
				if sum > primes[i] {
					break
				}
				if (sum == primes[i]) && (res_adds < n - j) {
					res_prime = sum
					res_adds = n - j
//					fmt.Println(i, primes[i], res_adds)
				}
			}

		}
	}

	time2 := time.Since(time1);
	fmt.Println(res_prime, " ", res_adds, "\n", time2)
}
/*-----------------------------------------------------------------------------*/
func create_primes (limit int64, primes *([]int64), prime_set*(map[int64]bool)) {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n int64;

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
			(*prime_set)[int64(i)] = true
		}
	}

	return
}
