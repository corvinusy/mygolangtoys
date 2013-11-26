package main

import (
    "fmt"
)

func main() {

	const LIMIT = 100000

	var (
		n uint
	)

	smap := make(map[uint]bool, 0)
	pmap := make(map[uint]bool, 0)


	for n = 1; n <= LIMIT; n++ {
		smap[sn(n)] = true
		pmap[pn(n)] = true
		if smap[tn(n)] && pmap[tn(n)] {
			fmt.Println(tn(n))
		}
	}

}
/*-----------------------------------------------------------------------------*/
func pn(n uint) uint {
	return (n * (3 * n - 1)) >> 1
}
/*-----------------------------------------------------------------------------*/
func tn(n uint) uint {
	return (n * (n + 1)) >> 1
}
/*-----------------------------------------------------------------------------*/
func sn(n uint) uint {
	return n * (2 * n - 1)
}
