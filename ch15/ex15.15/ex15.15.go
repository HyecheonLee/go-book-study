package main

import "fmt"

func main() {
	var str = "Hello World"
	var slice = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
