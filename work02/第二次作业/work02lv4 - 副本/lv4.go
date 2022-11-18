package main

import (
	"fmt"
)
type MyInterface interface {
	Len1() int
	Less1(i, j int) bool
	Swap1(i, j int)
}//定义自己的接口
type MyArray []int
func (receiver MyArray) Len1() int {
	return len(receiver)
} //返回长度
func (receiver MyArray) Less1(i, j int) bool {
	return receiver[i] < receiver[j]
} //比较出较小
func (receiver MyArray) Swap1(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
} //交换
func MySort(arr MyArray) {
	for i := 0; i < arr.Len1(); i++ {
		for j := 0; j < arr.Len1()-i-1; j++ {
			if !arr.Less1(j, j+1) {
				arr.Swap1(j, j+1)
			}
		}
	} //排序
	fmt.Println(arr) //输出排序后的结果
}//排序函数
func main() {
	var arr = MyArray{2, 3, 1, 5, 4, 7, 6, 9, 8, 0}//定义数组
	MySort(arr)//排序
}



