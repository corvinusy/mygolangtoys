package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const n = 10

func main() {
	rd := bufio.NewReader(os.Stdin)
	ph := make([]int, n)
	tph := make([]interface{}, n)
	for i := range ph {
		tph[i] = &ph[i]
	}
	s, _ := rd.ReadString('\n')
	fmt.Sscanln(s, tph...)

	hi := make([]int, n)
	thi := make([]interface{}, n)
	for i := range hi {
		thi[i] = &hi[i]
	}
	s, _ = rd.ReadString('\n')
	fmt.Sscanln(s, thi...)

	slope := getSlope(ph, hi)
	intercept := getIntercept(ph, hi)

	fmt.Println("pearson =", getPearson(ph, hi))
	fmt.Println("slope =", slope)
	fmt.Println("intercept =", intercept)
	hiScore := float64(10)
	fmt.Println("probable score = ", intercept+slope*hiScore)

}

func getPearson(x, y []int) float64 {
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

	var lowSumX, lowSumY float64
	for i := range x {
		lowSumX += (float64(x[i]) - xav) * (float64(x[i]) - xav)
		lowSumY += (float64(y[i]) - yav) * (float64(y[i]) - yav)
	}

	lowSum := math.Sqrt(lowSumX * lowSumY)

	return upSum / lowSum
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
