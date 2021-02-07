package main

import "fmt"

//闭包保留上一次的结果引用，可以返回使用
func getSum() func(num int) int {
	var sum int = 10
	return func(num int) int {
		sum = sum + num
		return sum
	}

}

func main() {
	f := getSum()
	fmt.Println(f(20))
	fmt.Println(f(30))

}
