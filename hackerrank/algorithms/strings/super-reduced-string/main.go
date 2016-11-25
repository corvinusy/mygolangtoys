package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var s string
	fmt.Scanln(&s)

	for i, ok := find(s); ok; i, ok = find(s) {
		if len(s) == 0 {
			break
		}
		s = reduce(s, i)
	}
	if len(s) == 0 {
		fmt.Println("Empty String")
	} else {
		fmt.Println(s)
	}
}

func find(s string) (int, bool) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return i, true
		}
	}
	return 0, false
}

func reduce(s string, i int) string {
	fmt.Println(s, i)
	return s[:i] + s[i+2:]
}
