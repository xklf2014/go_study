package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum >= 300 {
			break
		}
	}

	fmt.Println("1-100之间的累加和第一次超过300的值为", sum)
	fmt.Println("---------------分隔符-------------------")

	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i=%d and j=%d \n", i, j)
			if i == 2 && j == 2 {
				break
			}
		}
	}

	fmt.Println("---------------分隔符-------------------")
	//设置标签，可以通过break指定标签对应的循环
outer:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i=%d and j=%d \n", i, j)
			if i == 2 && j == 2 {
				break outer
			}
		}
	}
}
