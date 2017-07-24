package main

import "fmt"

func main() {
	var (
		t int
		s []byte
	)

	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&s)
		fmt.Println(string(permute(s)))
	}
}

func permute(p []byte) []byte {
	i := len(p) - 2
	if i < 0 {
		return []byte("no answer")
	}
	//find i such p[i] < p[i+1]
	for p[i] >= p[i+1] {
		i--
		if i < 0 {
			return []byte("no answer")
		}
	}
	// find minimum(p[i+1:])
	j := len(p) - 1
	for p[j] <= p[i] {
		j--
	}
	// permute
	p[i], p[j] = p[j], p[i]

	// process tail
	i, j = i+1, len(p)-1
	for i < j {
		p[i], p[j] = p[j], p[i]
		i, j = i+1, j-1
	}
	return p
}
