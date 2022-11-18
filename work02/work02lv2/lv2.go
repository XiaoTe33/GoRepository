package main

import "fmt"

type Behavior interface {
	Speak()
}
type Behavior1 interface {
	Eat()
}
type Dog struct {
}
type Cat struct {
}

func (d Dog) Speak() {
	fmt.Println("汪汪")
}
func (d Dog) Eat() {
	fmt.Println("狗在吃东西")
}
func (c Cat) Speak() {
	fmt.Println("喵喵")
}
func (c Cat) Eat() {
	fmt.Println("猫在吃东西")
}
func main() {
	var d Behavior = Dog{}
	d.Speak()
	var d1 Behavior1 = Dog{}
	d1.Eat()
	var c Cat
	c.Speak()
	c.Eat()

}
