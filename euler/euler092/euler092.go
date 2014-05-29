package main

import (
	"fmt"
)

const LIMIT = 1e7

func main() {

	count89 := 0

	for i := 1; i < LIMIT; i++ {
		if get_stuck(i) == 89 {
			count89++
		}
	}

	fmt.Println(count89)
}

/*-----------------------------------------------------------------------------*/
func get_stuck(n int) int {

	for {
		if n == 89 || n == 1 {
			return n
		}
		tmp := 0
		for n > 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		n = tmp
	}

}
