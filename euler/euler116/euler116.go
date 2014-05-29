package main

import (
	"fmt"
	"math/big"
)

/*
 * tileng n by m
 * tiles = n!/(k1!*k2!)
 */

func main() {

	res := big.NewInt(0)

	for m := int64(2); m <= 4; m++ {
		res.Add(res, calcTiles(50, m))
	}

	fmt.Println(res)

}

/*----------------------------------------------------------------------------*/
func calcTiles(n, m int64) *big.Int {

	res := big.NewInt(0)

	for i := int64(1); m*i <= n; i++ {
		res.Add(res, comb(n-i*m+i, i, n-i*m))
	}

	return res
}

/*----------------------------------------------------------------------------*/
func comb(n, m1, m2 int64) *big.Int {

	z := new(big.Int)
	z.MulRange(m1+1, n)

	z2 := new(big.Int)
	z2.MulRange(1, m2)

	z.Div(z, z2)

	return z

	//	return fact(n)/fact(m1)/fact(m2)

}

/*----------------------------------------------------------------------------*/
