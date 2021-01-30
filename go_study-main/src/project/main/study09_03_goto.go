package main

import "fmt"

func main() {
	fmt.Println("hello GoLang 1")
	fmt.Println("hello GoLang 2")
	if 1 == 1 {
		goto label1
	}
	fmt.Println("hello GoLang 3")
	fmt.Println("hello GoLang 4")
	fmt.Println("hello GoLang 5")
label1:
	fmt.Println("hello GoLang 6")
	fmt.Println("hello GoLang 7")
}
