package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	pattern = "Hello my Golang"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inout := make(chan string) // data transferring between
	// start reading input routine
	fmt.Print("Enter text ('quit' to stop): ")
	go readInput(reader, inout)
	// start processing input routine
	go processInput(inout)
	// wait proc
	time.Sleep(1 * time.Hour)
}

//------------
func readInput(reader *bufio.Reader, out chan string) {
	for {
		input, _ := reader.ReadString('\n')
		//fmt.Printf("%+q", input)
		if input == "quit\n" {
			os.Exit(0)
		}
		out <- input
	}
}

//------------
func processInput(in chan string) {
	for line := range in {
		total := 0
		for i := range line {
			if (i < len(pattern)) && (line[i] == pattern[i]) {
				total++
			}
		}
		fmt.Println(total)
	}
}
