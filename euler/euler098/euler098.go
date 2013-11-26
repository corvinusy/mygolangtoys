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

	for i, s := range words {
		words[i], err = strconv.Unquote(s)
		if err != nil {
			panic("Unqouting fails")
		}
	}

	
	fmt.Println(words[1],start_anagramm(words[1]))

}
/*-----------------------------------------------------------------------------*/
func start_anagramm(s string) string {

	template := [...]rune{'1','0','2','3','4','5','6','7','8','9'}

	r := make([]rune, len(s))

	k := 0

	for i, _ := range s {
		if r[i] == 0 {
			r[i] = template[k]
			k++

			for j, _ := range s {
				if s[j] == s[i] {
					r[j] = r[i]
				}
			}
		}
	}
	
	return string(r)
}
/*-----------------------------------------------------------------------------*/
func next_anagramm(s, a string) (string, bool) {
	return s, true
}
/*-----------------------------------------------------------------------------*/
func test_anagramm(s, a, word string) bool {
	if len(s) != len(word) {return false}

	for i, _ := range s {
		for j, _ := range word {
			if word[j] == s[i] {
				word[j] = a[i]
				break
			}
		}
		return false
	}
	return true
}
