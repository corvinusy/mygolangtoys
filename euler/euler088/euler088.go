package main

import (
    "fmt"
	"github.com/cznic/mathutil"
)

const LIMIT = 49

func main() {

	var i, sum, ss, count uint32

	bigsum := 1

	summap := make(map[uint32]uint32)

	summap[1] = 1

	for i = 2; i <= LIMIT*2; i++ {

		fts := mathutil.FactorInt(uint32(i))

		ss = 0
		count = 0

		for _, ft := range fts {
			ss += ft.Prime * ft.Power //error!
			count += ft.Power
		}

		for j := uint32(len(fts)); j <= count; j++ {
			sum = i - ss + j
			if sum <= LIMIT && summap[sum] == 0 {
				summap[sum] = i
			}
		}
	}

	for _, v := range summap {
		bigsum += int(v)
	}

	fmt.Println(bigsum, summap)

}
/*-----------------------------------------------------------------------------*/
