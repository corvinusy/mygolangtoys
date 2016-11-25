package main

import (
	"bufio"
	"fmt"
	"os"
)

type data struct {
	costs []int
	cash  int
}

func main() {
	var (
		t, costsLen int
		rdSlice     []interface{}
	)
	rd := bufio.NewReader(os.Stdin)
	tStr, _ := rd.ReadString('\n')
	fmt.Sscan(tStr, &t)
	input := make([]data, t)
	for i := range input {
		cashStr, _ := rd.ReadString('\n')
		fmt.Sscan(cashStr, &input[i].cash)
		lenStr, _ := rd.ReadString('\n')
		fmt.Sscan(lenStr, &costsLen)
		costsStr, _ := rd.ReadString('\n')
		input[i].costs = make([]int, costsLen)
		rdSlice = make([]interface{}, costsLen)
		for k := range rdSlice {
			rdSlice[k] = &input[i].costs[k]
		}
		fmt.Sscanln(costsStr, rdSlice...)
	}
	for i := range input {
		fmt.Println(getFlavorIndexes(input[i]))
	}
}

func getFlavorIndexes(d data) (int, int) {
	var (
		i, ii int
		found bool
	)
	for i = range d.costs {
		found, ii = searchValue(d.costs, i+1, d.cash-d.costs[i])
		if !found {
			continue
		}
		return i + 1, ii + 1
	}
	return -1, -1
}

func searchValue(a []int, start, v int) (bool, int) {
	for i := start; i < len(a); i++ {
		if a[i] == v {
			return true, i
		}
	}
	return false, -1
}
