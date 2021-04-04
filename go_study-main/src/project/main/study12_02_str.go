package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("gogogo", "go", "go lang\n", -1)) //全部替换
	fmt.Println(strings.Replace("gogogo", "go", "go lang\n", 1))  //替换一个

	split := strings.Split("go-java-python-scala", "-")
	fmt.Println(split)

	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("gogogo"))
	fmt.Println(strings.TrimRight("   ha ha ha   ", " "))
	fmt.Println(strings.TrimLeft("   ha ha ha   ", " "))
	//去除左右两边的空格
	fmt.Println(strings.TrimSpace("   ha ha ha   "))
	//去除两边指定的字符
	fmt.Println(strings.Trim("---ha ha ha---", "-"))

	fmt.Println(strings.HasPrefix("https://www.baidu.com", "https"))
	fmt.Println(strings.HasSuffix("https://www.baidu.com", "com"))

}
