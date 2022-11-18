package main

import "fmt"

// 通过数组或者切片或者map 实现用户的注册以及注册后的用户登录
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
	fmt.Println("register ok")

}
