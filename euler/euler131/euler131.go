package main

import (
    "fmt"
	"github.com/cznic/mathutil"
	"math/big"
)

const PRIME_LIMIT = 1e6
const NUM_LIMIT = 1e3

func main() {

	var ( 
		i uint64
		previous, j int64
	)

	cubemap := prepareCubes(PRIME_LIMIT*1e3)

//	fmt.Println(cubemap)

	count := 0
	prime_count := -1
	previous = 1

	z1 := new(big.Int)
	z2 := new(big.Int)

	for i = 1; i < PRIME_LIMIT; i, _ = mathutil.NextPrimeUint64(i) {
		prime_count++
		for j = previous; j < NUM_LIMIT; j++ {

			z1.SetInt64(j)
			z2.Mul(z1, z1)
			z2.Mul(z2, z1)
			
			z1.Set(z2) // z1 = j*j*j

			z2.Mul(z1, z1)
			z2.Mul(z2, z1) //z2 = z1*z1*z1

			z1.Mul(z1, z1)
			z1.Mul(z1, big.NewInt(int64(i))) 
			z1.Add(z1, z2) //z1 = j*j*j + j*j*i

			if cubemap[z1.String()] {
				fmt.Println(prime_count,":", i, j, z1)
				previous = j
				count++
				break
			}
		}
	}
	
	fmt.Println(count)
}
/*----------------------------------------------------------------------------*/
func prepareCubes(lim int64) map[string]bool {

	var i int64

	cubemap := make(map[string]bool)

	for i = 1; i * i * (i+1) < lim; i += 1 {

		k := i * i * (i+1)

		z1 := new(big.Int)
		z2 := new(big.Int)

		z1.SetInt64(k)
		z2.Mul(z1, z1)
		z2.Mul(z2, z1)

		cubemap[z2.String()] = true
	}

	return cubemap
}
