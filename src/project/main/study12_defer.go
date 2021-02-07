package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
}

func add(num1 int, num2 int) int {
	//在GoLang中，程序遇到defer关键字，不会立即执行defer后面的语句，
	//而是将defer后面的语句压入栈中，然后继续执行后续的语句（栈:先进后出）
	defer fmt.Println("num1:", num1) // 第三个打印
	defer fmt.Println("num2:", num2) // 第二个打印

	//在函数执行完毕后，从栈中取出语句开始执行
	/*
		打印顺序
		sum: 180
		num2: 20
		num1: 10
		180
	*/
	num1 += 100
	num2 += 50
	var sum int = num1 + num2
	fmt.Println("sum:", sum) // 第一个打印
	return sum

}
