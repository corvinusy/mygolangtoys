package main

import (
    "fmt"
	"time"
)

const LIMIT = 1e6

func main() {

	var count uint64 = 0

	var tots[LIMIT + 1]uint64

	t1 := time.Now()

	// prepare totient_sieve
	for i := 1; i <= LIMIT ; i++ {
		tots[i] = uint64(i)
	}

	for i := 1; i <= LIMIT; i++ {
		for j := 2 * i; j <= LIMIT; j += i {
			tots[j] -= tots[i]
		}
	}

	// counting parts
	for i := 2; i <= LIMIT; i++ {
		count += tots[i] 
	}


	t2 := time.Since(t1)
	fmt.Println("LIMIT = ", LIMIT, "result = ", count, "\ntime =", t2)
}
