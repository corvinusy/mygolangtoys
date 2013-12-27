package main

import (
    "fmt"
)

const LIMIT = 30

func main() {

	a := make([]int, LIMIT)

	count := 1 // All OK
	total := 1

	for advance(a) {
		total += 1
		if isOk(a) {
			count += 1
		}
	}
	
	fmt.Println(count, total)
}
/*----------------------------------------------------------------------------*/
func advance(a []int) bool {

	has2 := -1

	for i := range a {
		if a[i] == 2 {
			has2 = i
			break
		}
	}

	if has2 == -1 {
		a[0] += 1
		return true
	}

	for i := 0; i < has2; i++ {
		if a[i] == 0 {
			a[i] += 1
			for j := 0; j < i; j++ {
				a[j] = 0
			}
			return true
		}
	}

	for i := has2; i < LIMIT; i++ {
		if a[i] != 2 {
			a[i] += 1
			for j := 0; j < i; j++ {
				a[j] = 0
			}
			return true
		}

	}

	return false
}
/*----------------------------------------------------------------------------*/
func isOk(a []int) bool {

	for i := 0; i < LIMIT-2; i++ {
		if a[i] == 1 && a[i+1] == 1 && a[i+2] == 1 {
			return false
		}
	}

	return true
	
}
