package main

import (
	"fmt"
	"math/big"
)

/*
1/x + 1/y = 1/n
*/

const LIMIT = 1e5

var pvector []int64

func main() {

	var (
		n int64
		x int64
		count int64
	)

	pvector = append(pvector, 2,3,5,7,11,13,17,19,23,29,31)

	z  := new(big.Rat)
	z1 := new(big.Rat)

	cmap := make(map[int64]int64)

	for x = 0; x < 6; x++  {
		cmap[pvector[x]] = 1
	}

	for {
		cmap = inc(cmap)
		
		n = 1

		for x, _ = range cmap {
			if cmap[x] > 0 {
				n *= cmap[x] * x
			}
		}

		z.SetFrac64(1, n)
		count = 0
		
		for x = n; x <= n * 2; x++ {

			z1.SetFrac64(1, x)
			z1.Sub(z,z1)

			if z1.Num().String() == "1" {
				count++
			}

			if count > 1e4 {
				return
			}
		}
		if count > 1e3 {
			fmt.Println(n, count)				
		}

	}
}
/*-----------------------------------------------------------------------------*/
func inc(cmap map[int64]int64) map[int64]int64 {

	for i:=0; i < 10; i++ {
		if cmap[pvector[i]] < 10 {
			cmap[pvector[i]]++
			return cmap
		} else {
			if cmap[pvector[i]] == 10 {
				cmap[pvector[i]] = 1
				cmap[pvector[i+1]]++
			}
		}
	}
	panic("fuck")
	return nil

}
