package main
import "fmt"

func main(){

	var age int = 18
	fmt.Println("age的地址为",&age)

	//定义一个指针变量
	// *int 为指向int类型的指针
	// &age是一个地址，是point变量的值
	var point *int = &age
	fmt.Println(point)
	fmt.Println("point自己的地址为",&point)
	fmt.Println("point指向的值为",*point)

	fmt.Println("------------分隔符------------")

	var num int = 10
	fmt.Println("num : ",num)
	var ptr *int = &num
	*ptr = 18
	fmt.Println("修改后的num : ",num)
}