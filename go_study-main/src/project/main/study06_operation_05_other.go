package main
import "fmt"

func main(){

	var age int = 19
	fmt.Println("age对应的存储空间地址为",&age)

	var ptr *int = &age
	fmt.Println("ptr对应的存储空间地址为",&ptr)
	fmt.Println("ptr对应的值为",*ptr)
}