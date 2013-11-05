package main

import (
    "fmt"
	"math/big"
)

func main() {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	ftmp := big.NewInt(1)

    i := 2

	for {
		ftmp.Add(f1, f2)
		f2.Set(f1)
		f1.Set(ftmp)
		i++
		if len(f1.String()) == 1000 {
			break
		}
//		fmt.Println(i, " ", f1.String(), " ", f2.String())
	}
	
	fmt.Println(i, " ", f1.String())
	
}

