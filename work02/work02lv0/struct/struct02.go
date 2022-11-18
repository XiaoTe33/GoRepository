package main

import "fmt"

type people struct {
	Name string
	Age  int
}
type teacher struct {
	together people
	workTime int
}
type student struct {
	together  people
	learnTime int
}

func test03() {
	var stu student
	stu.together.Name = "xiaoHong"
	stu.learnTime = 24
	teac := teacher{
		together: people{
			Name: "huaHua",
			Age:  25,
		},
		workTime: 9,
	}
	fmt.Println(teac)
}

func main() {
	test03()
}
