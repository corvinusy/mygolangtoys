package main

import (
    "fmt"
	"github.com/cznic/mathutil"
)

const LIMIT = 10
const MINLIMIT = 1e9

func main() {

	var (
		d, r, n uint64
	)

	rmap := make(map[uint64]uint64)

	sum := uint64(0)

	for r = LIMIT-1; r > LIMIT-3; r-- { // r - number of digit repeats
		for d = 0; d < 10; d++ { // current digit
			smap := make(map[uint64]bool)
			masks := formMasks(r) //maskslice like 1111111110
			for _, mask := range masks {
				numbers := formNumbers(mask, r, d)
				for _, n = range numbers { 
					if mathutil.IsPrimeUint64(n) {
						if rmap[d] == 0 {
						 rmap[d] = r
						}
						if rmap[d] == r {
							fmt.Println(r, d, n)
							smap[n] = true
						}
					}
				}
			}

			for n := range smap {
				sum += n
			}
		}
	}

	fmt.Println("sum = ", sum)
}

/*----------------------------------------------------------------------------*/
func formMasks(r uint64) [][LIMIT]uint8 {

	a := make([][LIMIT]uint8, 0)

	var m, count uint64
	
	for m = 0; m <= 1 << LIMIT; m++ {
		q := m
		var aa [LIMIT]uint8
		count = 0
		for i := range aa {
			if q & 1 != 0 {
				count += 1
				aa[i] = 1
			}
			q = q >> 1
		}
		if count == r {
			a = append(a, aa)
		}
	}

	return a
}
/*----------------------------------------------------------------------------*/
func formNumbers(mask [LIMIT]uint8, r, d uint64) []uint64 {
	
	var k, n uint64
	
	a := make([]uint64, 0)

	maxnum := 1

	for ;r < LIMIT; r++ {
		maxnum *= 10
	}
	
	minnum := maxnum / 10
	if minnum == 1 { minnum = 0 }

	for i := minnum; i < maxnum; i++ {
		isContainD := false
		k = uint64(i)
		for k > 0 {
			if k % 10 == d {
				isContainD = true
				break 
			}
			k /= 10
		}

		if isContainD {continue}

		k = uint64(i)
		n = 0
		for i := range mask {
			if mask[i] == 1 {
				n = n * 10 + d
			} else {
				n = n * 10 + k % 10
				k /= 10
			}
		}

		if n > MINLIMIT {
			a = append(a, n)
		}
	}

	return a

}
/*-----------------------------------------------------------------------------*/
