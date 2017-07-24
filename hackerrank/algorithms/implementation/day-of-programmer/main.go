package main

import "fmt"
import "strconv"

func main() {
	var year int
	fmt.Scan(&year)
	fmt.Println(solve(year))
}

func solve(year int) string {
	var leap bool
	switch {
	case year == 1918:
		return "26.09.1918"
	case year < 1918:
		leap = (year%4 == 0)
	case year > 1918:
		leap = (year%400 == 0) || ((year%4 == 0) && (year%100 != 0))
	}

	if leap {
		return "12.09" + strconv.Itoa(year)
	}
	// else
	return "13.09" + strconv.Itoa(year)
}
