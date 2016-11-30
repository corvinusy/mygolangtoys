package main

import (
	"fmt"
	"math"
)

func main() {
	var upper, lower int
	fmt.Scan(&lower, &upper)
	sqrUp := int(math.Sqrt(float64(upper))) + 1
	preprimes := newPrePrimesErathosphen(sqrUp)
	primes := newPrimesErathosphen(lower, upper, preprimes)
	twins := countTwins(primes)
	fmt.Println(twins)
	// test
	//primesAtkin := newPrimesAtkin(lower, upper)
}

func newPrimesErathosphen(lower, upper int, preprimes []int) []int {
	sieve := make([]bool, upper-lower+1)
	var h, k int
	// prepare
	for i := range sieve {
		sieve[i] = true
	}
	// runner
	for i := range preprimes {
		h = lower % preprimes[i]
		if h == 0 {
			k = 0
		} else {
			k = preprimes[i] - h
		}
		for ; k <= upper-lower; k += preprimes[i] {
			sieve[k] = false
		}
	}
	// add preprimes to primes if ranges crossed
	var primes []int
	for i := range preprimes {
		if preprimes[i] >= lower {
			primes = append(primes, preprimes[i])
		}
	}
	// resolve sieve to primes
	for i := range sieve {
		if lower+i <= preprimes[len(preprimes)-1] {
			continue
		}
		if sieve[i] {
			primes = append(primes, lower+i)
		}
	}
	return primes
}

func newPrePrimesErathosphen(upper int) []int {
	sieve := make([]bool, upper+1)
	// prepare
	for i := range sieve {
		sieve[i] = true
	}
	sieve[1] = false
	// runner
	for i := 2; i*i <= upper; i++ {
		if sieve[i] {
			for j := i * i; j <= upper; j += i {
				sieve[j] = false
			}
		}
	}
	// convert sieve to slice
	var primes []int
	for i := 2; i <= upper; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func countTwins(a []int) int {
	var count int
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 2 {
			count++
		}
	}
	return count
}

func newPrimesAtkin(lower, upper int) []int {
	var sieve = make([]bool, upper+1)
	var n int
	sqrUp := int(math.Sqrt(float64(upper)))
	sieve[2] = true
	sieve[3] = true
	// main runner
	for x := 1; x <= sqrUp; x++ {
		for y := 1; y <= sqrUp; y++ {
			n = 4*x*x + y*y // 4*xx + yy
			if (n <= upper) && ((n%12 == 1) || (n%12 == 5)) {
				sieve[n] = !sieve[n]
			}
			n = n - x*x // 3*xx+yy ( 6*k+1 case )
			if (n <= upper) && (n%12 == 7) {
				sieve[n] = !sieve[n]
			}
			n = n - 2*y*y // 3*xx - yy (12k-1)
			if (x > y) && (n <= upper) && (n%12 == 11) {
				sieve[n] = !sieve[n]
			}
		}
	}
	// cleaning from 5*n
	for i := 5; i <= sqrUp; i++ {
		if sieve[i] {
			n = i * i
			for j := n; j <= upper; j += n {
				sieve[j] = false
			}
		}
	}
	// convert sieve to slice
	var primes []int
	for i := lower; i <= upper; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
