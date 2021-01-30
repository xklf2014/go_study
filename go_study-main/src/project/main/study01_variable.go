package main
import "fmt"

//定义全局变量
var global_var = "我是全局变量"
//一次性声明
var (
	first = "hello"
	second = "go"
	third = "lang"
)

func main(){
	//变量声明
	var age int
	//变量的赋值
	age = 18
	fmt.Println("age=",age)

	//变量声明和赋值在一行
	var age2 int = 20
	fmt.Println("age2=",age2)

	var i int
	fmt.Println("未初始化=",i)

	//类型自动推断
	var j = 20
	fmt.Println("未声明类型=",j)

	var float = 3.14
	fmt.Println("float=",float)

	var str = "哈哈哈"
	fmt.Println("我是字符串=",str)

	//省略var，符号为:=，不能写成=
	gender := "男"
	fmt.Println("性别为:",gender)

	fmt.Println("==================分割线===================")

	//多变量声明
	var n1,n2,n3 int
	fmt.Println("n1",n1,"n2",n2,"n3",n3)

	var length,name, d =100,"张三",3.1415
	fmt.Println("length",length,"name",name,"double",d)

	a1,a2,a3 := "haha",999,5.555
	fmt.Println("a1,a2,a3:",a1,a2,a3)

	fmt.Println("此处打印全局变量:",global_var)

	fmt.Println("打印一次性声明:",first,second,third)
}