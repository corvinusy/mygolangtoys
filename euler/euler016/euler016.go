package main

import (
    "fmt"
	"math/big"
	"strconv"
)

func main() {

	z := big.NewInt(1);
	z = z.Lsh(z, 1000);

	s2k := z.String();

	var temp, sum int64 = 0, 0

	for _, s := range s2k {
		temp, _ = strconv.ParseInt(string(s), 10, 0)
		sum = sum + temp
	}
	
	fmt.Println(sum)

}


