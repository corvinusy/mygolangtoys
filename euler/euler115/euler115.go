package main

import (
    "fmt"
)


const LIMIT = 1e6
const TSIZE = 50

func main() {

	for llen := TSIZE; llen <= TSIZE*10; llen++ {

		cache := make(map[int]int)

		if comb(TSIZE, llen, cache) > LIMIT {
			fmt.Println("TSIZE =", TSIZE, ", LLEN = ", llen)
			break
		}
	}

}
/*----------------------------------------------------------------------------*/
func comb(n, m int, cache map[int]int) int {

	if cache[m] != 0 {
		return cache[m]
	}

	result := 1

	if m < n {
		return 1
	}

	for pos := 0; pos <= m-n; pos++ {
		for blen := n; blen <= m-pos; blen++ {
			result += comb(n, m - pos - blen - 1, cache)
		}
	}

	cache[m] = result

	return result
}
