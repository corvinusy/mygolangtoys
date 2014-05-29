package main

import (
	"fmt"
	"time"
)

func main() {

	const LIMIT = 1e4

	var (
		n, cube, digits uint64
	)

	t1 := time.Now()

	cubes := make(map[uint64]uint64, 0)
	mincubes := make(map[uint64]uint64, 0)

	for n = 1; n <= LIMIT; n++ {
		cube = n * n * n
		digits = sort_digits(cube)
		if cubes[digits] == 0 {
			mincubes[digits] = cube
		}
		cubes[digits]++
		if cubes[digits] > 4 {
			t2 := time.Since(t1)
			fmt.Printf("result = %d \ntime = %v\n", mincubes[digits], t2)
			break
		}
	}

}

/*-----------------------------------------------------------------------------*/
func sort_digits(n1 uint64) uint64 {
	ds := [...]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for n1 > 0 {
		ds[n1%10]++
		n1 /= 10
	}

	var res uint64 = 0

	for i := len(ds) - 1; i >= 0; i-- {
		res = res*10 + ds[i]
	}
	return res
}
