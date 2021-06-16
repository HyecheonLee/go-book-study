package main

import "fmt"

func divided(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divided(9, 3)
	divided(9, 0)
}
