package main

import (
    "fmt"
	"math/big"
	"strconv"
)

func main() {

	product := big.NewInt(1)

	for i:=2; i <=100; i++ {
		product = product.Mul(product, big.NewInt(int64(i)))
	}
	
	str := product.String()

	fmt.Println(str)

	var sum int64 = 0

	for _, s := range str {
		n, _ := strconv.ParseInt(string(s), 10, 0)
		sum += n
	}
	
	fmt.Println(sum)

}


