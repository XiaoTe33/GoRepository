package main

import "fmt"

func solution(s string) {
	var a string
	a = s
	for _, j := range a {
		if i := 1; i <= len(s)/2 {
			i++
			fmt.Printf("%c", j)
		}
	}

	return
}
func solution2(s string) {
	var a string
	a = s
	i := 0
	for _, j := range a {
		if i < int(len(s)/6) {
			i++
			fmt.Printf("%c", j)
			fmt.Println(i, len(s)/6)
		}
	}

	return
}
func main() {
	var s string = "12321"
	for _, i2 := range s {
		fmt.Printf("%c", i2)
	}
	fmt.Println(s)
	solution2("可见度发挥时刻提防")
}
