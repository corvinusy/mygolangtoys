package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// read file
	words := readFile("./input.txt")
	fmt.Println("Num of words =", len(words))
	// define word
	const word string = "Copernicus"
	// call
	isContain := containsString1(words, word)
	fmt.Println(isContain)
}

func readFile(fileName string) []string {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	result := make([]string, 0)

	for i := range lines {
		result = append(result, strings.Split(lines[i], " ")...)
	}

	return result
}

func containsString1(arr []string, element string) bool {
	for i := range arr {
		if arr[i] == element {
			return true
		}
	}
	return false
}

func containsString2(arr []string, element string) bool {
	for _, elem := range arr {
		if elem == element {
			return true
		}
	}
	return false
}
