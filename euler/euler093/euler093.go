package main

import (
	"fmt"
	"github.com/cznic/mathutil"
	"math"
	"sort"
)

const (
	MUL = iota
	DIV = iota
	ADD = iota
	SUB = iota
)

const (
	ORD1 = iota
	ORD2 = iota
	ORD3 = iota
)

const FAIL = -1e6
const LIMIT = 10

func main() {

	nums := []int{1, 2, 3, 4}

	resVector := make([]int, 4)
	result := 0

	for nums[3] < LIMIT {
		seria := getMaxSeria(getSeria(nums))
		//		fmt.Println(nums, seria)
		if result < seria {
			result = seria
			copy(resVector, nums)
			fmt.Println(result, resVector)
		}
		advance(nums)
	}

	fmt.Println(result, resVector)

	return

}

/*----------------------------------------------------------------------------*/
func advance(nums []int) {

	for i := 0; i < 3; i++ {
		if nums[i] < nums[i+1]-1 {
			nums[i] += 1
			for j := 0; j < i; j++ {
				nums[j] = j + 1
			}
			return
		}
	}

	nums[len(nums)-1] += 1
	for j := 0; j < len(nums)-1; j++ {
		nums[j] = j + 1
	}
	return

}

/*----------------------------------------------------------------------------*/
func getMaxSeria(m map[float64]bool) int {

	i := 1
	for m[float64(i)] {
		i++
	}

	return i - 1
}

/*----------------------------------------------------------------------------*/
func getSeria(vec []int) map[float64]bool {

	vv := make([]int, 4)

	copy(vv, vec)

	v := sort.IntSlice(vv)

	vf := make([]float64, 4)

	ord := sort.IntSlice([]int{ORD1, ORD2, ORD3})

	m := make(map[float64]bool)

	ops := getComb([]int{MUL, DIV, ADD, SUB}, 3)

	mathutil.PermutationFirst(v)
	for ok1 := true; ok1; ok1 = mathutil.PermutationNext(v) {
		for _, op := range ops {
			mathutil.PermutationFirst(ord)
			for ok2 := true; ok2; ok2 = mathutil.PermutationNext(ord) {
				for i := range v {
					vf[i] = float64(v[i])
				}
				k := makeCalc(ord, op, vf)
				if k > 0 && math.Floor(k) == k {
					if !m[k] {
						m[k] = true
					}
				}
			}
		}
	}

	return m
}

/*----------------------------------------------------------------------------*/
func makeCalc(ord, op []int, v []float64) float64 {

	result := 0.0

	for i := range ord {
		if ord[i] == ORD1 {
			if FAIL == makeOP(op[i], v, i) {
				return FAIL
			}
		}
	}

	for i := range ord {
		if ord[i] == ORD2 {
			if FAIL == makeOP(op[i], v, i) {
				return FAIL
			}
		}
	}

	for i := range ord {
		if ord[i] == ORD3 {
			result = makeOP(op[i], v, i)
			if result == FAIL {
				return FAIL
			}
		}
	}

	return result
}

/*----------------------------------------------------------------------------*/
func makeOP(opi int, v []float64, i int) float64 { // bad function

	switch opi {
	case MUL:
		{
			v[i] *= v[i+1]
			v[i+1] = v[i]
		}
	case ADD:
		{
			v[i] += v[i+1]
			v[i+1] = v[i]
		}
	case SUB:
		{
			v[i] -= v[i+1]
			v[i+1] = v[i]
		}
	default: //case DIV:
		if v[i+1] == 0 {
			return FAIL
		} else {
			v[i] /= v[i+1]
			v[i+1] = v[i]
		}
	}

	return v[i]

}

/*----------------------------------------------------------------------------*/
func getComb(ops []int, n int) [][]int {

	if n == 0 {
		return nil
	}

	result := make([][]int, 0)

	// find tail-combinations
	for i := range ops {

		tail := getComb(ops, n-1)

		if tail == nil {
			tmp := make([]int, 1)
			tmp[0] = ops[i]
			result = append(result, tmp)
			continue
		}

		for j := range tail {
			tmp := make([]int, 1)
			tmp[0] = ops[i]
			tmp = append(tmp, tail[j]...)
			result = append(result, tmp)
		}
	}

	return result

}
