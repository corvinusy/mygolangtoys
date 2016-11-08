package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	s, err := rd.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	if strings.HasSuffix(s, "PM") {
		h, err := strconv.ParseInt(s[:2], 10, 64)
		if err != nil {
			panic(err.Error())
		}
		if h != 12 {
			h = (h + 12) % 24
		}
		hStr := strconv.FormatInt(h, 10)
		if len(hStr) == 1 {
			hStr = "0" + hStr
		}
		s = hStr + s[2:]
	} else if s[:2] == "12" {
		s = "00" + s[2:]
	}

	fmt.Println(s[:len(s)-2])

}
