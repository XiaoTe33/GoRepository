package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err := os.Create("plan.txt")
	if err != nil {
		return
	}
	var s = []string{"I’m not afraid of difficulties and insist on learning programming"}
	for _, i := range s {
		n, err := file1.Write([]byte(i))
		if err != nil {
			fmt.Println("err = ", err)
			os.Exit(1)
		}
		if n != len(i) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}

	}
	file2, err := os.Open("plan.txt")
	if err != nil {
		fmt.Println(err)
	}
	reader := make([]byte, len(s[0]))
	for {
		n, _ := file2.Read(reader)
		if n == 0 {
			break
		}
	}
	fmt.Println(string(reader))
}
