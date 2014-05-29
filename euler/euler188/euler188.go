package main

import (
	"fmt"
	"github.com/cznic/mathutil"
)

// a***1855 = a**(a***1854)

const LIMIT = 1855

const BASE = 1777

func main() {

	fmt.Println(superExpMod(BASE, LIMIT, 1e9))
}

/*----------------------------------------------------------------------------*/
func superExpMod(b, e, m uint64) uint64 {

	if e == 1 {
		return b
	}

	return mathutil.ModPowUint64(b, superExpMod(b, e-1, m), m)
}
