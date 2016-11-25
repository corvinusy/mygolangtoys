package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		aCount, bCount int
	)
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	rd := bufio.NewReader(os.Stdin)
	a, _ := readIntSlice(rd, 3)
	b, _ := readIntSlice(rd, 3)
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

func readIntSlice(rd *bufio.Reader, n int) ([]int, error) {
	s, err := rd.ReadString('\n')
	if err != nil {
		return nil, err
	}
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err = fmt.Sscan(s, y...)
	x = x[:n]
	return x, err
}
