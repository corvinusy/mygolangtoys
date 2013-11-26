package main

import (
	"fmt"
)

func main() {
	
	start := 2*3*5*7*11*13*17*19

	for i := start; i < 4e9; i += start {
		if is_good(i) {
			fmt.Println(i)
			break;
		}
	}
}

func is_good(p int) bool {

	for i := 4; i <= 20; i++ {
		if p % i !=0 {
			return false;
		}
	}
	return true;
}

