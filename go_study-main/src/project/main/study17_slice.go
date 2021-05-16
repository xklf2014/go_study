package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	fmt.Println(arr[1:3]) //左闭右开

	sliceArr := arr[0:3]
	fmt.Println(sliceArr)

	fmt.Println("slice arr 长度", len(sliceArr))
	fmt.Println("slice arr 容量", cap(sliceArr))

	sliceArr[2] = 20
	fmt.Println("arr:", arr)
	fmt.Println("sliceArr:", sliceArr)

	//定义切片，make函数的三个参数，1.切片的类型 2.切片的长度 3.切片的容量
	slice := make([]int, 4, 20)

	for i := 0; i < 4; i++ {
		slice[i] = i
	}

	//slice[0] = 1000000000000000000

	fmt.Println(slice)

	for v := range slice {
		fmt.Println("value is ", v)
	}

	fmt.Println(arr[:4])
	fmt.Println(arr[1:])

	var a []int = []int{1, 2, 3, 4, 5, 6, 7}
	var b []int = make([]int, 10)
	copy(b, a)
	fmt.Println(b)
}
