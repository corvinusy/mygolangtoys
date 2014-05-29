package main

import (
	"fmt"
	"github.com/cznic/mathutil"
	"math/big"
)

const LIMIT = 51

func main() {

	var (
		n, k uint64
	)

	cache := makeCache(LIMIT - 1)

	fmt.Println("cache prepared")

	umap := make(map[uint64]bool)

	usum := uint64(0)

	for n = 1; n < LIMIT; n++ {
		for k = 1; k <= n; k++ {

			q := binomial(n, k)

			if umap[q] {
				continue
			}

			umap[q] = true

			if !isSquareFree(q, cache) {
				continue
			}

			usum += q

		}
	}

	fmt.Println("limit =", LIMIT, "sum = ", usum)

}

/*----------------------------------------------------------------------------*/
func makeCache(limit uint64) []uint64 {

	cache := make([]uint64, 0)

	for i := uint64(2); i*i <= binomial(limit, limit/2+1); i, _ = mathutil.NextPrimeUint64(i) {
		cache = append(cache, i*i)
	}

	return cache
}

/*----------------------------------------------------------------------------*/
func binomial(n, k uint64) uint64 {

	z := new(big.Int)

	z.Binomial(int64(n), int64(k))

	res, _ := mathutil.Uint64FromBigInt(z)

	return res

}

/*----------------------------------------------------------------------------*/
func isSquareFree(n uint64, cache []uint64) bool {

	for _, ii := range cache {
		if n%ii*ii == 0 {
			return false
		}
	}

	return true
}
