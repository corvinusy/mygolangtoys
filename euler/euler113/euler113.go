package main

import (
    "fmt"
	"math/big"
)

// Binomial(n+9,n) + Binomial(n+10,n) - 10*n -1

const LIMIT = 100

func main() {

	z1 := new(big.Int)
	z2 := new(big.Int)

	z1.Binomial(LIMIT+9, LIMIT)
	z2.Binomial(LIMIT+10, LIMIT)

	z1.Add(z1, z2)

	z2.SetInt64(LIMIT*10)

	z1.Sub(z1, z2)
	z1.Sub(z1, big.NewInt(1))

	fmt.Println(z1)
}
