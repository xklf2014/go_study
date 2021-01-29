package main
import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 96
	var s1 string = fmt.Sprintf("%d",i)
	fmt.Printf("s1对应的数据类型为%T,结果为%v \n",s1,s1)

	var f float32 = 3.1415
	var s2 string = fmt.Sprintf("%f",f)
	fmt.Printf("s2对应的数据类型为%T,结果为%q \n",s2,s2)

	var b bool = true
	var s3 string = fmt.Sprintf("%t",b)
	fmt.Printf("s3对应的数据类型为%T,结果为%q \n",s3,s3)

	var bt byte = 'a'
	var s4 string = fmt.Sprintf("%c",bt)
	fmt.Printf("s4对应的数据类型为%T,结果为%q \n",s4,s4)

	fmt.Println("------------------分割线---------------------")

	var i1 int = 18
	//第一个参数需要int64类型，第二个为进制
	var s5 string = strconv.FormatInt(int64(i1),10)
	fmt.Printf("s5对应的数据类型为%T,结果为%q \n",s5,s5)

	var f1 float64 = 4.141241
	// 参数一：需要转换的类型
	// 参数二：'f'[ddd.ddd]
	// 参数三：保留小数点后几位
	// 参数四：float64类型
	var s6 string = strconv.FormatFloat(f1,'f',9,64)
	fmt.Printf("s6对应的数据类型为%T,结果为%q \n",s6,s6)

	var b1 bool = false
	var s7 string = strconv.FormatBool(b1)
	fmt.Printf("s7对应的数据类型为%T,结果为%q \n",s7,s7)
}