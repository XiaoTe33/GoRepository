package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Eat() {
	fmt.Println("狗在吃东西")
}
func (d Dog) Bark() {
	fmt.Println("汪汪")
}
func main() {

}
