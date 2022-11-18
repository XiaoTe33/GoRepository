package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func test01() {
	xiaoMing := Student{Name: "小明", Age: 18}
	fmt.Println(xiaoMing.Age, xiaoMing.Name)
}
func test02() {
	xiaoMing := Student{Name: "小明", Age: 18}
	fmt.Println(xiaoMing)
	xiaoMing.Age = 19
	fmt.Println(xiaoMing)

}
func main() {
	test02()

}
