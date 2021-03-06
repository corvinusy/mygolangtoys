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
		i    int64
		size int64
	)

	primes := make([]int64, 0, 6e7)

	time1 := time.Now()

	size = create_primes(1e7, &primes)

	time2 := time.Since(time1)
	fmt.Println("primes created for ", time2)

	for i = size - 1; i > 0; i-- {
		if is_pandigit(primes[i]) {
			fmt.Println(primes[i])
			break
		}
	}

	time3 := time.Since(time1)
	fmt.Println("pandigit prime found for ", time3)

}

/*-----------------------------------------------------------------------------*/
func create_primes_atkin(limit int64, primes *([]int64)) int64 {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n, count int64

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

	count = 0
	for i = 0; i <= limit; i++ {
		if sieve_nums[i] {
			*primes = append(*primes, i)
			count++
		}
	}

	return count
}

/*-----------------------------------------------------------------------------*/
func is_pandigit(num int64) bool {

	str := strconv.FormatInt(num, 10)

	for i, _ := range str {
		if strings.IndexByte(str, '1'+byte(i)) < 0 {
			return false
		}
	}
	return true

}
