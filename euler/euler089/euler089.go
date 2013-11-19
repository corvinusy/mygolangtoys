package main

import (
    "fmt"
	"strings"
	"io/ioutil"
)

func main() {

	raw_bytes, err := ioutil.ReadFile("roman.txt")
	if err != nil {
		panic("File not read")
	}

	strs := strings.Split(string(raw_bytes), "\n") 

//	fmt.Println(len(strs))
	count := 0

	for i, s1 := range strs {
		n := rom2int(s1)
		s2 := int2rom(n)
		count += len(s1)-len(s2)
		fmt.Println(i, s1, s2, n, len(s1)-len(s2) )
	}

	fmt.Println(count)


}
/*-----------------------------------------------------------------------------*/
func rom2int(s string) int {
	result := 0
	
	s = strings.Replace(s,"IV","IIII",-1 )
	s = strings.Replace(s,"IX","VIIII",-1 )
	s = strings.Replace(s,"XL","XXXX",-1 )
	s = strings.Replace(s,"XC","LXXXX",-1 )
	s = strings.Replace(s,"CD","CCCC",-1 )
	s = strings.Replace(s,"CM","DCCCC",-1 )
	
	for _, a := range s {
		switch a {
		case 'M': result += 1000
		case 'D': result += 500
		case 'C': result += 100
		case 'L': result += 50
		case 'X': result += 10
		case 'V': result += 5
		case 'I': result += 1
		}
 	}
	return result
}
/*-----------------------------------------------------------------------------*/
func int2rom(n int) string {
	s := ""
	for n > 0 {
		switch {
		case n / 1000 > 0 :
			s += "M"
			n -= 1000
		case n / 900 > 0 :
			s += "CM"
			n -= 900
		case n / 500 > 0 :
			s += "D"
			n -= 500
		case n / 400 > 0 :
			s += "CD"
			n -= 400
		case n / 100 > 0 :
			s += "C"
			n -= 100
		case n / 90 > 0 :
			s += "XC"
			n -= 90
		case n / 50 > 0 :
			s += "L"
			n -= 50
		case n / 40 > 0 :
			s += "XL"
			n -= 40
		case n / 10 > 0 :
			s += "X"
			n -= 10
		case n / 9 > 0 :
			s += "IX"
			n -= 9
		case n / 5 > 0 :
			s += "V"
			n -= 5
		case n / 4 > 0 :
			s += "IV"
			n -= 4
		default :
			s += "I"
			n -= 1
		}
	}
	return s
}
