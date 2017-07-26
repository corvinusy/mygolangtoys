package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type strand struct {
	first int
	last  int
	line  string
}

func main() {
	// read genes
	rd := bufio.NewReader(os.Stdin)
	rd.ReadString('\n') // skip useless n
	s, _ := rd.ReadString('\n')
	s = strings.TrimSpace(s)
	genes := strings.Split(s, " ")
	// read gene healths
	healths := make([]int, len(genes))
	s, _ = rd.ReadString('\n')
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	for i := range healths {
		healths[i], _ = strconv.Atoi(ss[i])
	}
	// read strands
	s, _ = rd.ReadString('\n')
	s = strings.TrimSpace(s)
	strandsNum, _ := strconv.Atoi(s)

	strands := make([]strand, strandsNum)
	for i := range strands {
		s, _ = rd.ReadString('\n')
		s = strings.TrimSpace(s)
		fmt.Sscanf(s, "%d %d %s", &strands[i].first, &strands[i].last, &strands[i].line)
	}
	// calc min max healths
	var min, max, h int
	const MaxInt = int(^uint(0) >> 1)
	min = MaxInt

	for i := range strands {
		h = calcGeneHealth(strands[i], healths, genes)
		if h > max {
			max = h
		}
		if h < min {
			min = h
		}
	}
	fmt.Println(min, max)
}

func calcGeneHealth(st strand, healths []int, genes []string) int {
	var (
		total      int
		line, gene string
	)
	for i := st.first; i <= st.last; i++ {
		gene = genes[i]
		line = st.line
		for x := strings.LastIndex(line, gene); x != -1; x = strings.LastIndex(line, gene) {
			total += healths[i]
			line = line[:x+len(gene)-1]
		}
	}
	return total
}
