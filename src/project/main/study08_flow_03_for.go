package main
import "fmt"

func main(){
	var sum int = 0;
/* 	var i int
	for i = 1 ; i<=5 ; i++ {
		sum += i
	} */

	// for循环 --> 初始声明不可以用var声明变量，可以用:=
	for i := 1 ; i<=5 ; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("---------------分割线------------------")
	var sum1 int = 0
	for i := 1; i<=100;i++ {
		if i % 2 == 0 {
			sum1 += i
		}
	}

	fmt.Println("100以内的偶数和为",sum1)

	j := 1
	for j<= 5 {
		fmt.Println("Go Lang")
		j++
	}
}