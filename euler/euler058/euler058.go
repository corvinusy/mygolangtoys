package main

import (
	"fmt"
	"math"
)

func main() {

	miller_rabin_solution()
	//  prime_map_solution()
}

/*-----------------------------------------------------------------------------*/
func miller_rabin_solution() {

	const SIZE = 26250

	var (
		i, upright uint64
	)

	count_all := 1
	count_primes := 0

	for i = 3; i <= SIZE; i += 2 {
		upright = i * i
		count_all += 4
		if is_mr_prime(upright - i + 1) {
			count_primes++
		}
		if is_mr_prime(upright - i - i + 2) {
			count_primes++
		}
		if is_mr_prime(upright - i - i - i + 3) {
			count_primes++
		}
		if (float64(count_primes) / float64(count_all)) < 0.1 {
			fmt.Println(count_primes, count_all, float64(count_primes)/float64(count_all), i)
			return
		}
	}
}

/*-----------------------------------------------------------------------------*/
func is_mr_prime(n uint64) bool {
	var i, upper uint64

	upper = 2 * log2_n(n) * log2_n(n)
	for i = 3; (i < upper) && (i < n); i += 1 + upper/10 {
		if !is_witness(i, n) {
			return false
		}
	}

	return true
}

/*-----------------------------------------------------------------------------*/
func is_witness(a, n uint64) bool {
	u := n / 2
	t := 1
	for u%2 == 0 {
		u /= 2
		t++
	}

	prev := exp_a_n_mod(a, u, n)

	var curr uint64

	for i := 1; i <= t; i++ {
		curr = (prev * prev) % n
		if (curr == 1) && (prev != 1) && (prev != n-1) {
			return false
		}
		prev = curr
	}

	if curr != 1 {
		return false
	}
	return true
}

/*-----------------------------------------------------------------------------*/
func exp_a_n_mod(a, n, mod uint64) uint64 { // fast (a ** n ) % mod
	var result uint64 = 1

	power := a

	for n != 0 {
		if n%2 != 0 {
			result = (result * power) % mod
		}
		power = (power * power) % mod
		n /= 2
	}
	return result
}

/*-----------------------------------------------------------------------------*/
func log2_n(n uint64) uint64 { // fast log2(n)
	var result uint64 = 0

	if n >= 1<<32 {
		n >>= 32
		result += 32
	}
	if n >= 1<<16 {
		n >>= 16
		result += 16
	}
	if n >= 1<<8 {
		n >>= 8
		result += 8
	}
	if n >= 1<<4 {
		n >>= 4
		result += 4
	}
	if n >= 1<<2 {
		n >>= 2
		result += 2
	}
	if n >= 1<<1 {
		result += 1
	}

	return result
}

/*-----------------------------------------------------------------------------*/
func prime_map_solution() {

	const SIZE = 30001

	var (
		i, upright uint64
	)

	prmap := make(map[uint64]bool, 0)

	create_primes_atkin(SIZE*SIZE, prmap)

	fmt.Println("passed")

	count_all := 1
	count_primes := 0

	for i = 3; i <= SIZE; i += 2 {
		upright = i * i
		count_all += 4
		if prmap[upright-i+1] {
			count_primes++
		}
		if prmap[upright-i-i+2] {
			count_primes++
		}
		if prmap[upright-i-i-i+3] {
			count_primes++
		}
		if (float64(count_primes) / float64(count_all)) < 0.1 {
			fmt.Println(count_primes, count_all, float64(count_primes)/float64(count_all), i)
			return
		}
	}
}

/*-----------------------------------------------------------------------------*/
func create_primes_atkin(limit uint64, prmap map[uint64]bool) {

	var sqr_lim uint64 = uint64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n uint64

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
			prmap[i] = true
		}
	}

	return
}

/*-----------------------------------------------------------------------------*/
