package main

import (
    "fmt"
)

func main() {
	
	const LIMIT = 1e6

	var n, sublim uint64

	sublim = 17

	for n = 1; n < LIMIT; next(&n, &sublim) {
		if is_satisfies(n, n * 2) && is_satisfies(n, n * 3) &&
			is_satisfies(n, n * 4) && is_satisfies(n, n * 5) && 
			is_satisfies(n, n * 6) {
				fmt.Println(n)
				return
		}
	}
}
/*-----------------------------------------------------------------------------*/
func next(n, sublim *uint64) {
	if *n == *sublim {
		*sublim *= 10
		*n = *sublim - uint64(*n * 70 / 17)
	} else {
	*n++
}
}
/*-----------------------------------------------------------------------------*/
func is_satisfies(n, m uint64) bool {

	d := make(map[uint64]int,0)

	for n > 0 {
		d[n % 10] = 1
		n /= 10
	}

	for m > 0 {
		if d[m % 10] != 1 {
			return false
		} else {
			d[m % 10]++
			m /=10
		}
	}

	for _, v := range d {
		if v == 1 {
			return false
		}
	}
	return true
}
