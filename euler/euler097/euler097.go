package main

import (
	"fmt"
)

const LIMIT = 7830457

func main() {

	result := int64(1)
	for i := 1; i <= LIMIT; i++ {
		result = (result * 2) % 1e10
	}
	fmt.Println(LIMIT, 28433*result+1)

}
