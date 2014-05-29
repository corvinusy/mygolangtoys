package main

import (
	"fmt"
	"math"
)

// Sum = n*(n+1)/2 * m*(m+1)/2
// m * (m +1) = 2 * Sum / n / (n+1) * 2 = 4 * Sum / n / (n+1)
// m < math.Sqrt(4 * Sum / n / (n+1)) = 2 * math.Sqrt(Sum / n / (n+1) ))

//LIMIT = 2e6
const LIMIT = 2e6

func main() {

	//	nlim := int64(math.Sqrt(LIMIT))

	result := int64(LIMIT)

	resm := int64(0)
	resn := int64(0)

	for n := int64(1); n <= LIMIT-1; n++ {

		m := 2 * int64(math.Sqrt(float64(LIMIT/n/(n+1))))

		for ; rect_num(n, m-2) < LIMIT; m++ {
			delta := rect_num(n, m) - LIMIT
			if delta < 0 {
				delta = -delta
			}
			//			fmt.Println(n, m, delta)
			if result > delta {
				result = delta
				resn = n
				resm = m
			}
		}
	}

	fmt.Println("result", resn, resm, resn*resm)
}

/*-----------------------------------------------------------------------------*/
func rect_num(n, m int64) int64 {
	return int64(n * (n + 1) * m * (m + 1) / 4)
}
