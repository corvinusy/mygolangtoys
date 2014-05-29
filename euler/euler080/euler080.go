package main

import (
	"fmt"
	"math"
	"math/big"
)

/*-----------------------------------------------------------------------------*/
/*
Newton algorythm
*/
/*-----------------------------------------------------------------------------*/
const LIMIT = 1e2

func main() {

	sum := 0
	var n int64

	for n = 1; n <= 100; n++ {
		sqn := int64(math.Sqrt(float64(n)))
		if sqn*sqn == n {
			continue
		}
		sum += sum100digits(get_sqrt(n))
	}

	fmt.Println(sum)

}

/*-----------------------------------------------------------------------------*/
func get_sqrt(n int64) *big.Rat {

	zn := big.NewRat(n, 1)
	eps := big.NewRat(1, 1e14) //14
	eps.Mul(eps, eps)          //28
	eps.Mul(eps, eps)          //56
	eps.Mul(eps, eps)          //112

	x0 := big.NewRat(1, 1)

	z := new(big.Rat)

	for {
		x1 := new(big.Rat)
		x1.Inv(x0)
		x1.Mul(x1, zn)
		x1.Add(x1, x0)
		x1.Mul(x1, big.NewRat(1, 2))

		z.Mul(x1, x1)
		z.Abs(z.Sub(z, zn))

		if z.Cmp(eps) == -1 {
			return x1
		} else {
			x0 = x1
		}
	}
}

/*-----------------------------------------------------------------------------*/
func sum100digits(z *big.Rat) int {
	s := z.FloatString(105)
	sum := 0
	count := 0
	for i := 0; count < 100; i++ {
		if s[i] != '.' {
			sum += int(s[i] - '0')
			count++
		}
	}
	return sum
}
