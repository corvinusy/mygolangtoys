package main

import (
    "fmt"
    "time"
)

/*
	//50 = 6
	//100 = 12
	//1e3 = 114
	//1e4 = 1161
	//1e5 = 11624
*/

type Triple struct {
    a   int
    b   int
    c   int
}

func main() {

    var LIMIT int = 1.5e6

    pslice := make([]Triple, 0)

    var start, next Triple

    start.a = 3
    start.b = 4
    start.c = 5

    t1 := time.Now()

    pslice = append(pslice, start)

    // create primitive Pyth-Triplets

    for i := 0; ; i++ {
        next = nextU(pslice[i])
        if next.perimeter() <= LIMIT {
            pslice = append(pslice, next)
        }

        next = nextA(pslice[i])
        if next.perimeter() <= LIMIT {
            pslice = append(pslice, next)
        }

        next = nextD(pslice[i])
        if next.perimeter() <= LIMIT {
            pslice = append(pslice, next)
        }

        if len(pslice) == i+1 {
            break
        }
    }

    // append non-primitive Pyth-Triplets

    prim_len := len(pslice)

    for i := 0; i < prim_len; i++ {
        for j := 2; ; j++ {
            next = pslice[i].multiply(j)
            if next.perimeter() <= LIMIT {
                pslice = append(pslice, next)
            } else {
                break
            }
        }
    }

    perimap := make(map[int]int, 0)

    for i := 0; i < len(pslice); i++ {
        perimap[pslice[i].perimeter()]++
    }

    count := 0
    for _, n := range perimap {
        if n == 1 {
            count++
        }
    }

    t2 := time.Since(t1)

    fmt.Println(LIMIT, count, t2)

    return
}

/*-----------------------------------------------------------------------------*/
/*
Pythagorean primitive matrices
  1  2  2
 -2 -1 -2  U
  2  2  3

  1  2  2
  2  1  2  A
  2  2  3

 -1 -2 -2
  2  1  2  D
  2  2  3
*/

func nextU(t Triple) Triple {
    var next Triple

    next.a = t.a - 2*t.b + 2*t.c
    next.b = 2*t.a - t.b + 2*t.c
    next.c = 2*t.a - 2*t.b + 3*t.c

    return next
}

/*-----------------------------------------------------------------------------*/
func nextA(t Triple) Triple {
    var next Triple

    next.a = t.a + 2*t.b + 2*t.c
    next.b = 2*t.a + t.b + 2*t.c
    next.c = 2*t.a + 2*t.b + 3*t.c

    return next
}

/*-----------------------------------------------------------------------------*/
func nextD(t Triple) Triple {
    var next Triple

    next.a = -1*t.a + 2*t.b + 2*t.c
    next.b = -2*t.a + t.b + 2*t.c
    next.c = -2*t.a + 2*t.b + 3*t.c

    return next
}

/*-----------------------------------------------------------------------------*/
func (t *Triple) perimeter() int {
    return t.a + t.b + t.c
}

/*-----------------------------------------------------------------------------*/
func (t *Triple) multiply(n int) Triple {
    var next Triple

    next.a = t.a * n
    next.b = t.b * n
    next.c = t.c * n

    return next
}
