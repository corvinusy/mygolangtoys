package main

import (
    "fmt"
)

func main() {

	const LIMIT = 10000

	var (
		n, m, cur uint
	)

	cur = 1e7

	pmap := make(map[uint]bool, 0)

	for n = 1; n <= LIMIT; n++ {
		pmap[pn(n)] = true
	}

	for n = 1; n < LIMIT/2; n++ {
		for m = n + 1; m < LIMIT/2; m++ {
			if pmap[pn(m) - pn(n)] && pmap[pn(m)+pn(n)] {
				fmt.Println(pn(n), pn(m))
				if pn(m) - pn(n) < cur {
					cur = pn(m) - pn(n)
				}
			}
		}
	}
	fmt.Println(cur)
}
/*-----------------------------------------------------------------------------*/
func pn(n uint) uint {
	return (n * (3 * n - 1)) >> 1
}
