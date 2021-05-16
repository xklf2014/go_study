package main

import (
	"errors"
	"fmt"
)

func main() {
	//test2()
	err := test3()
	if err != nil {
		fmt.Println("自定义异常")
		panic(err)
	}
	fmt.Println("程序正确执行")
}

func test2() {
	//利用defer recover来捕获错误
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获错误,err:", err)
		}
	}()
	a := 10
	b := 0
	result := a / b
	fmt.Println(result)
}

func test3() (err error) {
	a := 10
	b := 0

	if b == 0 {
		return errors.New("除数不能为0")
	} else {
		result := a / b
		fmt.Println(result)
		return nil
	}

}
