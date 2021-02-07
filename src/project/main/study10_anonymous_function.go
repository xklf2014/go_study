package main

import "fmt"

var sub = func(num1 int, num2 int) int {
	return num1 - num2
}

func main() {
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println("result is ", result)

	fmt.Println("sub result is", sub(20, 10))

	by := func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Println("by result is ", by(10, 10))
}
