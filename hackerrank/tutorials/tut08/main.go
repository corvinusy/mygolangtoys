package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n           int
		name, phone string
		ok          bool
	)

	fmt.Scan(&n)
	d := make(map[string]string, n)
	rd := bufio.NewScanner(os.Stdin)
	rd.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		if rd.Scan() {
			name = rd.Text()
		}
		if rd.Scan() {
			phone = rd.Text()
		}
		d[name] = phone
	}

	for rd.Scan() {
		name = rd.Text()
		phone, ok = d[name]
		if ok {
			fmt.Printf("%s=%s\n", name, phone)
		} else {
			fmt.Println("Not found")
		}
	}
}
