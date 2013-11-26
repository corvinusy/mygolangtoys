package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {

	var (
		strs []string
		count, wordnum int
	)

	//read file into source

	const LIMIT = 1000

	triangles := make(map[int]bool, 0)

	for i := 1; i <= LIMIT; i++ {
		triangles[(i * (i + 1)) >> 1] = true
	}

	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic("File not read")
	}

	source := string(content)

	strs = strings.SplitAfter(source,",")

	count = 0

	for _, s := range strs {
		if s[len(s)-1] == ',' {
			s = s[:len(s)-1]
		}
		s, err := strconv.Unquote(s)
		if err != nil {
			panic("Unqouting fails")
		}

		wordnum = 0
		for _, a := range s {
			wordnum += int(a - 'A' + 1)
		}
		if triangles[wordnum] {
			count++
		}
	}

	fmt.Println(count)

}


