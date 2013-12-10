package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {

	rslice := make([][]int, 0)

	content, err := ioutil.ReadFile("euler105/sets.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content), "\n")

	lines := strings.Split(source,"\n")
	
	for i := range lines {
		rslice = append(rslice, make([]int, 0))
		numbers := strings.Split(lines[i], ",")
		for j := range numbers  {
			num, err := strconv.Atoi(numbers[j])
			if err != nil { panic(err) }
			rslice[i] = append(rslice[i], num)
		}
	}

  // read ok

	sum := 0
	
  // sort and check sequence
	for i := range rslice {
		sort.Sort(sort.IntSlice(rslice[i]))
		if isSatisfies(rslice[i]) {
  // fmt.Println("passed", rslice[i], sliceSum(rslice[i]))
			sum += sliceSum(rslice[i])
		} 
	}

	fmt.Println("sum=", sum)

}
/*----------------------------------------------------------------------------*/
func isSatisfies(a []int) bool {

	// verify rule 1

	// enough to compare "low half of slice" and "high part minus 1 element"

	mid := len(a)/2
	if sliceSum(a[0:mid+len(a)%2]) <= sliceSum(a[mid+1:]) {
		return false
	}

	//verify rule 2

	// build subset-sums map

	seen := make(map[int]uint)

	var (
		lenA uint = uint(len(a))
		bf uint
	)
	
	for bf = 1; bf < 1 << lenA; bf++ {
		// count sum of bit elements
		sum := 0
		for q, j := bf, 0; q != 0; q, j = q >> 1, j + 1 {
			if q & 1 != 0 {
				sum += a[j]
			}
		}
		if seen[sum] == 0 {
			seen[sum] = bf
		} else {
			if seen[sum] & bf == 0 {
				return false
			} else {
				seen[sum] = seen[sum] | bf
			}
		}
	}

	// if map was successfully builded, that sequence is ok
	
	return true
}
/*----------------------------------------------------------------------------*/
func sliceSum(a []int) int {

	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
