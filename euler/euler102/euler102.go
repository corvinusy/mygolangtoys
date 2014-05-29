package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Triangle struct {
	a Point
	b Point
	c Point
}

func main() {

	//read file into source

	content, err := ioutil.ReadFile("euler102/triangles.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content), "\n")

	lines := strings.Split(source, "\n") //

	//create array and init it
	ts := make([]Triangle, 0)

	for _, line := range lines {
		var t Triangle
		nums := strings.Split(line, ",")
		t.a.x, _ = strconv.Atoi(nums[0])
		t.a.y, _ = strconv.Atoi(nums[1])
		t.b.x, _ = strconv.Atoi(nums[2])
		t.b.y, _ = strconv.Atoi(nums[3])
		t.c.x, _ = strconv.Atoi(nums[4])
		t.c.y, _ = strconv.Atoi(nums[5])
		ts = append(ts, t)
	}

	count := 0
	//check data
	for _, t := range ts {
		p := Point{0, 0}
		if pt_inside(p, t.a, t.b, t.c) {
			count++
		}
	}

	fmt.Println(count)

	return
}

/*-----------------------------------------------------------------------------*/
func cross_product(a, b Point) Point {
	return Point{0, a.x*b.y - a.y*b.x}
}

/*-----------------------------------------------------------------------------*/
func dot_product(a, b Point) int {
	return a.x*b.x + a.y*b.y
}

/*-----------------------------------------------------------------------------*/
func point_sub(a, b Point) Point {
	return Point{a.x - b.x, a.y - b.y}
}

/*-----------------------------------------------------------------------------*/
func same_side(p1, p2, a, b Point) bool {
	cp1 := cross_product(point_sub(b, a), point_sub(p1, a))
	cp2 := cross_product(point_sub(b, a), point_sub(p2, a))
	if dot_product(cp1, cp2) >= 0 {
		return true
	}
	return false
}

/*-----------------------------------------------------------------------------*/
func pt_inside(p, a, b, c Point) bool {
	if same_side(p, a, b, c) && same_side(p, b, a, c) &&
		same_side(p, c, a, b) {
		return true
	}
	return false
}
