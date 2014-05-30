package main

import (
	"fmt"
	"github.com/cznic/mathutil"
)

const limit = 1e9

func main() {

	var dr uint64
	var n uint64

	for n = 2; n <= limit; n += 2  {

		dr = getDiophantReciprocals(n)

		if dr > 1e3 {
			fmt.Printf("n = %4d : a(n) = %d\n", n, dr)
			break
		}
	}
	
	return
}
/*----------------------------------------------------------------------------*/
func getNumberOfDivisors(n uint64) uint64 {
	
	factorPowers := factorPowersUint64(n)

	//calculate sigma function
	result := uint64(1)

	for i := range factorPowers {
		result *= factorPowers[i] + 1
	}

	return result
}
/*----------------------------------------------------------------------------*/
func getDiophantReciprocals(n uint64) uint64 {
	return (getNumberOfDivisors(n*n) + 1) / 2
}
/*----------------------------------------------------------------------------*/
func factorPowersUint64(n uint64) []uint64 {

	factorPowers := make([]uint64, 0, 20)
	prime32 := uint32(0)

	for {
		var ok bool
		if prime32, ok = mathutil.NextPrime(prime32); !ok {
			break
		}
		
		prime := uint64(prime32)
		if prime*prime > n {
			break
		}
		
		power := uint64(0)
		
		for n % prime == 0 {
			n /= prime
			power++
		}

		if power != 0 {
			factorPowers = append(factorPowers, power)
		}

		if n == 1 {
			break
		}
	}

	if n != 1 {
		factorPowers = append(factorPowers, 1)
	}

	return factorPowers
}
