package main

import (
    "fmt"
	"math"
	"sort"
)

//const LIMIT = 5e7

const LIMIT = 1e5

type Irad struct {
	n int
	rad int
}

type ByRadInt []Irad

func main() {
	
	primes := create_primes_atkin(LIMIT*2)

	es := make([]Irad, 0)
	es = append(es, Irad{0,0}, Irad{1,1})

	for i := 2; i <= LIMIT; i++ {
		es = append(es, Irad{i, rad(i, primes)})
	}

	sort.Sort(ByRadInt(es))
	fmt.Println(es[1e4])
	
	
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit int) []int  {

    var i, x, y, n int

    sqr_lim := int(math.Sqrt(float64(limit)))
    sieve_nums := make([]bool, limit+1)
	primes := make([]int, 0)


    for i = 5; i <= limit ; i++ {
        sieve_nums[i] = false
    }

    sieve_nums[2] = true
    sieve_nums[3] = true
    
    for x = 1; x <= sqr_lim; x++ {
        for y = 1; y <= sqr_lim; y++ {

            n = 4 * x * x + y * y
            if (n <= limit) && ( (n % 12 == 1) || (n % 12 == 5) ) {
                sieve_nums[n] = !sieve_nums[n]
            }

            n = n - x * x
            if (n <= limit) && (n % 12 == 7) {
                sieve_nums[n] = !sieve_nums[n]
            }

            n = n - 2 * y * y
            if (x > y) && (n <= limit) && (n % 12 == 11) {
                sieve_nums[n] = !sieve_nums[n]
            }
        }
    }

    for i = 5; i <= sqr_lim; i++ {
        if sieve_nums[i] {
            n = i * i
            for j := n; j <= limit; j += n {
                sieve_nums[j] = false
            }
        }
    }
    
    for i = 0; i <= limit; i++ {
        if sieve_nums[i] {
            primes = append(primes, i)
        } 
    }

    return primes
}
/*-----------------------------------------------------------------------------*/
func rad(n int, primes []int) int {

	rad := 1
	for i := 0; primes[i] <= n; i++ {
		if n % primes[i] == 0 {
			rad *= primes[i]
		}
		for n % primes[i] == 0 {
			n /= primes[i]
		}
	}
	return rad
}
/*-----------------------------------------------------------------------------*/
func (e ByRadInt) Len() int {
	return len(e)
}
/*-----------------------------------------------------------------------------*/
func (e ByRadInt) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
/*-----------------------------------------------------------------------------*/
func (e ByRadInt) Less(i, j int) bool {
	switch {
	case e[i].rad < e[j].rad : return true
	case e[i].rad == e[j].rad: return e[i].n < e[j].n
	default: return false
	}
}
/*-----------------------------------------------------------------------------*/
