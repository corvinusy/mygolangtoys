package main

import "fmt"

const limit = 200

var chainLength [limit + 1]int
var chain [limit + 1]int

func additionChain(exponent, length int) {
	if exponent > limit || length > chainLength[exponent] {
		return
	}

	chainLength[exponent] = length
	chain[length] = exponent

	for i := length; i >= 0; i-- {
		additionChain(exponent+chain[i], length+1)
	}
}

func main() {

	//init arrays
	for i := 0; i < limit+1; i++ {
		chainLength[i] = limit // big number, which definitely greater than calculated
		chain[i] = 0
	}

	// recursive call
	additionChain(1, 0)

	// get result
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += chainLength[i]
	}

	// printres

	fmt.Println(sum)
}
