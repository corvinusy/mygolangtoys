package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n int
	)
	// read input
	fmt.Scan(&n)
	a := make([]uint, n)
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	var x uint64
	for i := 0; scanner.Scan(); i++ {
		x, _ = strconv.ParseUint(scanner.Text(), 10, 0)
		a[i] = uint(x)
	}

	// find max differ bit value
	maxBitValue, found := findMaxDifferBitValue(a)

	for !found && (maxBitValue != 0) {
		xorArrayWithMaxBitValue(a, maxBitValue)
		maxBitValue, found = findMaxDifferBitValue(a)
	}

	if maxBitValue == 0 {
		// trivial case, all numbers are equal
		fmt.Println(0)
		return
	}

	// split array by maxBit value
	low := make([]uint, 0, n)
	high := make([]uint, 0, n)
	for i := range a {
		if a[i]|maxBitValue == a[i] {
			high = append(high, a[i])
		} else {
			low = append(low, a[i])
		}
	}

	minScore := low[0] ^ high[0]

	for i := range high {
		for j := range low {
			if minScore > high[i]^low[j] {
				minScore = high[i] ^ low[j]
			}
		}
	}
	fmt.Println(minScore)
}

//------------------------------------------------------------------------------
func findMaxDifferBitValue(a []uint) (uint, bool) {
	// find max(a) and min(a)
	max, min := a[0], a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	// find max bit number of max(a)
	var bitNum uint
	tmp := max
	for tmp > 0 {
		tmp >>= 1
		bitNum++
	}
	// get maxBitValue
	maxBitValue := uint(1 << (bitNum - 1))
	// check if array can be split by maxBitValue
	if (min | maxBitValue) == min {
		return maxBitValue, false //no. Min(a) always has such bit
	}

	return maxBitValue, true // yes
}

//------------------------------------------------------------------------------
func xorArrayWithMaxBitValue(a []uint, maxBitValue uint) {
	for i := range a {
		a[i] ^= maxBitValue
	}
}
