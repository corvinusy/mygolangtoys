package main

import (
	"fmt"
	"math/big"
)

func main() {

	const (
		LIMIT = 9
	)

	var (
		i, n, d int64
	)

	x := new(big.Int)
	y := new(big.Int)
	z := new(big.Int)

	count := 1

	for i = 2; i <= LIMIT; i++ {
		for n = 1; true; n++ {
			x.SetInt64(i)
			y.SetInt64(n)
			d = int64(len(z.Exp(x, y, nil).String()))
			if n > d {
				break
			}
			if n == d {
				fmt.Println(i, n, d)
				count++
			}
		}
	}

	fmt.Println(count)
}
