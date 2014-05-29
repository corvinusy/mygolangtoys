package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {

	const LIMIT = 100

	var (
		i, j int64
	)

	r := big.NewRat(1, 1)

	for i = 11; i < LIMIT; i++ {
		if i%10 == 0 {
			continue
		}
		for j = i + 1; j < LIMIT; j++ {
			if j%10 == 0 {
				continue
			}
			if is_satisfies(i, j) {
				fmt.Println(i, "/", j)
				r.Mul(r, big.NewRat(i, j))
			}
		}
	}
	fmt.Println(r)
}

/*-----------------------------------------------------------------------------*/
func is_satisfies(i, j int64) bool {
	numstr := strconv.FormatInt(i, 10)
	denstr := strconv.FormatInt(j, 10)

	var r1, r2 *big.Rat

	switch {
	case !strings.ContainsAny(numstr, denstr),
		numstr[0] == numstr[1],
		denstr[0] == denstr[1]:
		return false
	}

	r1 = big.NewRat(i, j)

	switch {
	case numstr[0] == denstr[0]:
		r2 = big.NewRat(int64(numstr[1]-'0'), int64(denstr[1]-'0'))
	case numstr[0] == denstr[1]:
		r2 = big.NewRat(int64(numstr[1]-'0'), int64(denstr[0]-'0'))
	case numstr[1] == denstr[0]:
		r2 = big.NewRat(int64(numstr[0]-'0'), int64(denstr[1]-'0'))
	case numstr[1] == denstr[1]:
		r2 = big.NewRat(int64(numstr[0]-'0'), int64(denstr[0]-'0'))
	default:
		return false
	}

	if r1.Cmp(r2) == 0 {
		return true
	}
	return false
}
