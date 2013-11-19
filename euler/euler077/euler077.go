package main

import (
    "fmt"
	"math"
)

// Greedy algorythm
// Conjecture: for any number and prime set exists only one prime partition with coeff > 0
// It can be found by greedy algorythm
// #3# = 0 
// #4# = 1 (2 + 2)
// #5# = 1 (2 + 3)
// #6# = 2 (3*2; 2*3) #4# + 1
// #7# = 2 (2*2+3, 2+5) #5# + 1
// #8# = 3 (4*2, 2+2*3, 5+3) #6# + 1
// #9# = 4 (2*3+3, 2*2+5, 7+2, 3*3)
// #10# = 5 (5*2; 2*2 + 2*3; 2+3+5; 2*5; 3+7)
// #11# =   ()

const LIMIT = 98

func main() {

	primes := create_primes_atkin(LIMIT*2)
	fmt.Println(primes)


	c := make([]int, len(primes))

	for greedy_alg(LIMIT, primes, c) == 0 {
		greedy_alg(LIMIT, primes, c
	}

	fmt.Println(greedy_alg(LIMIT, primes))
}
/*-----------------------------------------------------------------------------*/
func greedy_alg(n_in int, primes []int, c []int) int {

	lenp := 0

	for i:=0; ; i++ {
		if primes[i] > n_in {
			lenp = i - 1
			break
		}
	}

	count := 0

	n := n_in

	for	i := lenp; (n > 0) && (i>=0); i-- {

		c[i] = n / primes[i]

		if n == c[i]*primes[i] {
			count++
			fmt.Println(primes[lenp], c)
		}

		n = n - c[i]*primes[i]	
	}

	return count
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
