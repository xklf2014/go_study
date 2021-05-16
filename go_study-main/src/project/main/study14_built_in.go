package main

import "fmt"

func main() {
	str := "hello go"
	fmt.Println(len(str))

	//new分配内存，new函数的实参是一个类型而不是具体的数值，new函数返回的对应类型指针num *int
	num := new(int)
	fmt.Printf("num的类型%T,num的值为:%v,num的地址为%v,num的指针指向%v", num, num, &num, *num)
}
