package main
import (
	"fmt"
	"strconv"
)

func main(){

	//string -> bool
	var s1 string = "true"
	var b bool
	// _可以忽略接收返回结果
	b,_ = strconv.ParseBool(s1)
	fmt.Printf("b的类型是%T,结果是%v \n",b,b)

	//string -> int64
	var s2 string = "199"
	var num int64
	num,_ = strconv.ParseInt(s2,10,64)
	fmt.Printf("num的类型是%T,结果是%v \n",num,num)

	//string -> float64
	var s3 string = "3.1415926"
	var f1 float64
	f1,_ = strconv.ParseFloat(s3,64)
	fmt.Printf("f1的类型是%T,结果是%v \n",f1,f1)

	//注意：无效类型转换会输出默认值,例如：s4 = "abc" ，将s1->bool
	var s4 string = "abc"
	var b2 bool 
	b2,_= strconv.ParseBool(s4)
	fmt.Printf("b2的类型是%T,结果是%v \n",b2,b2)

	var f2 float64
	f2,_= strconv.ParseFloat(s4,64)
	fmt.Printf("f2的类型是%T,结果是%v \n",f2,f2)
	

}