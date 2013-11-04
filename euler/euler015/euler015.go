package main

import (
    "fmt"
	"math/big"
)

func main() {

	z := big.NewInt(0)
	z = z.Binomial(40, 20);

	fmt.Println(z.String())

}


