package main

import (
    "fmt"
)

func main() {

	maskPete := []int{1,1,1,1,1,1,1,1,1}

	maskColin := []int{1,1,1,1,1,1}

	peteMap := make([]int, 40)
	colinMap := make([]int, 40)

	peteMap[9] = 1
	colinMap[6] = 1

	for advancePete(maskPete) {
		peteMap[getCombo(maskPete)] += 1
	}

	for advanceColin(maskColin) {
		colinMap[getCombo(maskColin)] += 1
	}

	sumPete := 0
	sumColin := 0

	for i := range peteMap {
		sumPete += peteMap[i]
	}

	for i := range colinMap {
		sumColin += colinMap[i]
	}

	win := 0

	for i := 0; i < len(peteMap); i++ {
		sum := 0
		for j := 6; j < i; j++ {
			sum += colinMap[j]
		}

		win += peteMap[i] * sum

	}

	fmt.Println(win, sumPete*sumColin)

	fmt.Println(float64(win)/float64(sumPete*sumColin))

}
/*----------------------------------------------------------------------------*/
func advancePete(mask []int) bool {

	for i := range mask {
		if mask[i] != 4 {
			mask[i] += 1
			for j := range mask[:i] {
				mask[j] = 1
			}
			return true
		}
	}

	return false
}
/*----------------------------------------------------------------------------*/
func getCombo(mask []int) int {

	sum := 0
	
	for _, m := range mask {
		sum += m
	}

	return sum
}
/*----------------------------------------------------------------------------*/
func advanceColin(mask []int) bool {

	for i := range mask {
		if mask[i] != 6 {
			mask[i] += 1
			for j := range mask[:i] {
				mask[j] = 1
			}
			return true
		}
	}

	return false

}
