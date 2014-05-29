package main

import (
	"fmt"
	"math/big"
)

func main() {

	var i, j, zsum, tmp int64
	const LIMIT = 100

	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)
	zsum = 1

	for i = LIMIT - 1; i > 1; i-- {
		for j = LIMIT; j > 1; j-- {
			x.SetInt64(i)
			y.SetInt64(j)
			z.Exp(x, y, nil)
			if len(z.String()) < 100 {
				break
			}
			tmp = getzsum(z.String())
			if tmp > zsum {
				zsum = tmp
			}

		}
	}

	fmt.Println(zsum)
}

/*-----------------------------------------------------------------------------*/
func getzsum(s string) int64 {
	var result int64 = 0
	for _, a := range s {
		result += int64(a - '0')
	}
	return result
}
