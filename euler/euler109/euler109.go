package main

import (
	"fmt"
)

const (
	curLimit = 100
)

func main() {

/*
	fmt.Println("checkouts = ", getScoreCheckouts(7))
*/
	checkouts := 0
	var co int

	for n := 1; n < curLimit; n++ {
		co = getScoreCheckouts(n)
		checkouts += co
		fmt.Printf("num = %3d, checkouts = %d\n", n, co)
	}

	fmt.Println("checkouts = ", checkouts)
}
/*----------------------------------------------------------------------------*/
func isFirstValidDart(n int) bool {
	if n == 25 || n == 50 {
		return true
	} 
	
	if n <= 20 {
		return true
	}

	if n <= 40 && n % 2 == 0 {
		return true
	}

	if n <= 60 && n % 3 == 0 {
		return true
	}

	return false
}
/*----------------------------------------------------------------------------*/
func isLastValidDart(n int) bool {

	if n == 0 {
		return false
	}

	if n == 50 {
		return true
	} 
	
	if n <= 40 && n % 2 == 0 {
		return true
	}

	return false
}
/*----------------------------------------------------------------------------*/
func isDouble(n int) bool {
	
	if n == 0 || n > 40 {
		return false
	}

	return (n % 2) == 0
}
/*----------------------------------------------------------------------------*/
func isTriple(n int) bool {
	
	if n == 0 || n > 60 {
		return false
	}

	return (n % 3) == 0
}
/*----------------------------------------------------------------------------*/
func dartVariants(n int) int {

	if n == 0 {
		return 1
	}

	if n == 25 || n == 50 {
		return 1
	}

	if n > 20 && isDouble(n) && isTriple(n) {
		return 2
	}

	if n > 20 && (isDouble(n) || isTriple(n)) {
		return 1
	}

	if n <= 20 && isDouble(n) && isTriple(n) {
		return 3
	}

	if n <= 20 && (isDouble(n) || isTriple(n)) {
		return 2
	}

	return 1
}
/*----------------------------------------------------------------------------*/
func getScoreCheckouts(n int) int {

	result := 0

	for i1 := 0; i1 < n; i1++ {

		if !isFirstValidDart(i1) {
			continue
		}

		for i2 := i1; i2 < n - i1 ; i2++ {

			if !isFirstValidDart(i2) {
				continue
			}

			i3 := n - i2 - i1

			if !isLastValidDart(i3) {
				continue
			}

			if i1 != i2 {
				result += dartVariants(i1) * dartVariants(i2)
			} else {
				result += dartVariants(i1)  + dartVariants(i1)*(dartVariants(i1)-1) / 2
			}

		}
	}

	return result
}
/*----------------------------------------------------------------------------*/
/*
6+6+2 : 3 => 6
s6+s6+d1
s6+d3+d1 =1
s6+t2+d1 =2
d3+s6+d1 =1
d3+d3+d1
d3+t2+d1 =3
t2+s6+d1 =2
t2+d3+d1 =3
t2+t2+d1

4+4+2 : 2 => 3
s4+s4+d1
s4+d2+d1 =1
d2+d2+d1
d2+s4+d1 =1

5+5+2 : 1 => 1
s5+s5+d2
*/
