package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	positive part
N(n) = N(2*n) для n = 2m
N(n+1) = N(n) + 1 для т = 2т+1
N(1) = N(1)

S(n) = S(2m) + S(2m+1)

*/
//negCache := make(map[uint32]uint32, 0)
//posCache := make(map[uint32]uint32, 0)

func main() {
	rd := bufio.NewReader(os.Stdin)
	tStr, _ := rd.ReadString('\n')
	var t int
	fmt.Sscan(tStr, &t)
	dat := make([][2]int32, t)
	for i := range dat {
		qStr, _ := rd.ReadString('\n')
		fmt.Sscan(qStr, &dat[i][0], &dat[i][1])
	}
	for i := range dat {
		fmt.Println(getOnesNumber(dat[i][0], dat[i][1]))
	}
}

func getOnesNumber(start, end int32) uint64 {
	var count uint64
	for i := start; i <= end; i++ {
		s := uint32(i)
		for s != 0 {
			count += uint64(s % 2)
			s >>= 1
		}
	}
	return count
}
