package main

import "fmt"

func test(num1 int, num2 int, testFunc func(int, int)) {
	testFunc(num1, num2)
}

type myFunc func(int, int)

func test1(num1 int, num2 int, testFunc myFunc) {
	testFunc(num1, num2)
}

func add(n1 int, n2 int) {
	fmt.Println(n1 + n2)
}

func addAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	test(5, 10, add)

	type myInt int
	var num1 myInt = 30
	fmt.Println("num1", num1)

	var num2 int = 30
	num2 = int(num1) // 类型需要强制转换
	fmt.Println("num2", num2)

	test1(10, 20, add)

	a, b := addAndSub(20, 10)
	fmt.Println(a, b)
}
