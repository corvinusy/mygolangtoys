package main

import (
    "fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {

	primes := make([]int64, 0)
	create_primes(1000, &primes)

	r := new(big.Rat)

	var count int

	rescount := 0
	resprime := primes[5]

	for i := 5; i < len(primes); i++ {
		r.SetString("1/" + strconv.FormatInt(primes[i], 10))
//		fmt.Println(primes[i], r.FloatString(150)[3:])
		if i < 100 {
			count = find_period(r.FloatString(1e5)[3:])
		} else {
			count = find_period(r.FloatString(1e5)[4:])			
		}

		if count > rescount {
			rescount = count
			resprime = primes[i]
		}
	}
	fmt.Println(resprime, count)
}
/*-----------------------------------------------------------------------------*/
func find_period(s string) int {
	
	lens := len(s)/2

	for i:=0; i < 20; i++ {
		for j := 35 + i; j < lens; j++ {
			if s[j] != s[i] {
				continue
			}
			if strings.Count(s, s[i:j]) == len(s)/(j - i) {
				return j - i
			}

		}
	}
	return 0
}
/*-----------------------------------------------------------------------------*/
func create_primes (limit int64, primes *[]int64) {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n int64;

	for i = 5; i <= limit ; i++ {
		sieve_nums[i] = false;
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
			*primes = append(*primes, i)
		}
	}
	return
}
