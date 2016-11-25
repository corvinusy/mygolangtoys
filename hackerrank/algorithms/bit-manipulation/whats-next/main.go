package main

import "fmt"

func main() {
	var trials, size int
	fmt.Scan(&trials)
	data := make([][]int, trials)
	results := make([][]int, trials)
	for t := 0; t < trials; t++ {
		fmt.Scan(&size)
		data[t] = make([]int, size)
		for i := range data[t] {
			fmt.Scan(&data[t][i])
		}
		results[t] = reduce(data[t])
	}
	for t := 0; t < trials; t++ {
		fmt.Println(len(results[t]))
		printlnIntArray(results[t])
	}
}

func reduce(d []int) []int {
	// convert to string
	b := unpack(d)
	// find smallest next
	bb := nextS(b)
	// packBack
	r := pack(bb)
	return r
}

func unpack(d []int) []byte {
	var b []byte
	for i := range d {
		for k := 0; k < d[i]; k++ {
			if i%2 == 0 {
				b = append(b, 1)
			} else {
				b = append(b, 0)
			}
		}
	}
	return b
}

func nextS(b []byte) []byte {
	i := len(b) - 1
	for i > 0 && b[i] != 1 {
		i--
	}
	for i > 0 && b[i] != 0 {
		i--
	}
	if i == 0 {
		b = append([]byte{1}, b...)
		b[1] = 0
		return b
	}
	b[i] = 1
	if b[i+1] == 0 {
		b[i+1] = 1
	} else {
		b[i+1] = 0
	}
	return b
}

func pack(b []byte) []int {
	var (
		r   []int
		i   int
		cur byte
		cnt int
	)

	for i < len(b) {
		cur = b[i]
		cnt = 0
		for i < len(b) && cur == b[i] {
			cnt++
			i++
		}
		r = append(r, cnt)
	}
	return r
}

func printlnIntArray(b []int) {
	for i := range b {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()
}
