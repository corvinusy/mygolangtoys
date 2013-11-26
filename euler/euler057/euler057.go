package main

import (
    "fmt"
	"math/big"
)

func main() {
	
	const LIMIT = 1000
	
	z := big.NewRat(3, 2)

	count := 0

	for i := 2; i < LIMIT; i++ {

		// z = 1 + 1/(1 + z)

		z.Add(z, big.NewRat(1, 1))
		z.Inv(z)
		z.Add(z, big.NewRat(1, 1))

		if len(z.Num().String()) > len(z.Denom().String()) {
			count++
		}
		
	}

	fmt.Println(count)
}
