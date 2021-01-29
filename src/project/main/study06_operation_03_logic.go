package main
import "fmt"

func main(){
	//逻辑运算符 &&(逻辑与) ||（逻辑或） !(逻辑非)
	fmt.Println("--------逻辑与----------")
	fmt.Println("true && true=",true && true)
	fmt.Println("true && true=",true && false)
	fmt.Println("true && true=",false && true)
	fmt.Println("true && true=",false && false)
	fmt.Println("--------逻辑或----------")
	fmt.Println("true || true=",true || true)
	fmt.Println("true || true=",true || false)
	fmt.Println("true || true=",false || true)
	fmt.Println("true || true=",false || false)
	fmt.Println("--------逻辑非----------")
	fmt.Println("!false=",!false)
	fmt.Println("!true",!true)
}