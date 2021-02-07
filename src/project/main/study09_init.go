package main

import (
	"fmt"
	testutils "project/main/utils"
)

//执行顺序： import(全局变量的定义-> init)->全局变量的定义-> init -> main
var num int = test()

func test() int {
	fmt.Println("test函数被调用")
	return 10
}

func init() {
	fmt.Println("init 函数被调用")
}

func main() {
	fmt.Println("main 函数被调用")
	fmt.Printf("testutils 被调用,姓名：%s,年龄：%d,性别:%s", testutils.Name, testutils.Age, testutils.Gender)
}
