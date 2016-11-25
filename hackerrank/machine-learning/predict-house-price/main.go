package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		f, n int
		scn  []interface{}
		s    string
	)
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &f, &n)

	x := make([][]float64, n)
	p := make([]float64, n)
	for i := range x {
		s, _ = rd.ReadString('\n')
		x[i] = make([]float64, f)
		scn = make([]interface{}, f+1)
		for k := range x[i] {
			scn[k] = &x[i][k]
		}
		scn[f] = &p[i]
		fmt.Sscanln(s, scn...)
	}

	var t int
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &t)

	xx := make([][]float64, t)
	//pp := make([]float64, t)
	for i := range xx {
		s, _ = rd.ReadString('\n')
		xx[i] = make([]float64, f)
		scn = make([]interface{}, f)
		for k := range xx[i] {
			scn[k] = &xx[i][k]
		}
		fmt.Sscanln(s, scn...)
	}
	fmt.Println(x, p, xx)
}

func getSlope(x, y []int) float64 {
	var xav, yav float64
	for i := range x {
		xav += float64(x[i])
	}
	xav /= float64(len(x))
	for i := range y {
		yav += float64(y[i])
	}
	yav /= float64(len(y))

	var upSum float64
	for i := range x {
		upSum += (float64(x[i]) - xav) * (float64(y[i]) - yav)
	}

	var lowSum float64
	for i := range x {
		lowSum += (float64(x[i]) - xav) * (float64(x[i]) - xav)
	}

	return upSum / lowSum

}

func getIntercept(x, y []int) float64 {
	var xav, yav float64
	for i := range x {
		xav += float64(x[i])
	}
	xav /= float64(len(x))
	for i := range y {
		yav += float64(y[i])
	}
	yav /= float64(len(y))

	return yav - getSlope(x, y)*xav

}
