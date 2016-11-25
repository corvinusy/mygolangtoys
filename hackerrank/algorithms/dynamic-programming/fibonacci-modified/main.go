package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	t := new(big.Int)
	u := new(big.Int)
	fmt.Scan(t, u, &n)
	x := big.NewInt(0)
	for i := 2; i < n; i++ {
		x = x.Mul(u, u)
		x = x.Add(x, t)
		t.Set(u)
		u.Set(x)
	}
	fmt.Print(u.String())
}
