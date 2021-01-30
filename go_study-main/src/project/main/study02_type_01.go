package main
import "fmt"

func main(){
	//basic date type
	// u开头为无符号
	var i1 int = -1
	var i2 int8 = 127
	//var i21 int8 = 128 constant 128 overflows int8
	var i3 int16 = -999
	var i4 int32 = -99999
	var i5 int64 = -99999999
	var u1 uint = 1
	var u2 uint8 = 255
	//var u21 uint8 = 256 constant 256 overflows uint8
	var u3 uint16 = 999
	var u4 uint32 = 9999
	var u5 uint64 = 99999999

	var f1 float32 = 3.14
	var f11 float32 = -3.14
	var f12 float32 = 3.14e2
	var f13 float32 = 314e-2
	var f2 float64 = 3.3333333333333333
	var f3 = 3.1415926

	var flag bool = true

	var str string = "strings"
	
	fmt.Println("-------int---------")
	fmt.Println(i1,i2,i3,i4,i5)
	fmt.Println("-------uint---------")
	fmt.Println(u1,u2,u3,u4,u5)
	fmt.Println("-------float---------")
	fmt.Println(f1,f11,f12,f13,f2)
	fmt.Println("-------boolean---------")
	fmt.Println(flag)
	fmt.Println("-------string---------")
	fmt.Println(str)

	fmt.Printf("i1的类型是%T \n",i1)
	fmt.Printf("f3的类型是%T",f3)

}