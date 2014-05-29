package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var sum int64 = 0
	res := make([]int64, 0, 1e6)
	time1 := time.Now()

	prime_list(2e6, &res)

	for i := 0; i < len(res); i++ {
		sum = sum + res[i]
	}

	time2 := time.Since(time1)
	fmt.Println(sum, "\n", time2)
}

/*func is_prime (n int64) bool {
	var i int64;
	for i = 3; i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n % i == 0 {
			return false;
			break;
		}
	}
	return true;
}
*/
func prime_list(limit int64, res *([]int64)) {

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
			*res = append(*res, i)
		}
	}

	return
}
