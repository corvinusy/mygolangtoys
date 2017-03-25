package main

import "fmt"

type square struct {
	r int
	c int
}

func main() {
	var n, k, r, c int
	fmt.Scan(&n, &k) //nolint
	var q square
	fmt.Scan(&r, &c) //nolint
	r--
	c--
	q = square{r, c}
	o := make(map[square]bool)
	for i := 0; i < k; i++ {
		fmt.Scan(&r, &c) //nolint
		r--
		c--
		if isStriked(q, r, c) {
			o[square{r, c}] = true
		}
	}
	fmt.Println(countStrikes(q, o, n))
}

func isStriked(q square, r, c int) bool {
	dc := c - q.c
	dr := r - q.r
	return dc == 0 || dr == 0 || abs(dc) == abs(dr)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func countStrikes(q square, o map[square]bool, n int) int {
	count := 0
	count += checkHorizontal(q, o, n)
	count += checkVertical(q, o, n)
	count += checkLLDiag(q, o, n)
	count += checkULDiag(q, o, n)
	return count
}

func checkVertical(q square, o map[square]bool, n int) int {
	count := 0
	p := q
	for p.r = q.r + 1; !isOutOfBoard(p, n); p.r++ {
		if o[p] {
			break
		}
		count++
	}
	for p.r = q.r - 1; !isOutOfBoard(p, n); p.r-- {
		if o[p] {
			break
		}
		count++
	}
	return count
}

func checkHorizontal(q square, o map[square]bool, n int) int {
	count := 0
	p := q
	for p.c = q.c + 1; !isOutOfBoard(p, n); p.c++ {
		if o[p] {
			break
		}
		count++
	}
	for p.c = q.c - 1; !isOutOfBoard(p, n); p.c-- {
		if o[p] {
			break
		}
		count++
	}
	return count
}

func checkLLDiag(q square, o map[square]bool, n int) int {
	count := 0
	var p square
	for p.r, p.c = q.r+1, q.c+1; !isOutOfBoard(p, n); p.r, p.c = p.r+1, p.c+1 {
		if o[p] {
			break
		}
		count++
	}
	for p.r, p.c = q.r-1, q.c-1; !isOutOfBoard(p, n); p.r, p.c = p.r-1, p.c-1 {
		if o[p] {
			break
		}
		count++
	}
	return count
}

func checkULDiag(q square, o map[square]bool, n int) int {
	count := 0
	var p square
	for p.r, p.c = q.r+1, q.c-1; !isOutOfBoard(p, n); p.r, p.c = p.r+1, p.c-1 {
		if o[p] {
			break
		}
		count++
	}
	for p.r, p.c = q.r-1, q.c+1; !isOutOfBoard(p, n); p.r, p.c = p.r-1, p.c+1 {
		if o[p] {
			break
		}
		count++
	}
	return count
}

func isOutOfBoard(p square, n int) bool {
	return p.r < 0 || p.c < 0 || p.r >= n || p.c >= n
}
