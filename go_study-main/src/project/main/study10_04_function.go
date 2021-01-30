package main

import "fmt"

func update(num *int) {
	*num = 30
}

func main() {
	var n int = 10
	update(&n)
	fmt.Println(n)

	a := update
	fmt.Printf("a的类型：%T,update的类型:%T", a, update)
}
