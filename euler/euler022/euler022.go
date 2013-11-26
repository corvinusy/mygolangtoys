package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func main() {
	content, err := ioutil.ReadFile("names.txt")
	if err != nil {
		panic("File not read")
	}

	lines := strings.Split(string(content), ",")
	for i , line := range lines {
		lines[i], _ = strconv.Unquote(line)
	}

	sort.Strings(lines)

	fmt.Println(lines[938])

	sum := 0
	
	for i, _ := range lines {
		sum += letsum(lines[i]) * (i + 1)
	}

	fmt.Println(sum)
}
/*-----------------------------------------------------------------------------*/
func letsum(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) - 64
	}
	return sum
}
