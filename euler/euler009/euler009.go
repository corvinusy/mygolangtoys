package main

import (
	"fmt"
)

func main () {

	var a, b, c int

	for a = 1; a < 999; a++ {
		for b = 1; b < 999; b++ {
			for c = 1; c < 999; c++ {
				if (a + b + c == 1000) && (((a*a) + (b*b)) == (c*c)) {
					fmt.Println(a*b*c)					
					return
				}
			}
		}
	}

}
