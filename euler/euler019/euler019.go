package main

import (
    "fmt"
)

func main() {

	sum := 0
	day1 := 1 // monday = 0, sunday = 6

	for i:=1901; i < 2001; i++ {
		sum += count_year_fmsundays(i, &day1);
	}

	fmt.Println(sum);
}
/*-----------------------------------------------------------------------------*/
func count_year_fmsundays(year int, day1 *int) int {


	is_leap := false
	maxday := 364
	if (year % 4 == 0) && (year != 1900) {
		is_leap = true
		maxday++
	}
	
	res := 0
	for i := 0; i <= maxday ; i++ {
		if is_first(i, is_leap) && is_sunday(i, *day1) {
			res++
		}
	}
	if is_leap {
		*day1 = (*day1 + 2) % 7
		fmt.Println(year, res)
	} else {
		*day1 = (*day1 + 1) % 7		
	}
	return res
}
/*-----------------------------------------------------------------------------*/
func is_first(day int, is_leap bool) bool {
	if is_leap {
		return (day == 0) || (day == 31) || (day == 60) || (day == 91) || (day == 121) ||
			(day == 152) || (day == 182) || (day == 213) || (day == 244) || (day == 274) ||
			(day == 305) || (day == 335)
	} else {
		return (day == 0) || (day == 31) || (day == 59) || (day == 90) || (day == 120) ||
			(day == 151) || (day == 181) || (day == 212) || (day == 243) || (day == 273) ||
			(day == 304) || (day == 334)
	}
}
/*-----------------------------------------------------------------------------*/
func is_sunday(day int, day1 int) bool {
	return (day + 1 + day1) % 7 == 0
}
