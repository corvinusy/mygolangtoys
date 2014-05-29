package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	z1 := big.NewInt(1)
	z2 := big.NewInt(1)
	z := new(big.Int)

	for i := 2; ; i++ {
		z.Set(z2)
		z2.Add(z2, z1)
		z1.Set(z)

		str := z.String()
		if i < 500 {
			continue
		}
		if i%1000 == 0 {
			fmt.Println(i)
		}
		if is_pandigital(str[0:9]) && is_pandigital(str[len(str)-9:]) {
			fmt.Println("start + end:", i)
			break
		}
	}

}

/*-----------------------------------------------------------------------------*/
func is_pandigital(s string) bool {
	for c := '1'; c <= '9'; c++ {
		if strings.IndexRune(s, c) == -1 {
			return false
		}
	}
	return true
}
