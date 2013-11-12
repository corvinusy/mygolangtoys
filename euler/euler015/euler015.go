package main

import (
    "fmt"
	"math/big"
)

func main() {

	z := new (big.Int)
	z.Binomial(40, 20);

	fmt.Println(z.String())

}


