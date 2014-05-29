package main

import (
	"fmt"
)

/*
 S(B) != S(C)
 if {B} > {A} then S(B) > S(A)
*/

const LIMIT = 12

type TSlice []int

func main() {

	var a TSlice

	a = make([]int, LIMIT)

	a[LIMIT-1] = 1

	allCount := 0
	resCount := 0

	for !a.isEmpty() {
		a.advance()
		if a.isQualified() {
			allCount += 1
			if a.isCounted() {
				resCount += 1
			}
		}
	}

	fmt.Println("allCount =", allCount, "resCount =", resCount)
}

/*----------------------------------------------------------------------------*/
func (a TSlice) advance() {

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] < 2 {
			a[i] += 1
			for j := i + 1; j < len(a); j++ {
				a[j] = 0
			}
			return
		}
	}

	// if not advanced than set allballs

	for j := range a {
		a[j] = 0
	}

	return
}

/*----------------------------------------------------------------------------*/
func (a TSlice) isEmpty() bool {

	for i := range a {
		if a[i] != 0 {
			return false
		}
	}

	return true
}

/*----------------------------------------------------------------------------*/
func (a TSlice) isQualified() bool {

	has1 := false

	for i := range a {

		if a[i] == 1 {
			has1 = true
		}

		if a[i] == 2 {
			return has1 // allowed "2 after 1" only
		}
	}

	return false
}

/*----------------------------------------------------------------------------*/
func (a TSlice) isCounted() bool {

	sumOf1 := 0
	sumOf2 := 0
	check_trigger := false

	for i := range a {
		if a[i] == 1 {
			sumOf1 += 1
		}
		if a[i] == 2 {
			sumOf2 += 1
			if sumOf2 > sumOf1 {
				check_trigger = true
			}
		}
	}

	return check_trigger && (sumOf1 == sumOf2)
}
