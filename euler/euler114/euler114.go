package main

import (
    "fmt"
	"math/big"
)

/*
 * tileng n by m
 * tiles = n!/(k1!*k2!...km!)
 */

const SIZE = 7
const TSIZE = 3

func main() {

	res := big.NewInt(0) // 1 == all black

	for i4 := int64(TSIZE); i4 <= SIZE ; i4++ {
		m := make([]int64, 1)

		for i3 := i4; i3 >= TSIZE; i3-- {
			m = append(m, 0)

			for i := 1; i < len(m); i-- {
				
			}
			
			res.Add(res, comb(i3+i2, i3, i2))

//			fmt.Println("size = ", i4, i3, i2, "\tcomb =", comb(i3+i2, i3, i2))
		}
	}

	fmt.Println(res)
	
}
/*----------------------------------------------------------------------------*/
func comb(n, m []int64) *big.Int {

	z := new(big.Int)
	z.MulRange(1, n)

	for i, v := range m {
		if v != 0 {
			z1 := big.NewInt(1)
			z1.MulRange(1, v)
			z.Div(z, z1)
		}
	}

	z.Div(z, z1).Div(z, z2)

	return z

}
/*----------------------------------------------------------------------------*/
