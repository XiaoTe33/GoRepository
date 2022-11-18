package main

import "fmt"

type MyFunc func(float64, float64) float64

func calculator(a float64, counter MyFunc, b float64) (result float64) { //计算器本体
	result = counter(a, b)
	return
}
func add(a float64, b float64) (result float64) { //加
	temp := a
	a += b
	println("----------------------------")
	fmt.Println(temp, "+", b, "=")
	result = a
	return result
}
func sub(a float64, b float64) (result float64) { //减
	temp := a
	a -= b
	println("----------------------------")
	fmt.Println(temp, "-", b, "=")
	result = a
	return
}
func mul(a float64, b float64) (result float64) { //乘
	temp := a
	a *= b
	println("----------------------------")
	fmt.Println(temp, "*", b, "=")
	result = a
	return
}
func div(a float64, b float64) (result float64) { //除
	temp := a
	a /= b
	println("----------------------------")
	fmt.Println(temp, "/", b, "=")
	result = a
	return
}
func main() {

	var sign rune
	var a, b float64

	println("请输入第一个数字：")
	fmt.Scan(&a)
	result := a
	for i := 1; i == 1; i += 0 {
		fmt.Println("请输入运算符号：")
		fmt.Println("(1.+  2.-  3.*  4./)")
		fmt.Scan(&sign)
		if sign == 0 {
			fmt.Println("欢迎下次使用")
			break
		}
		println("请输入第二个数字：")
		fmt.Scan(&b)
		switch sign {
		case 1:
			result = calculator(result, add, b)
		case 2:
			result = calculator(result, sub, b)
		case 3:
			result = calculator(result, mul, b)
		case 4:
			result = calculator(result, div, b)
		}
		fmt.Println(result)
	}

}
