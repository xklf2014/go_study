package main
import "fmt"

func main(){
	//+ 可以作为数值类型运算
	var n1 int = 5
	var n2 int = 4 + 7
	var n3 int = n1 + 10
	fmt.Printf("n1: %d , n2: %d , n3: %d \n",n1,n2,n3)

	// + 可以作为字符串拼接
	var s1 string = "Hello " + "GoLang"
	fmt.Println(s1)

	var n4 int = 100-50
	fmt.Println(n4)

	fmt.Println(10/3) //取整，按整数计算
	fmt.Println(10.0/3) //按照浮点类型计算

	//取模计算
	fmt.Println(10%3)
	fmt.Println(-10%3)
	fmt.Println(10%-3)
	fmt.Println(-10%-3)

	var a int = 10
	// ++ -- 自增自减操作，只能单独使用，不能参与计算
	// ++ -- 不能放在变量前面
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	var num int = (10 + 20) % 3  + 3 - 5 * 7
	fmt.Println(num)

	var num2 int = 10
	num2 += 20
	fmt.Println("num2 += 20 is ",num2)
	num2 -= 10
	fmt.Println("num2 -= 10 is ",num2)
	num2 *= 10
	fmt.Println("num2 *= 10 is ",num2)
	num2 /= 2
	fmt.Println("num2 /= 2 is ",num2)

	var a1 int = 8
	var b1 int = 6
	fmt.Printf("a=%d  b=%d \n",a1,b1)
	// 通过中间变量进行交换
	var t int
	t=a1
	a1=b1
	b1=t
	fmt.Printf("a=%d  b=%d \n",a1,b1)

	a1,b1 = b1,a1
	fmt.Printf("a=%d  b=%d \n",a1,b1)
}