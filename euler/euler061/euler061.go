package main

import (
    "fmt"
	"strconv"
	"time"
)

func main() {

    const LIMIT = 1e3

    var (
		i, n uint64
		
	)

	t1 := time.Now()

	tr := make([]uint64, 0)
	qu := make([]uint64, 0)
	pe := make([]uint64, 0)
	si := make([]uint64, 0)
	se := make([]uint64, 0)
	oc := make([]uint64, 0)


	for i = 1; i <= LIMIT; i++ {
		n = (i * (i + 1)) >> 1
		if len(strconv.FormatUint(n, 10)) == 4 {
			tr = append(tr, n)
		}

		n = i * i
		if len(strconv.FormatUint(n, 10)) == 4 {
			qu = append(qu, n)
		}


		n = (i * (3 * i - 1)) >> 1
		if len(strconv.FormatUint(n, 10)) == 4 {
			pe = append(pe, n)
		}

		n = i * (2 * i - 1)
		if len(strconv.FormatUint(n, 10)) == 4 {
			si = append(si, n)
		}

		n = (i * (5 * i - 3)) >> 1
		if len(strconv.FormatUint(n, 10)) == 4 {
			se = append(se, n)
		}

		n = i * (3 * i - 2)
		if len(strconv.FormatUint(n, 10)) == 4 {
			oc = append(oc, n)
		}
	}

	fmt.Println("passed")

	for i1 := 0; i1 < len(tr); i1++ {
		for i2 := 0; i2 < len(qu); i2++ {
			for i3 := 0; i3 < len(pe); i3++ {
				for i4 := 0; i4 < len(si); i4++ {
					if !is_semi_cyclic(tr[i1], qu[i2], pe[i3], si[i4]) {
						break
					}
					for i5 := 0; i5 < len(se); i5++ {
						if !is_semi_cyclic2(tr[i1], qu[i2], pe[i3], si[i4], se[i5]) {
							break
						}
						for i6 := 0; i6 < len(oc); i6++ {
							if is_cyclic(tr[i1], qu[i2], pe[i3], si[i4], se[i5], oc[i6]) {
								t2 := time.Since(t1)
								fmt.Println(tr[i1], qu[i2], pe[i3], si[i4], se[i5], oc[i6],"\n",  
									tr[i1] + qu[i2] + pe[i3] + si[i4] + se[i5] + oc[i6], "\ntime = ", t2)
								
							}
						}
					}
				}
			}
		}
	}
}
/*-----------------------------------------------------------------------------*/
func is_cyclic(n1, n2, n3, n4, n5, n6 uint64) bool {

	if (n1 % 100 == n1 /100) || (n2 % 100 == n2 /100) || (n3 % 100 == n3 /100) || 
		(n4 % 100 == n4 /100) || (n5 % 100 == n5 /100) || (n6 % 100 == n6 /100) {
		return false
	}

	ns := make(map[uint64]int, 0)
	ns[n1 % 100]++
	ns[n1 / 100]--
	ns[n2 % 100]++
	ns[n2 / 100]--
	ns[n3 % 100]++
	ns[n3 / 100]--
	ns[n4 % 100]++
	ns[n4 / 100]--
	ns[n5 % 100]++
	ns[n5 / 100]--
	ns[n6 % 100]++
	ns[n6 / 100]--

	for _, d := range ns {
		if d != 0 {
			return false
		}
	}
	return true
}
/*-----------------------------------------------------------------------------*/
func is_semi_cyclic(n1, n2, n3, n4 uint64) bool {

	ns := make(map[uint64]int, 0)
	ns[n1 % 100]++
	ns[n1 / 100]--
	ns[n2 % 100]++
	ns[n2 / 100]--
	ns[n3 % 100]++
	ns[n3 / 100]--
	ns[n4 % 100]++
	ns[n4 / 100]--

	count := 0

	for _, d := range ns {
		if d != 0 {
			count++
		}
	}

	return (count < 6)
}
/*-----------------------------------------------------------------------------*/
func is_semi_cyclic2(n1, n2, n3, n4, n5 uint64) bool {

	ns := make(map[uint64]int, 0)
	ns[n1 % 100]++
	ns[n1 / 100]--
	ns[n2 % 100]++
	ns[n2 / 100]--
	ns[n3 % 100]++
	ns[n3 / 100]--
	ns[n4 % 100]++
	ns[n4 / 100]--
	ns[n5 % 100]++
	ns[n5 / 100]--

	count := 0

	for _, d := range ns {
		if d != 0 {
			count++
		}
	}

	return (count < 6)
}
