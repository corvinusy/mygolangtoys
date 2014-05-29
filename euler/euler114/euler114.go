package main

import (
	"fmt"
)

/*
 * tileng n by m
 * tiles = n!/(k1!*k2!...km!)
 */

const SIZE = 50
const TSIZE = 3

func main() {

	cache := make(map[int]int)

	fmt.Println(comb(TSIZE, SIZE, cache))

}

/*----------------------------------------------------------------------------*/
func comb(n, m int, cache map[int]int) int {

	if cache[m] > 0 {
		return cache[m]
	}

	result := 1

	if m < n {
		return 1
	}

	for pos := 0; pos <= m-n; pos++ {
		for blen := n; blen <= m-pos; blen++ {
			result += comb(n, m-pos-blen-1, cache)
		}
	}

	cache[m] = result

	return result
}
