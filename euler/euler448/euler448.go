package main

import (
	"fmt"
)

//99999999019

// for prime avelcm(n) =  n * n * (n/2) + n

// A(n) = sumLcm(n)

// a(n) = Sum_{k=1..n} n/GCD(n,k)
// A(n) = (1+a(n))/2, where a(n) = Sum_{k=1..n} n/GCD(n,k).
// a(x*y) = a(x) * a(y)
// a(prime) =  prime * (prime-1) + 1

const LIMIT = 1e5

const MOD = 999999017

//const MOD = 1e9

const SEEK = 99999999019

func main() {

	for n := int64(2); n <= 100; n++ {
		slc := aveLcmMod(n, MOD)
		slp := n*n*(n-1)/2 + n
		// slp := (n*(n-1)/2 + 1) * n = n*n - n*1/2 + 1
		k := int64(n)
		Ckn := k*n*(n-1)/2 + 1
		fmt.Println("sumLcm(", n, ") =", slc, "\tsumLcmPrime = ", slp, "\tdelta = ", (slp-slc*n)/n, "\tCkn", Ckn)
	}

	return

	fmt.Println("sf(10) =", sf(10))

	fmt.Println("sf(100) =", sf(100))

	fmt.Println(sf(SEEK))

}

/*-----------------------------------------------------------------------------*/

/*-----------------------------------------------------------------------------*/
func lcmMod(a, b, mod int64) int64 {

	return ((((a % mod) / gcdMod(a, b, MOD)) % mod) * (b % mod)) % mod

}

/*-----------------------------------------------------------------------------*/
func gcdMod(a, b, mod int64) int64 {

	for b != 0 {
		a, b = b, a%b
	}

	return a % mod
}

/*-----------------------------------------------------------------------------*/
func aveLcmMod(n, mod int64) int64 {

	var (
		sum int64 = 0
	)

	for p := int64(1); p <= n; p++ {

		sum = (sum + lcmMod(n, p, MOD)) % mod

	}

	return (sum / (n % mod)) % mod
}

/*-----------------------------------------------------------------------------*/
func sf(n int64) int64 {

	var (
		sum int64 = 0
		res int64
	)

	for i := int64(1); i <= n; i++ {
		res = aveLcmMod(i, MOD)
		sum = (sum%MOD + res%MOD) % MOD

		if i%1e4 == 0 {
			fmt.Println("sf(", i, ") =", sum)
		}

	}

	return sum
}
