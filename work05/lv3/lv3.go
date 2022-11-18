package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 把用户信息写进文件 每次启动的时候通过文件读取用户信息
// 并且用户注册后的信息也可以追加到文件中
var userData = map[string]string{}

func main() {
	a := 0
	for {
		showMenu()
		fmt.Print(":")
		fmt.Scan(&a)
		switch a {
		case 1:
			login()
		case 2:
			register()
		case 3:
			return
		default:
			fmt.Println("please input again!")
		}
	}
}
func showMenu() {
	fmt.Println("**********")
	fmt.Println("* 1.登录 *")
	fmt.Println("* 2.注册 *")
	fmt.Println("* 3.退出 *")
	fmt.Println("**********")
}
func login() {
	username := ""
	password := ""
	fmt.Print("username:")
	fmt.Scan(&username)
	fmt.Print("password:")
	fmt.Scan(&password)
	readFromFile()
	if userData[username] == password {
		fmt.Println("login ok")
	} else {
		fmt.Println("fail to login")
	}
}

func register() {
	username := ""
	password := ""
	fmt.Print("username:")
	fmt.Scan(&username)
	fmt.Print("password:")
	fmt.Scan(&password)
	if userData[username] != "" {
		fmt.Println("fail to register because username is used")
		return
	}
	userData[username] = password
	addToFile()
	fmt.Println("register ok")

}
func addToFile() {
	f, err1 := os.Create("user.data")
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	defer f.Close()
	buf, err2 := json.Marshal(userData)
	if err2 != nil {
		fmt.Println("err = ", err2)
		return
	}
	_, err3 := f.Write(buf)
	if err3 != nil {
		fmt.Println("err3 = ", err3)
		return
	}

}
func readFromFile() {
	f, err1 := os.Open("user.data")
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}
	defer f.Close()
	fileData := ""
	buf := make([]byte, 1024)
	for {
		n, err2 := f.Read(buf)
		if err2 != nil && n != 0 {
			fmt.Println("err = ", err2)
		}
		if n == 0 {
			break
		}
		fileData += string(buf[:n])
	}
	json.Unmarshal([]byte(fileData), userData)
}
