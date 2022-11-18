package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0)

	for i := 0; i < 100; i++ {
		num := rand.Intn(10)
		slice = append(slice, num)
	}
	for i := 0; i < 100; i++ {
		fmt.Print(slice[i])
	}
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	fmt.Println()
	for i := 0; i < 100; i++ {
		fmt.Print(slice[i])
	}

}
