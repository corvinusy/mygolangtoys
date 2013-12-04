package main

import (
    "fmt"
	"math/big"
)

/*
 * tileng n by m
 * tiles = n!/(k1!*k2!...km!)
 */

const SIZE = 50

func main() {

	res := big.NewInt(0)

	for i4 := int64(SIZE/4); i4 >= 0 ; i4-- {
		for i3 := int64((SIZE - i4*4)/3); i3 >= 0 ; i3-- {
			for i2 := int64((SIZE - i4*4 - i3*3)/2); i2 >= 0 ; i2-- {
				i1 := SIZE - i4*4 - i3*3 - i2*2
				res.Add(res, comb(i4+i3+i2+i1, i4, i3, i2, i1))
			//	fmt.Println( i4, i3, i2, i1, "\tcomb =", comb(i4+i3+i2+i1, i4, i3, i2, i1))
			}
		}
	}

	fmt.Println(res)
	
}
/*----------------------------------------------------------------------------*/
func comb(n, m4, m3, m2, m1 int64) *big.Int {

	z := new(big.Int)
	z.MulRange(1, n)

	z1 := big.NewInt(1)
	if m1 != 0 {
		z1.MulRange(1, m1)
	}

	z2 := big.NewInt(1)
	if m2 != 0 {
		z2.MulRange(1, m2)
	}

	z3 := big.NewInt(1)
	if m3 != 0 {
		z3.MulRange(1, m3)
	}

	z4 := big.NewInt(1)
	if m4 != 0 {
		z4.MulRange(1, m4)
	}

	z.Div(z, z1).Div(z, z2).Div(z, z3).Div(z, z4)

	return z

}
/*----------------------------------------------------------------------------*/
