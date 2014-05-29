package main

import (
	"fmt"
	"math"
)

const LIMIT = 1e8 // there is no diffdigital primes that are greater than 1e8

var ddPrimes []int

func main() {

	ddPrimes = newAtkinPrimes(LIMIT)

	fmt.Println(len(ddPrimes), ddPrimes[len(ddPrimes)-1])

	for i := 0; i < len(ddPrimes); i++ {

	}

}

/*----------------------------------------------------------------------------*/
func getSetsNumber(nums []int) int {

	result := 0

	for i := 0; i < len(ddPrimes); i++ {
		if isSatisfies(ddPrimes[i]) {
			// remove digits
			// if digitslen == 0 return set
			//result += setNumber(newDigits)
		}
	}

	return result

}

/*----------------------------------------------------------------------------*/
func newAtkinPrimes(limit int) []int {

	var i, x, y, n int

	sqr_lim := int(math.Sqrt(float64(limit))) + 1
	sieve_nums := make([]bool, limit+1)
	primes := make([]int, 0)

	sieve_nums[2] = true
	sieve_nums[3] = true

	for x = 1; x < sqr_lim; x++ {
		for y = 1; y < sqr_lim; y++ {

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

	for i = 2; i < sqr_lim; i++ {
		if sieve_nums[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				sieve_nums[j] = false
			}
		}
	}

	for i = 0; i <= limit; i++ {
		if sieve_nums[i] && isDiffDigital(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

/*----------------------------------------------------------------------------*/
func isDiffDigital(n int) bool {

	var ds [10]int

	for n > 0 {
		if n%10 == 0 {
			return false
		}

		ds[n%10] += 1

		if ds[n%10] == 2 {
			return false
		}

		n /= 10
	}

	return true
}
