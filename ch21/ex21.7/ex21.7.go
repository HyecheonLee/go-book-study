package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer file.Close()
	writeHello(func(mes string) {
		fmt.Fprintln(file, mes)
	})
}
