package main

import (
	"fmt"
	"sort"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		err  error
		n, r int
	)
	_, err = fmt.Scan(&n, &r)
	panicIfErr(err)
	ms := make([]int, r)
	for i := range ms {
		_, err = fmt.Scan(&ms[i])
		panicIfErr(err)
	}
	sort.Ints(ms)
	fmt.Println(remains(n, ms))
}

func remains(n int, ms []int) int {
	var (
		ok bool
	)
	p := map[int]map[int]bool{}
	for i := 0; i <= n; i++ {
		p[n] = make(map[int]bool)
	}
	m := map[int]int{}
	m[0] = 0
	p[ms[0]][1] = true
	for i := 1; i <= n; i++ {
		if _, ok = m[i]; !ok {
			getRemains(i, ms, m, p)
		}
	}
	fmt.Printf("%+v\n, %+v\n", m, p)
	return m[n]
}

func getRemains(n int, ms []int, m map[int]int, p map[int]map[int]bool) {
	var (
		ok bool
		k  int
	)
	m[n] = 0
	p[n] = make(map[int]bool)
	for _, v := range ms {
		if n-v < 0 {
			continue
		}
		if n-v == 0 && !p[n][k] {
			m[n]++
			k++
			p[n][k] = true
		}
		if _, ok = m[n-v]; !ok {
			getRemains(n-v, ms, m, p)
		}
		m[n] += m[n-v]
		k = k << 8
	}

}
