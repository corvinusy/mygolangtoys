package main

import (
	"fmt"
	"math/big"
)

const LIMIT = 12000

func main() {

	var i, j int64
	var count int64 = 0

	z1 := big.NewRat(1, 3)
	z2 := big.NewRat(1, 2)

	// dirty brute force
	for i = 4; i <= LIMIT; i++ {
		for j = i / 3; j <= i/2; j++ {
			if gcd(i, j) != 1 {
				continue
			} else {
				z := big.NewRat(j, i)
				if (z.Cmp(z1) == 1) && (z.Cmp(z2) == -1) {
					count++
				}
			}
		}
	}
	fmt.Println("LIMIT = ", LIMIT, "result = ", count)
}

/*-----------------------------------------------------------------------------*/
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
