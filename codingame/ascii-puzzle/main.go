// Read inputs from Standard Input
// Write outputs to Standard Output
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		w     int
		h     int
		text  string
		input []string
		data  []string
	)
	// reading input
	input = make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := scanner.Text()
		input = append(input, str)

	}
	//parsing input
	fmt.Sscan(input[0], &w)
	fmt.Sscan(input[1], &h)

	// preparing text
	for _, v := range input[2] {
		switch {
		case (v >= 'A') && (v <= 'Z'):
			text += string(v)
			continue
		case (v >= 'a') && (v <= 'z'):
			text += string(v - ('a' - 'A'))
			continue
		default:
			text += "?"
		}
	}

	data = input[3:]

	// writing output
	printText(w, h, newText, data)
}

func getLetterAscii(w, h int, letter rune, data []string) []string {

	var shift int

	if letter == '?' {
		shift = int('Z'-'A') + 1
	} else {
		shift = int(letter - 'A')
	}

	result := make([]string, 0)
	for row := 0; row < h; row++ {
		result = append(result, data[row][shift*w:w*(shift+1)])
	}
	return result
}

func printText(w, h int, text string, data []string) {

	asciiLetters := make([][]string, 0)

	for _, v := range text {
		asciiLetters = append(asciiLetters, getLetterAscii(w, h, v, data))
	}

	for i := 0; i < h; i++ {
		for j := 0; j < len(asciiLetters); j++ {
			fmt.Printf(asciiLetters[j][i])
		}
		fmt.Println()
	}
}
