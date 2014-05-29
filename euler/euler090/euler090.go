package main

import (
	"fmt"
)

func main() {

	count := 0

	for nA := 0; nA < 1<<10; nA++ {

		if bitLen(nA) != 6 {
			continue
		}

		a := getArrayFromMask(nA)

		for nB := nA + 1; nB < 1<<10; nB++ {

			if bitLen(nB) != 6 {
				continue
			}

			b := getArrayFromMask(nB)

			if isOk(a, b) {
				count++
			}
		}

	}

	fmt.Println(count)
}

/*----------------------------------------------------------------------------*/
func isOk(a, b [6]int) bool {

	var ok01, ok04, ok09, ok16, ok25, ok36, ok49, ok64, ok81 bool

	for i := range a {
		for j := range b {
			switch a[i]*10 + b[j] {
			case 1:
				ok01 = true
			case 4:
				ok04 = true
			case 6, 9:
				ok09 = true
			case 16, 19:
				ok16 = true
			case 25:
				ok25 = true
			case 36, 39:
				ok36 = true
			case 46, 49:
				ok49 = true
			case 64, 94:
				ok64 = true
			case 81:
				ok81 = true
			}

			switch a[i] + b[j]*10 {
			case 1:
				ok01 = true
			case 4:
				ok04 = true
			case 6, 9:
				ok09 = true
			case 16, 19:
				ok16 = true
			case 25:
				ok25 = true
			case 36, 39:
				ok36 = true
			case 46, 49:
				ok49 = true
			case 64, 94:
				ok64 = true
			case 81:
				ok81 = true
			}

			if ok01 && ok04 && ok09 && ok16 && ok25 && ok36 && ok49 && ok64 && ok81 {
				return true
			}
		}
	}

	return false
}

/*----------------------------------------------------------------------------*/
func bitLen(n int) int {

	result := 0

	for n > 0 {
		if n&1 != 0 {
			result++
		}
		n = n >> 1
	}

	return result
}

/*----------------------------------------------------------------------------*/
func getArrayFromMask(n int) [6]int {

	var a [6]int

	index := 0

	for i := 0; i < 10; i++ {
		if n&1 != 0 {
			a[index] = i
			index++
		}
		n = n >> 1
	}

	return a
}
