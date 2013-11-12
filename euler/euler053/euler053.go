package main

import (
    "fmt"
	"math/big"
)

func main() {

	var i, j, count int64
	const LIMIT = 100

	z := new (big.Int)
	count = 0

	for i = 1; i <= LIMIT; i++ {
		for j = i ; j <= LIMIT; j++ {
			z.Binomial(j, i);
			if len(z.String()) > 6 {
				count += LIMIT - j + 1
				break
			}
		}
	}


	fmt.Println(count)

}


