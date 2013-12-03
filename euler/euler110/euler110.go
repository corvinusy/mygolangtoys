package main

import (
	"fmt"
	"math"
)


/* If n = (p1^a1)(p2^a2)...(pt^at), 
 * 2-partitions = ((2 a1 + 1)(2 a2 + 1) ... (2 at + 1) + 1)/2. 
 * We want ((2 a1 + 1)(2 a2 + 1) ... (2 at + 1) + 1)/2 > 4e6 
 * (2 a1 + 1)(2 a2 + 1) ... (2 at + 1) > 8e6
 *
 */

const LIMIT = 4e6

var primes []uint

func main() {

	primes = create_primes_atkin(1000)

	vec := make([]uint,0)

	var num, res  uint

	for  {

		vec = growVec(vec)

		if calcUnits(vec) > LIMIT {
			res = getNumFromVec(vec)

			for i := 0; i < 100; i++  {
				if getNumFromVec(vec) <= res {
					fmt.Println(vec, getNumFromVec(vec), calcUnits(vec))
				}
				vec = advanceVec(vec)
				num = getNumFromVec(vec)
				if num < res && calcUnits(vec) > LIMIT {
					res = num
					continue
				}
//				break
			}

			fmt.Println(res)
			break
		}
	}

}
/*----------------------------------------------------------------------------*/
func calcUnits(vec []uint) uint {
	units := uint(1)

	for i, _ := range vec {
		units *= 2*vec[i] + 1
	}

	units += 1; units /= 2

	return units

}
/*----------------------------------------------------------------------------*/
func growVec (vec []uint) []uint {

	vec = append(vec, 1)

	return vec
}
/*----------------------------------------------------------------------------*/
func advanceVec (vec []uint) []uint {

	low := 0
	high := len(vec)-1;

	for i := len(vec)-1; i >= 0; i-- {
		if vec[i] != 0 {
			high = i
			break
		}		
	}

	for i := 0; i < high-1; i++ {
		if vec[i] <= vec[i+1] {
			continue
		}
	    low = i+1
	}
	
	if calcUnits(vec) > LIMIT {
		vec[high]--
	} else {
		vec[low]++
	}


	return vec
}
/*----------------------------------------------------------------------------*/
func getFactors (num uint) []uint {

//	sqrnum := uint(math.Sqrt(float64(num)))
	
	vector := make([]uint, 0)

	for i := 0; primes[i] <= num; i++ {

		vector = append(vector, 0)

		for num % primes[i] == 0 {
			vector[ len(vector)-1 ]++
			num /= primes[i]
		}
	}

	return vector
}
/*----------------------------------------------------------------------------*/
func getNumFromVec(vec []uint) uint {
	
	num := uint(1)

	for i, _ := range vec {
		for j := vec[i]; j > 0; j-- {
			num *= primes[i]
		}
	}
	return num
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit uint) []uint  {

    var i, x, y, n uint

    sqr_lim := uint(math.Sqrt(float64(limit)))
    sieve_nums := make([]bool, limit+1)
	primes := make([]uint, 0)


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
/*----------------------------------------------------------------------------*/
