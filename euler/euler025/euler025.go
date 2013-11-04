package main

import (
    "fmt"
	"math/big"
)

func main() {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	ftmp := big.NewInt(1)

    i := 1

	for len(f1.String()) < 1000 {
		ftmp =  f1
		f1 = f1.Add(f1, f2)
		f2 = ftmp
		i++
	}
	
	fmt.Println(i, " ", f1.String())
	
}

