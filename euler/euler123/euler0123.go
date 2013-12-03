package main

import (
    "fmt"
	"math"
)

/*  (a % n + b % n ) % n = (a + b) % n
 *  (a % n) * (b % n) % n = (a * b) % n
 *  ((p-1) ** n + (p+1) ** n) % p*p = (((p-1)**n ) % p*p + ((p+1)**n ) % p*p) % p*p =
 *  // (p+1) % p = 1 , (p-1) % p = p-1 //
 *
 */

const LIMIT = 5e7

func main() {
	
	primes := create_primes_atkin(LIMIT)

	fmt.Println(remNPower(primes[2], 3))
	
	delta := uint64(10)

	for i := uint64(1); i < uint64(len(primes)); i++ {

		rem := remNPower(primes[i-1], i)
		if rem > delta {
			fmt.Println( "p [", i, "] =", primes[i-1], "delta = ", rem )
			delta *= 10
		}

	}

//	fmt.Println("result =", len(resmap))
	
}
/*-----------------------------------------------------------------------------*/
func remNPower(p, n uint64) uint64 {

	num1 := slowExpNMod(p-1, n, p*p)
	num2 := slowExpNMod(p+1, n, p*p)

	return (num1 + num2) % (p*p)
	
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit uint64) []uint64  {

    var i, x, y, n uint64

    sqr_lim := uint64(math.Sqrt(float64(limit)))
    sieve_nums := make([]bool, limit+1)
	primes := make([]uint64, 0)


    for i = 5; i <= limit ; i++ {
        sieve_nums[i] = false
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
            primes = append(primes, i)
        } 
    }

    return primes
}
/*-----------------------------------------------------------------------------*/
func expNMod (p, n, mod uint64) uint64 { // fast (p ** n ) % mod for p < 60000

    var result uint64 = 1

    for n != 0 {
        if n & 1 != 0 {
            result = (result * p) % mod
        }
        p = (p * p) % mod
        n /= 2
    }
    return result
}
/*-----------------------------------------------------------------------------*/
func slowExpNMod (p, n, mod uint64) uint64 { // fast (p ** n ) % mod for p < 60000

    var result uint64 = 1

    for n != 0 {
        result = (result * p) % mod
        n--
    }
    return result
}
/*-----------------------------------------------------------------------------*/
