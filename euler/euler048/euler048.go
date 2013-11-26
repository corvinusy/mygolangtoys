package main

import (
    "fmt"
	"math/big"
)

func main() {
	

	sum := big.NewInt(0)
	f1 := new(big.Int)
	f2 := new(big.Int)

	var i int64

	for i = 1; i <= 1000; i++ {
		f2.SetInt64(i)
		f1.Exp(f2, f2, nil)
		sum.Add(sum, f1)
	}

	fmt.Println(sum.String()[len(sum.String())-10:])
	
}
/*-----------------------------------------------------------------------------*/
