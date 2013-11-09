package main

import (
    "fmt"
	"strconv"
	"strings"
)

func main() {

	const LIMIT = 1e4

	nummap := make(map[int64]bool, 0)

	var n, finsum int64

	finsum = 0

	for n = 1; n < LIMIT; n++ {
		if is_pandigital(n) {
			nummap[n] = true
			fmt.Println(n)
		}
	}

	for k, _ := range nummap {
		finsum += k
	}
	
	fmt.Println(nummap, finsum);
	
}
/*-----------------------------------------------------------------------------*/
func is_pandigital(n int64) bool {

	var i, j int64

	for i = 1; i < n/2; i++ {
		for j = 1; j < n/2; j++ {
			if i * j > n {
				break
			}
			if i * j == n {
				if is_satisfy(i, j, n) {
					return true
				}
			}
		}
	}
	return false
}
/*-----------------------------------------------------------------------------*/
func is_satisfy(i, j, n int64) bool {
	str := strconv.FormatInt(i, 10) + strconv.FormatInt(j, 10) + strconv.FormatInt(n, 10);
	if len(str) != 9 {
		return false
	}
	for b := '1'; b <='9'; b++ {
		if strings.IndexRune(str, b) == -1 {
			return false
		}
	}
	return true
}
