package main

import (
    "fmt"
	"github.com/cznic/mathutil"
)

const LIMIT = 10

func main() {

	summap := make(map[int]int)

	for n := 2; n <= LIMIT*2; n++ {

		fts := mathutil.FactorInt(uint32(n))

		divisors := make([]int, 0)
		
		// get slice of prime  divisors
		for _, ft := range fts {
			for ft.Power > 0 {
				divisors = append(divisors, int(ft.Prime))
				ft.Power--
			}
		}

		divs := getCombos(n, divisors)

		fmt.Println(n, divs)

		for i := range divs {

			product := 1
			for j := range divs[i] {
				product *= divs[i][j]
			}

			sum := 0
			for j := range divs[i] {
				sum += divs[i][j]
			}

			k := product - sum + len(divs[i])

			if k == 10 {
				fmt.Println(i, divs)
			}

			if summap[k] == 0 {
				summap[k] = n
			} 

		}

	}

	fmt.Println(summap)

}
/*-----------------------------------------------------------------------------*/
func getCombos(n int, divs []int) [][]int {

	if len(divs) == 0 {
		return [][]int{}
	}

	combos := make([][]int, 0)	

	for i := range divs {

		a := make([]int, 1)

		if n % divs[i] != 0 {
			continue
		}

		a[0] = n / divs[i]

		newcombos := getCombos(a[0], divs)

		for i := range newcombos {
			for j := range newcombos[i] {
				a = append(a, newcombos[i][j])
			}
		}

		if len(a) > 1 {
			combos = append(combos, a)
		}

	}


	return combos
}
