package main

import (
    "fmt"
	"math"
)

const LIMIT = 100

func main() {

	primes := create_primes_atkin(LIMIT*2)

	for i:= 10; i < LIMIT; i++ {
		n := count_rep(i, primes)
		fmt.Println(i, n)
		if n > 5000 {
			break
		}
	}
}
/*-----------------------------------------------------------------------------*/
func count_rep(n int, primes []int) int {

	lenp := 0

	for i:=0; ; i++ {
		if primes[i] > n {
			lenp = i - 1
			break
		}
	}

	c := make([]int, lenp+1)

	/*
     * c = slice of prime_number counts,
     * i.e. c[0] = count of '2', c[1] = count of '3', c[2] = count of '5'...
     */

	count := 0

	//init c

	c[0] = 1

	for !is_c_null(c) {

		if is_c_ok(n, c, primes) {
			count++
		}
		inc(n, c, primes)
	}


	return count
}
/*-----------------------------------------------------------------------------*/
func is_c_null(c []int) bool {
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}
/*-----------------------------------------------------------------------------*/
func is_c_ok(n int, c []int, primes []int) bool {
	for i := 0; i < len(c); i++ {
		n -= c[i]*primes[i]
		if n < 0 {return false}
	}
	return n == 0
}
/*-----------------------------------------------------------------------------*/
func inc(n int, c []int, primes []int) []int {

	for i := 0; i < len(c); i++ {
		c[i]++
		if sumc(c, primes) <= n {
			return c
		}
		c[i] = 0
	}
	return c
}
/*-----------------------------------------------------------------------------*/
func sumc(c []int, primes []int) int {

	result := 0

	for i, _ := range c {
		result += c[i]*primes[i]
	}
	return result
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
