package main

import "fmt"

type tst struct {
	total, threshold int
	times            []int
}

type tstTable struct {
	trials int
	tsts   []tst
}

func main() {
	table := tstTable{
		trials: 5,
		tsts: []tst{
			{total: 4, threshold: 3,
				times: []int{-1, 0, 4, 2}},
			{total: 5, threshold: 2,
				times: []int{0, -1, 2, 1, 4}},
			{total: 7, threshold: 6,
				times: []int{2, 0, -1, 1, 1, 1, 1}},
			{total: 3, threshold: 1,
				times: []int{-1, 0, 4}},
			{total: 6, threshold: 4,
				times: []int{0, -1, 1, 4, 5, 6}},
		},
	}

	printInput(table)
}

func printInput(tt tstTable) {
	fmt.Println(tt.trials)
	for i := range tt.tsts {
		fmt.Println(tt.tsts[i].total, tt.tsts[i].threshold)
		printArray(tt.tsts[i].times)
	}
}

func printArray(a []int) {
	for i := range a {
		fmt.Print(a[i])
		if i != len(a)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
