package main

import (
	"fmt"
	"math/big"
)

func main() {

	const LIMIT = 100

	var (
		i, j int64
	)

	nummap := make(map[string]bool, 0)

	bigI := new(big.Int)
	bigJ := new(big.Int)

	for i = 2; i <= LIMIT; i++ {
		bigI.SetInt64(i)
		for j = 2; j <= LIMIT; j++ {
			bigJ.SetInt64(j)
			num := new(big.Int)
			num.Exp(bigI, bigJ, nil)
			nummap[num.String()] = true
		}
	}
	fmt.Println(len(nummap))
}
