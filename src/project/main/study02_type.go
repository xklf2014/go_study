package main
import "fmt"

func main(){
	//basic date type
	// u开头为无符号
	var i1 int = -1
	var i2 int8 = -100
	var i3 int16 = -999
	var i4 int32 = -99999
	var i5 int64 = -99999999
	var u1 uint = 1
	var u2 uint8 = 100
	var u3 uint16 = 999
	var u4 uint32 = 9999
	var u5 uint64 = 99999999

	var f1 float32 = 3.14
	var f2 float64 = 3.3333333333333333

	var flag bool = true

	var str string = "strings"
	
	fmt.Println("-------int---------")
	fmt.Println(i1,i2,i3,i4,i5)
	fmt.Println("-------uint---------")
	fmt.Println(u1,u2,u3,u4,u5)
	fmt.Println("-------float---------")
	fmt.Println(f1,f2)
	fmt.Println("-------boolean---------")
	fmt.Println(flag)
	fmt.Println("-------string---------")
	fmt.Println(str)

}