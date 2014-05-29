package main

import (
	"fmt"
	"math/big"
)

const LIMIT = 1e6

func main() {

	var i, j int64

	z := big.NewRat(3, 7)
	zres := big.NewRat(2, 5)

	for i = 14; i <= LIMIT; i++ {
		for j = i / zres.Denom().Int64() * zres.Num().Int64(); j < i; j++ {
			z1 := big.NewRat(j, i)
			if z.Cmp(z1) <= 0 {
				break
			}
			if z1.Cmp(zres) == 1 {
				zres = z1
				//				fmt.Println(zres.String())
			}
		}
	}
	fmt.Println("result = ", zres.String())
}
