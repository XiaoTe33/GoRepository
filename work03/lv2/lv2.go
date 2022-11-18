package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 1)

	go func(ch chan int) {
		for i := 1; i <= 100; i += 2 {
			ch <- i
			fmt.Println(i)
			<-ch
		}
		wg.Done()
	}(ch)
	time.After(10 * time.Millisecond)
	go func(ch chan int) {
		for i := 2; i <= 100; i += 2 {
			ch <- i
			fmt.Println(i)
			<-ch
		}
		wg.Done()
	}(ch)

	wg.Add(2)
	wg.Wait()
}
