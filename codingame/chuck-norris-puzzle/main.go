package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()

	//	fmt.Println(text)

	encoded := strings.Trim(encodeText(getBinaryText(text)), " ")
	fmt.Println(encoded) // Write answer to stdout
}

func getBinaryText(text string) string {
	var res string
	for _, v := range text {
		res += getBinaryLetter(v)
	}
	return res
}

func getBinaryLetter(r rune) string {
	var res string
	for i := 0; i < 7; i++ {
		if r%2 == 0 {
			res += "0"
		} else {
			res += "1"
		}
		r = r >> 1
	}
	// reverse the result string
	runes := []rune(res)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func splitSeries(text string) []string {
	res := make([]string, 0)
	var cur rune = rune(text[0])
	var str string
	for _, v := range text {
		if v == cur {
			str += string(v)
			continue
		}
		res = append(res, str)
		cur = v
		str = string(cur)
	}
	res = append(res, str) // last seria
	return res

}

func encodeText(text string) string {

	var res string
	//	fmt.Println(text)
	seria := splitSeries(text)
	//	fmt.Printf("%v\n", seria)

	for _, v := range seria {
		if v[0] == '1' {
			res += "0"
		} else {
			res += "00"
		}
		res += " "
		for _ = range v {
			res += "0"
		}
		res += " "
	}
	return res
}
