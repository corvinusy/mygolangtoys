package main

import (
	"fmt"
	"github.com/cznic/mathutil"
)

const LIMIT = 1e7

func main() {

	var (
		dLen, prevDivLen, i int
	)

	count := 0

	prevDivLen = divLen(2)

	for i = 3; i < LIMIT; i, prevDivLen = i+1, dLen {

		dLen = int(divLen32(uint32(i)))

		if dLen == prevDivLen {
			count++
		}

	}

	fmt.Println(count)

}

/*----------------------------------------------------------------------------*/
func divLen32(n uint32) uint32 {

	result := uint32(1)

	fts := mathutil.FactorInt(n)

	for _, ft := range fts {
		result *= (ft.Power + 1)
	}

	return result
}

/*----------------------------------------------------------------------------*/
func divLen(n int) int {

	var i, count int

	count = 0

	for i = 1; i <= n; i++ {
		if n%i == 0 {
			count += 1
		}
	}

	return count
}

/*----------------------------------------------------------------------------*/
