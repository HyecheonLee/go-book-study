package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", t)
	case float64:
		fmt.Printf("v is float64 %f\n", t)
	case string:
		fmt.Printf("v is string %s\n", t)
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("hello")
	PrintVal(Student{15})
}
