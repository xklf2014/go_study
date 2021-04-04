package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello Go语言"
	fmt.Println(len(str))

	for i, value := range str {
		fmt.Printf("索引为:%d,值为:%c \n", i, value)
	}

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("索引为:%d,值为:%c \n", i, r[i])
	}

	//字符串转整数
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)
	//整数转字符串
	s1 := strconv.Itoa(666)
	fmt.Println(s1)

	//计算字符再字符串中出现的次数
	count := strings.Count("golangandjava", "a")
	fmt.Printf("包含a字符%d个", count)

	//不区分大小写比较字符串
	rs := strings.EqualFold("golang", "GOLAND")
	fmt.Println(rs)
	//区分大小写比较
	fmt.Println("golang" == "GOLANG")

	//计算字符在字符串中第一次出现的位置，不存在为-1
	index := strings.Index("golangandjava", "a")
	index2 := strings.Index("golangandjava", "ab")
	fmt.Println(index)
	fmt.Println(index2)
}
