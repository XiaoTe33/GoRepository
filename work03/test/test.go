package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := []byte(s)
	fmt.Println(a)
}
