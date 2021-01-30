package main

import "fmt"

func test(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	test()
	fmt.Println("---------------")
	test(1)
	fmt.Println("---------------")
	test(11, 22, 33, 44, 55, 66, 77, 88, 99)
	fmt.Println("---------------")
}
