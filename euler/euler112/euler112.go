package main

import (
    "fmt"
)

const LIMIT = 1e7

func main() {

	froggy := 0
	i := int64(1)

	reached05 := false
	reached09 := false

	for ; i < LIMIT; i++ {
		if is_froggy(i) {
			froggy++
		}
		if (float64(froggy)/float64(i) >= 0.5) && !reached05 {
			reached05 = true
			fmt.Println("reached 0.5", i)
		}

		if (float64(froggy)/float64(i) >= 0.9) && !reached09 {
			reached09 = true
			fmt.Println("reached 0.9", i)
		}

		if float64(froggy)/float64(i) >= 0.99 {
			fmt.Println("reached 0.99", i)
			break
		}

	}
}
/*-----------------------------------------------------------------------------*/
func is_froggy(n int64) bool {

	d := make([]int64,0)

	if n < 100 {
		return false
	}
	
	if n > 1000000000 {
		return true
	}

	for n > 0 {
		d = append(d, n % 10)
		n /= 10
	}

	is_upper := true
	is_lower := true

	for i := 1; i < len(d); i++ {
		if d[i] > d[i-1] {
			is_lower = false
		}
		if d[i] < d[i-1] {
			is_upper = false
		}

		if !is_lower && !is_upper {
			return true
		}
	}
	return false
}
