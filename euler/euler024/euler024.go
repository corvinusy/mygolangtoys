package main

import (
    "fmt"
)

func main() {

	const LIMIT = 1e6
	
	subj := []byte  {0,1,2,3,4,5,6,7,8,9}

	fmt.Println(subj)
	for i:=0; i < LIMIT-1; i++ {
		next_permutation(0, len(subj) - 1, subj);
	}
	fmt.Println(subj)
}
/*-----------------------------------------------------------------------------*/
func next_permutation(start int, end int, s []byte) {

	if start >= end {
		return
	}

	var (
		i, j, k int
		tmp byte
	)

	i = end
	
	for {
		j = i
		i--

		if (s[i] < s[j]) {
			for k = end; s[i] > s[k]; k-- {
				;
			}
			//swap(s[i], s[k])
			tmp = s[i]
			s[i] = s[k]
			s[k] = tmp
		
			reverse(j, end, s)
			return 
		}

		if (i == start) {
			reverse(start, end, s)
			return
		}

	}
}
/*-----------------------------------------------------------------------------*/
func reverse(start int, end int, s []byte) {

	if start >= end {
		return
	}

	var tmp byte
	
	for {
		if (start >= end) {
			return
		} else {
			tmp = s[start]
			s[start] = s[end]
			s[end] = tmp
			end--
			start++
		}
	}
}
