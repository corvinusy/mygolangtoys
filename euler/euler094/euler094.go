package main

import (
    "fmt"
)

/*  Herons S = sqrt(p(p-a)(p-b)(p-c)
 *  S*S = p(p-a)(p-b)(p-c)
 * if a==b and (b == c+1 or b == c-1) --> SS <= (2/3 * p) ** 3 * p; SS <= 8/27 * p**4;  SS <= 8/27 P**4 / 16; SS <= 1/50 * P**4
 */

const LIMIT = 1e9

func main() {

	var i, j, s1, sum uint64

	sum = 0

	i = 3; j = 1

	for  {

		if 3*i - 1 > LIMIT {
			break
		}

		s1 = (3*i - 1) * (i+1)

		switch {
		case s1 > j*j : {j++; continue}

		case s1 < j*j : {i++; continue}

		case s1 == j*j : {sum += 3*i - 1; i++ }

		}
	}

	i = 3; j = 1

	for  {

		if 3*i + 1 > LIMIT {
			break
		}

		s1 = (3*i + 1) * (i - 1)

		switch {
		case s1 > j*j : {j++; continue}

		case s1 < j*j : {i++; continue}

		case s1 == j*j : {sum += 3*i + 1; i++ }

		}

	}


	fmt.Println(sum)
	
}
