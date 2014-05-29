package main

import (
	"fmt"
)

const LIMIT = 1e9

func main() {

	n := 0
	count := 0

	for i := 1; i < LIMIT; i++ {
		if i%10 == 0 {
			continue
		}
		n = i + reverse(i)
		if is_oddful(n) {
			count++
			//			fmt.Println (i, reverse(i), n)
		}
	}
	fmt.Println(count)

}

/*-----------------------------------------------------------------------------*/
func reverse(n int) int {

	var ds [10]int

	nl := 0

	for n > 0 {
		ds[nl] = n % 10
		n /= 10
		nl++
	}

	n = 0

	for i := 0; i < nl; i++ {
		n *= 10
		n += ds[i]
	}

	return n

}

/*-----------------------------------------------------------------------------*/
func is_oddful(n int) bool {

	for n > 0 {
		if n&1 == 0 {
			return false
		}
		n /= 10
	}
	return true
}
