package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	var (
		first int
		delta int64
	)
	primes := make([]int64, 0, 1e6)
	time1 := time.Now()

	start := 0
	end := 1

	prime_list(2e6, &primes)

	for i := 4; len(strconv.FormatInt(primes[i], 10)) < 4; i++ {
		start = i + 1
	}

	for i := start; len(strconv.FormatInt(primes[i], 10)) < 5; i++ {
		end = i + 1
	}

	fmt.Println(start, end)

	for i := start; i < end; i++ {
		if (primes[i] != 1487) && (is_progression(i, end, &primes, &delta)) {
			first = i
			fmt.Println(primes[first], primes[first]+delta, primes[first]+delta*2)
		}
	}

	time2 := time.Since(time1)
	fmt.Println(primes[first], primes[first]+delta, primes[first]+delta*2, "\n", time2)
}

/*----------------------------------------------------------------------------*/
func prime_list(limit int64, primes *([]int64)) {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n int64

	for i = 5; i <= limit; i++ {
		sieve_nums[i] = false
	}

	sieve_nums[2] = true
	sieve_nums[3] = true

	for x = 1; x <= sqr_lim; x++ {
		for y = 1; y <= sqr_lim; y++ {

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

/*----------------------------------------------------------------------------*/
func is_progression(start int, end int, primes *([]int64), delta *int64) bool {

	var (
		str1, str2, str3 string
	)

	for i := start + 1; i < end; i++ {
		*delta = (*primes)[i] - (*primes)[start]
		for j := i; j < end; j++ {

			if ((*primes)[j] - (*primes)[i]) == *delta {
				str1 = strconv.FormatInt((*primes)[start], 10)
				str2 = strconv.FormatInt((*primes)[start]+*delta, 10)
				str3 = strconv.FormatInt((*primes)[start]+(*delta)*2, 10)

				if strings.Contains(str2, string(str1[0])) && strings.Contains(str3, string(str1[0])) &&
					strings.Contains(str2, string(str1[1])) && strings.Contains(str3, string(str1[1])) &&
					strings.Contains(str2, string(str1[2])) && strings.Contains(str3, string(str1[2])) &&
					strings.Contains(str2, string(str1[3])) && strings.Contains(str3, string(str1[3])) {
					return true
				}
			}
		}
	}
	return false
}
