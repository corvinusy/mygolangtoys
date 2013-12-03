package main

import (
	"fmt"
	"math"
)


/* If n = (p1^a1)(p2^a2)...(pt^at), 
 * 2-partitions = ((2 a1 + 1)(2 a2 + 1) ... (2 at + 1) + 1)/2. 
 * We want ((2 a1 + 1)(2 a2 + 1) ... (2 at + 1) + 1)/2 > 1000 
 * (2 a1 + 1)(2 a2 + 1) ... (2 at + 1) >= 2000
 *
 */

const LIMIT = 1e7

var primes []int

func main() {

	primes = create_primes_atkin(1e6)


	for n := int(1); n < LIMIT; n++ {

		primeVec := getFactors(n)

		units := 1

		for i, _ := range primeVec {
			units *= 2*primeVec[i] + 1
		}

		units += 1; units /= 2

		if units > 1e3 {
			fmt.Println(n, units )
			break
		}
	}

}
/*----------------------------------------------------------------------------*/
func getFactors (num int) []int {

	sqrnum := int(math.Sqrt(float64(num)))
	
	vector := make([]int, 0)

	for i := 0; primes[i] <= sqrnum; i++ {

		if num % primes[i] == 0 {
			vector = append(vector, 0)
		}

		for num % primes[i] == 0 {
			vector[ len(vector)-1 ]++
			num /= primes[i]
		}
	}

	return vector
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit int) []int  {

    var i, x, y, n int

    sqr_lim := int(math.Sqrt(float64(limit)))
    sieve_nums := make([]bool, limit+1)
	primes := make([]int, 0)


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
/*----------------------------------------------------------------------------*/
