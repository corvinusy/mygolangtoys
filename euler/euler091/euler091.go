package main

import (
	"fmt"
)

// number points for point B
// number points for point C

type Point struct {
	x, y int
}

const LIMIT = 50 // actually grid has indexes +1: 0, 1, 2 for 2

func main() {

	a := Point{0, 0}

	count := 0

	for xb := a.x; xb <= LIMIT; xb++ {
		for yb := a.y + 1; yb <= LIMIT; yb++ {
			b := Point{xb, yb}
			for xc := a.x + 1; xc <= LIMIT; xc++ {
				for yc := a.y; yc <= b.y; yc++ {
					c := Point{xc, yc}
					if c.x == b.x && c.y == b.y {
						continue
					}
					q_ab := b.x*b.x + b.y*b.y
					q_ac := c.x*c.x + c.y*c.y
					q_bc := (c.x-b.x)*(c.x-b.x) + (b.y-c.y)*(b.y-c.y)
					if q_ab == q_ac+q_bc {
						count++
						fmt.Println(b, c)
					}
					if q_ac == q_ab+q_bc {
						count++
						fmt.Println(b, c)
					}
					if q_bc == q_ab+q_ac {
						count++
						fmt.Println(b, c)
					}
				}
			}
		}

		/*			switch {
					case a.x == b.x && b.y % 2 != 0: count_c1 += 2*LIMIT
					case a.x == b.x && b.y % 2 == 0: count_c2 += 2*LIMIT + 1
					case b.x == b.y && b.y <= LIMIT/2: count_c3 += 2
					default: count_df += 1
					}
		*/
	}

	fmt.Println(count)

}
