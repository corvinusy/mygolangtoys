package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	rd := bufio.NewReader(os.Stdin)
	bs, _ := rd.ReadBytes('\n')
	fmt.Println(totalValleys(bs))
}

func totalValleys(bs []byte) int {
	var level, total int
	for i := range bs {
		if bs[i] == 'U' {
			level++
			continue
		}
		if bs[i] == 'D' {
			if level == 0 {
				total++
			}
			level--
		}
	}
	return total
}
