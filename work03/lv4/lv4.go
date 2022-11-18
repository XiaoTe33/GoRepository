package main

import "fmt"

func main() {
	over := make(chan bool)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			ch <- 1
			//for循环还没等输出就把i加上去了，用个chanel锁一下
			if i == 9 {
				over <- true
			}
		}()
		<-ch
		//if i == 9 {
		//	over <- true
		//}
		//此段if代码放在主协程会和<-over形成一个死锁，所以放到goroutine里
	}
	<-over
	fmt.Println("over!!!")
}
