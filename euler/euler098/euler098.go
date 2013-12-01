package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {

	//read file into source

	content, err := ioutil.ReadFile("euler098/words.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content),"\n")

	words := strings.Split(source,",")

	// unquote words

	for i, s := range words {
		words[i], err = strconv.Unquote(s)
		if err != nil {
			panic("Unqouting fails")
		}
	}

	// find anagrams

	anamap := make(map[string]string, 0)

	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if isAnagramm(words[i], words[j]) {
				anamap[words[i]] = words[j]
				fmt.Println(words[i], words[j])
			}
		}
	}

	fmt.Println(isMaps("CARE", "RACE", 1296, 9216))

	fmt.Println(isMaps("RACE", "CARE", 9216, 1296))

//	return

	sqrs := prepareSquares()

	for _, n1 := range sqrs {
		for _, n2 := range sqrs {
			if isMatches(n1, n2) {
				for w1, _ := range anamap {
					if isMaps(w1, anamap[w1], n1, n2) {
						fmt.Println(w1, anamap[w1], n1, n2)
					}
				}
			}
		}
	}
}
/*-----------------------------------------------------------------------------*/
func isAnagramm(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	alphs := make(map[rune]int, 0)

	for _, a := range s1 {
		alphs[a]++
	}

	for _, a := range s2 {
		alphs[a]--
	}

	for _, a := range s1 {
		if alphs[a] != 0 {
			return false
		}
	}

	return true
}
/*-----------------------------------------------------------------------------*/
func prepareSquares() []int {

	sqrs := make([]int, 0)

	for i := 1; i * i < 1e10; i++ {
		switch numlen(i*i)  {
		case 1, 2, 3, 4, 5, 6, 7, 9: sqrs = append(sqrs, i*i)
		}
	}
	return sqrs
}
/*-----------------------------------------------------------------------------*/
func numlen(n int) int {

	i := 0
	for n > 0 {
		n /= 10
		i++
	}
	return i
}
/*-----------------------------------------------------------------------------*/
func isMaps(w1, w2 string, n1, n2 int) bool {

	// check if n1 is a map of w1
	
	if len(w1) != numlen(n1) { return false }

	anmap := make(map[uint8]int, 0)
	dmap := make(map[int]bool, 0)

	for i := len(w1) - 1; i >= 0; i-- {
		if anmap[w1[i]] == 0 {
			if dmap[n1 % 10] { return false }
			dmap[n1 % 10] = true
			anmap[w1[i]] = n1 % 10 + 1
		} else {
			if anmap[w1[i]] != n1 % 10 + 1 {	return false }
		}

		n1 /= 10
	}

	for i := len(w2) - 1; i >= 0; i-- {
		if anmap[w2[i]] != n2 % 10 +1 {
			return false
		}
		n2 /= 10
	}
		
	return true
}
/*-----------------------------------------------------------------------------*/
func isMatches(n1, n2 int) bool {

	//check if n1 && n2 consists of same digits

	if n1 == n2 { return false }

	if numlen(n1) != numlen(n2) { return false }

	digits := make([]int, 10)

	for n1 > 0 {
		digits[n1 % 10]++
		n1 /= 10
	}

	for n2 > 0 {
		digits[n2 % 10]--
		n2 /= 10
	}
	
	for _, v := range digits {
		if v != 0 {
			return false
		}
	}

	return true
}
