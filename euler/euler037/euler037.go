package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {

	var sum int64 = 0
	primes := make([]int64, 0, 1e6)
	time1 := time.Now()

	prime_list(2e6, &primes)

	for i := 4; i < 100000; i++ {
		if is_round_prime(i, &primes) {
			sum = sum + primes[i]
			fmt.Println(primes[i])
		}
	}

	time2 := time.Since(time1)
	fmt.Println(sum, "\n", time2)
}

/*------------------------------------------------------*/
func prime_list(limit int64, primes *([]int64)) {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n int64

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
			*primes = append(*primes, i)
		}
	}

	return
}

/*------------------------------------------------------*/
func is_round_prime(id int, primes *[]int64) bool {
	var str string = strconv.FormatInt((*primes)[id], 10)
	var r, l int64
	var r_prime, l_prime bool

	for i := 1; i < len(str); i++ {
		r, l = split_num(str, i)

		r_prime = false
		for j := 0; j < id; j++ {
			if r == (*primes)[j] {
				r_prime = true
				break
			}
		}
		if !r_prime {
			return false
		}

		l_prime = false
		for j := 0; j < id; j++ {
			if l == (*primes)[j] {
				l_prime = true
				break
			}
		}
		if !l_prime {
			return false
		}
	}
	return true
}

/*------------------------------------------------------*/
func split_num(s string, c int) (int64, int64) {
	r, _ := strconv.ParseInt(s[:c], 10, 0)
	l, _ := strconv.ParseInt(s[c:], 10, 0)

	return r, l
}
