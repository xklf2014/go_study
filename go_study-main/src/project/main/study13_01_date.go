package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	fmt.Printf("%v------类型:%T", time.Now(), time.Now())
	fmt.Println()
	now := time.Now()
	fmt.Printf("年: %v \n", now.Year())
	fmt.Printf("月: %v \n", (int)(now.Month()))
	fmt.Printf("日: %v \n", now.Day())
	fmt.Printf("时: %v \n", now.Hour())
	fmt.Printf("分: %v \n", now.Minute())
	fmt.Printf("秒: %v \n", now.Second())

	fmt.Println("-----------------")
	fmt.Printf("当前日期为: %d-%d-%d %d:%d:%d", now.Year(), (int)(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	sprintf := fmt.Sprintf("当前日期为: %d-%d-%d %d:%d:%d", now.Year(), (int)(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(sprintf)

}
