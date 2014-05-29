package main

import (
	"fmt"
	"github.com/cznic/mathutil"
)

const LIMIT = 1e8

func main() {

	var ftslen int

	count := 0

	for i := 4; i < LIMIT; i++ {

		fts := mathutil.FactorInt(uint32(i))
		ftslen = len(fts)

		switch ftslen {
		case 2:
			if fts[0].Power == 1 && fts[1].Power == 1 {
				count++
			}

		case 1:
			if fts[0].Power == 2 {
				count++
			}
		}

	}

	fmt.Println(count)

}
