package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quite chan bool) {

	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quite:
			wg.Done()
			return
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true
	wg.Wait()
}
