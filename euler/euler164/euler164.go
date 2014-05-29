package main

// try to bruteforce for low parts

import (
	"flag"
	"fmt"
)

var (
	LIMIT   int
	NUMERAL int
	UPPER   int
)

type vector []int

func init() {
	flag.IntVar(&LIMIT, "lim", 3, "power of 10")
	flag.IntVar(&NUMERAL, "num", 9, "numeral system")
	flag.IntVar(&UPPER, "up", 9, "upper limit")
}

func main() {

	flag.Parse()

	v := make(vector, LIMIT)

	v[LIMIT-1] = 1

	count := 1

	for v.advance() {
		//		fmt.Println(v)
		count++
	}

	fmt.Println("LIMIT =", LIMIT, "NUMERAL =", NUMERAL, "UPPER =", UPPER, "vector count =", count)

	ncount := 0
	oncount := 0

	start := 1

	for i := 0; i < LIMIT-1; i++ {
		start *= 10
	}

	for i := start; i < start*10; i++ {
		if isOk(i) {
			if i%(start/10) == 0 {
				fmt.Println(ncount, ncount-oncount)
				fmt.Print(i/(start/10), ":")
				oncount = ncount

			}
			//			fmt.Print(i, " ")
			/*			if i >= 8000 {
							fmt.Println(i)
						}
			*/
			ncount++
		}
	}

	fmt.Println("number count = ", ncount)

}

/*----------------------------------------------------------------------------*/
func (v vector) advance() bool {

	for i := range v {

		if v[i] == NUMERAL {
			continue
		}
		// else
		v[i] += 1
		for j := 0; j < i; j++ {
			v[j] = 0
		}

		if v.isOk() {
			return true
		}
	}

	return false
}

/*----------------------------------------------------------------------------*/
func (v vector) isOk() bool {

	for i := 0; i < LIMIT-2; i++ {
		if v[i]+v[i+1]+v[i+2] > UPPER {
			return false
		}
	}

	return true
}

/*----------------------------------------------------------------------------*/
func isOk(n int) bool {

	ds := make([]int, 0, LIMIT)

	for n > 0 {
		ds = append(ds, n%10)
		n /= 10
	}

	for i := 0; i < LIMIT-2; i++ {
		if ds[i]+ds[i+1]+ds[i+2] > UPPER {
			return false
		}
	}

	return true
}
