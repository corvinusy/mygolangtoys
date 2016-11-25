package main

import (
	"fmt"
	"math"
)

var (
	sieve    map[uint]bool
	sieveMax uint
)

func main() {
	var n uint
	fmt.Scan(&n)
	data := make([]uint, n)
	for i := range data {
		fmt.Scan(&data[i])
	}

	sieve = make(map[uint]bool)
	sieve[2] = true
	sieve[3] = true
	sieve[5] = true
	sieve[7] = true

	sieveMax = 7

	for i := range data {
		if isPrime(data[i]) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(x uint) bool {
	switch {
	case x == 1:
		return false
	case x == 2:
		return true
	case x%2 == 0:
		return false
	default:
		sqr := uint(math.Sqrt(float64(x))) + 1
		extendSieve(sqr)
		return !isDivisable(x)
	}
}

func extendSieve(cap uint) {
	var ok bool
	for i := sieveMax + 2; sieveMax <= cap; i += 2 {
		ok = true
		for k := range sieve {
			if i%k == 0 {
				ok = false
				break
			}
		}
		if ok {
			sieve[i] = true
			sieveMax = i
		}

	}
}

func isDivisable(x uint) bool {
	if sieve[x] {
		return false
	}
	for k := range sieve {
		if x%k == 0 {
			return true
		}
	}
	return false
}
