package main

import (
	"fmt"
)

func main() {

	const LIMIT = 1000

	var (
		i, j, p int
	)

	rescount := 0
	resp := 1

	for p = 1; p <= LIMIT; p++ {
		count := 0
		cmap := make(map[int]bool, 0)
		for i = 1; i < p; i++ {
			if cmap[i] {
				continue
			}
			for j = 1; j < p; j++ {
				if cmap[j] {
					continue
				}
				if i+j >= p {
					break
				}
				if is_satisfies(i, j, p) {
					count++
					cmap[i] = true
					cmap[j] = true
				}
			}
		}
		if count > rescount {
			rescount = count
			resp = p
		}
	}
	fmt.Println(rescount, resp)
}

/*-----------------------------------------------------------------------------*/
func is_satisfies(i, j, p int) bool {
	sqr := (p - i - j) * (p - i - j)
	return sqr == i*i+j*j
}
