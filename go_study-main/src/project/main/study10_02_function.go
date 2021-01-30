package main

import "fmt"

func exchangeNum(num1 int, num2 int) (int, int) {
	var tmp = num1
	num1 = num2
	num2 = tmp
	return num1, num2
}

func main() {
	var n1 int = 10
	var n2 int = 20
	fmt.Printf("交换前为n1=%d , n2=%d \n", n1, n2)
	n1, n2 = exchangeNum(n1, n2)
	fmt.Printf("交换后为n1=%d , n2=%d \n", n1, n2)
}
