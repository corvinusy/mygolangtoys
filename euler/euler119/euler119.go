package main

import (
	"fmt"
	"sort"
)

type Uint64Slice []uint64

func main() {

	count := 0
	result := make([]uint64, 0)
	for i := uint64(2); count <= 32; i++ {
		p := i * i
		for j := uint64(2); j < 10; j++ {
			if digitSum(p) == i {
				result = append(result, p)
				count++
			}
			p = p * i
		}
	}
	sort.Sort(Uint64Slice(result))
	fmt.Println(result[29])

}

/*-----------------------------------------------------------------------------*/
func digitSum(p uint64) uint64 {

	res := uint64(0)
	for p > 0 {
		res += p % 10
		p /= 10
	}
	return res
}

/*-----------------------------------------------------------------------------*/
func (slice Uint64Slice) Len() int {
	return len(slice)
}

/*-----------------------------------------------------------------------------*/
func (slice Uint64Slice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

/*-----------------------------------------------------------------------------*/
func (slice Uint64Slice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

/*-----------------------------------------------------------------------------*/
