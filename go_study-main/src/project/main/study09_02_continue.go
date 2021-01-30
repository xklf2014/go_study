package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("--------下面为打印1-100以内不能被6整除的数--------------")

	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
