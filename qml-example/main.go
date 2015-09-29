package main

import (
	"fmt"
	"os"

	qml "gopkg.in/qml.v1"
)

func main() {
	err := qml.Run(run)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()
	component, err := engine.LoadFile("/home/corvinus/ws/go/local/src/qml-example/example.qml")
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)

	win.Show()
	win.Wait()

	return nil
}
