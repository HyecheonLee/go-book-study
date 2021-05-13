package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str = "Hello World"
	var slice = []byte(str)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%X\n", stringHeader.Data)
	fmt.Printf("str:\t%X\n", sliceHeader.Data)
}
