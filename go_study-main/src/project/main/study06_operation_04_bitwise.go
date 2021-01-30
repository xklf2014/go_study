package main
import "fmt"

func main(){
	var a int = -2
	fmt.Println(a>>1)
	var c int = -2
	fmt.Println(c<<1)

	var n1 int = 7
	var n2 int = 5
	fmt.Println("n1&n2=",n1&n2)
	fmt.Println("n1|n2=",n1|n2)
	fmt.Println("n1^n2=",n1^n2)
}