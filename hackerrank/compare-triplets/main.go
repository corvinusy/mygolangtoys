package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		aCount, bCount int
	)
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	rd := bufio.NewReader(os.Stdin)
	a := readIntSlice(rd)
	b := readIntSlice(rd)
	if len(a) != len(b) {
		panic("different slice lengths")
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			aCount++
		} else if b[i] > a[i] {
			bCount++
		}
	}
	fmt.Println(aCount, bCount)
}

func readIntSlice(rd *bufio.Reader) []int {
	var (
		x int
		s string
		a []int
	)
	s, _ = rd.ReadString('\n')
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &x)
		a = append(a, x)
	}
	return a[:3]
}
