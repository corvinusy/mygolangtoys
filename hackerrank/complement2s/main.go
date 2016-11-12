package main

import (
	"bufio"
	"fmt"
	"os"
)

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
