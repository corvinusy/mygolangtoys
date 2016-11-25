package main

import (
	"fmt"
)

func main() {

	var t, count int
	fmt.Scan(&t)
	inp := make([]int, t)
	for i := range inp {
		fmt.Scan(&inp[i])
	}

	for i := range inp {
		count = 0
		for k := 11; k < inp[i]; k += 2 {
			if isOddful(k + reverse(k)) {
				count += 2
			}
		}
		fmt.Println("ans =", count)
	}
}

/*-----------------------------------------------------------------------------*/
func reverse(n int) int {
	var m int
	for n > 0 {
		m *= 10
		m = m + n%10
		n /= 10
	}

	return m
}

/*-----------------------------------------------------------------------------*/
func isOddful(n int) bool {
	for n > 0 {
		if n&1 == 0 {
			return false
		}
		n /= 10
	}
	return true
}
