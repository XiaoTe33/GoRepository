package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add(ch chan int) {

	for i := 0; i < 50000; i++ {
		ch <- i
		x = x + 1
		<-ch
	}
	wg.Done()
}
func main() {
	ch := make(chan int, 1)
	wg.Add(2)
	go add(ch)
	go add(ch)
	wg.Wait()
	fmt.Println(x)
}
