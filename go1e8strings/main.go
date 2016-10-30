//+build linux freebsd darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	a := make([]string, 1e8)

	for i := 0; i < len(a); i++ {
		a[i] = ""
	}

	fmt.Printf("pushed %d strings\n", len(a))
	cmd := exec.Command("/usr/bin/cat", "/proc/"+strconv.Itoa(os.Getpid())+"/status")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
