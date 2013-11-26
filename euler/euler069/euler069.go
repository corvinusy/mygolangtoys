package main

import (
    "fmt"
	"time"
)

const LIMIT = 1e6 + 1

func main() {

	var (
		ratio, res_ratio float64 = 1, 1
		res_n int = 1
	)

	var tots [LIMIT+1]int

	t1 := time.Now()

	// prepare totient_sieve
	for i := 1; i <= LIMIT ; i++ {
		tots[i] = i
	}

	for i := 1; i <= LIMIT; i++ {
		for j := 2 * i; j < LIMIT; j += i {
			tots[j] -= tots[i]
		}
	}

	// find best ratio
	for i, n := range tots {
		ratio = float64(i) / float64(n)
		if ratio > res_ratio {
			res_ratio = ratio
			res_n = i
		}
	}

	t2 := time.Since(t1)

	fmt.Println(res_n, res_ratio, t2)

	return
}
/*-----------------------------------------------------------------------------*/
